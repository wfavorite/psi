package main

import (
	"os/exec"
	"syscall"
	"time"

	"github.com/wfavorite/initq"
)

/*
	STUB: This is heavily stubbed.
	STUB:
	STUB: What should happen
	STUB: - Launch clients from a local (internal) copy.
	STUB: - Launch client(s) based on config.
	STUB: - Pass things like args and socket name to client.
	STUB: - Keeps PID and process info.
*/

/* ------------------------------------------------------------------------ */

// ClientCmd is the stats and info wrapper for the launched client.
type ClientCmd struct {
	cmd           *exec.Cmd `json:"-"`
	PID           int       `json:"pid"`
	MonitorTarget string    `json:"monitor_target"`
	MonitorArgs   string    `json:"monitor_arg"`
	Started       string    `json:"start_timestamp"`
}

/* ------------------------------------------------------------------------ */

type ClientLaunch struct {
	List []*ClientCmd `json:"list"`
}

/* ======================================================================== */

func (cd *CoreData) ClientLaunch() initq.ReqResult {

	if cd.List == nil {
		return initq.TryAgain
	}

	clil := new(ClientLaunch)
	clil.List = make([]*ClientCmd, 0)
	cd.Clil = clil

	cd.LaunchClient("io", "some 150000000 1000000000")
	cd.LaunchClient("io", "full 100000000 1000000000")

	return initq.Satisfied

}

/* ======================================================================== */

// ShutdownClients walks the process list and sends TERM signals to each.
func (cd *CoreData) ShutdownClients() {
	for _, cmdWrap := range cd.Clil.List {
		cmdWrap.cmd.Process.Signal(syscall.SIGTERM)
	}
}

/* ======================================================================== */

// LaunchClient starts a client using the provided arguments.
func (cd *CoreData) LaunchClient(target string, arg string) {

	cmdWrap := new(ClientCmd)
	cmdWrap.MonitorTarget = target
	cmdWrap.MonitorArgs = arg
	cmdWrap.cmd = exec.Command("../cpsi/cpsi", cd.List.Filename, target, arg)
	cmdWrap.cmd.Start()
	cmdWrap.PID = cmdWrap.cmd.Process.Pid
	cmdWrap.Started = time.Now().UTC().Format(time.RFC3339)
	// Don't wait.\

	// Add to list
	cd.Clil.List = append(cd.Clil.List, cmdWrap)

}
