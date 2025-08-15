package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/wfavorite/initq"
)

/* ------------------------------------------------------------------------ */

// Listener is the structure that contains the listener instance and related
// arguments / stats.
type Listener struct {
	l        net.Listener `json:"-"`
	Filename string       `json:"unix_socket_filename"`
}

/* ======================================================================== */

// StartListener starts the unix socket listener.
func (cd *CoreData) StartListener() initq.ReqResult {

	if cd == nil {
		log.Fatal("InitQ method called on a nil CoreData")
	}

	if cd.EvtQ == nil {
		return initq.TryAgain
	}

	if cd.Logr == nil {
		return initq.TryAgain
	}

	if cd.List != nil {
		return initq.Satisfied
	}

	list := new(Listener)
	list.Filename = UnixSocketName // Option to make it different.

	fmt.Print("Looking for a stale Unix socket...")

	if _, sErr := os.Stat(list.Filename); errors.Is(sErr, os.ErrNotExist) {
		fmt.Println("None.")
	} else {
		fmt.Println("Found.")

		fmt.Print("Removing the (stale) Unix socket...")
		os.Remove(list.Filename)
		fmt.Println("Done.")
	}

	// Listen on the Unix socket
	fmt.Print("Listening on the socket..")
	l, err := net.Listen("unix", list.Filename)
	if err != nil {
		fmt.Println("Failed.")

		fmt.Println("listen error:", err)
		return initq.Stop
	}
	fmt.Println("Good.")

	// We are good so make this available.
	list.l = l
	cd.List = list

	cd.Logr.Verbose.Println("Listener established and listening.")

	// Now start the listener thread.
	go cd.listen()

	return initq.Satisfied
}

/* ======================================================================== */

// ShutdownListener is used to shutdown the listener and remove the Unix
// socket.
func (cd *CoreData) ShutdownListener() {

	cd.List.l.Close()
	os.Remove(cd.List.Filename)
	cd.Logr.Normal.Println("Listener shutdown.")
}

/* ======================================================================== */

func (cd *CoreData) listen() {

	for {
		// Accept incoming connections
		conn, err := cd.List.l.Accept()

		if err != nil {

			if errors.Is(err, net.ErrClosed) {
				return
			}

			cd.Logr.Normal.Println("accept error:", err)
			continue
		}

		// Handle the connection in a goroutine
		go cd.handleClientConn(conn)
	}
}

/* ======================================================================== */

func (cd *CoreData) handleClientConn(conn net.Conn) {

	cd.Logr.Verbose.Println("Received client connection.")

	cliCon := cd.Cash.New()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)

		if err != nil {
			if errors.Is(err, io.EOF) {
				cd.NewClientExitsEvent(cliCon)
				return
			}

			// STUB: Perhaps keep track of these.
			cd.Logr.Normal.Printf("Failed to read client send - %s\n", err.Error())
			continue
		}

		// Convert to a properly sized string
		data := string(buf[:n])
		cd.processClientMessage(cliCon, data)

	}

}

/* ======================================================================== */

func (cd *CoreData) processClientMessage(cc *ClientConn, msg string) {

	switch {
	case strings.HasPrefix(msg, "ClientUp("):

		pidStr := strings.TrimSuffix(strings.TrimPrefix(msg, "ClientUp("), ")")

		if iParse, eParse := strconv.ParseInt(pidStr, 10, 0); eParse == nil {
			cc.PID = int(iParse)
		} else {
			// STUB: Return error, bail? What?
			fmt.Println("ClientUp pid parse fail:", eParse.Error())
			return
		}

		cd.NewClientRegistrationEvent(cc)

	case strings.HasPrefix(msg, "HeartBeat"):
		cd.NewHeartbeatEvent(cc)

	default:
		// STUB: Track these in stats and send to log
		fmt.Printf("Client sent unknown data.\n")
	}

}
