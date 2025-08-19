package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/wfavorite/initq"
)

/* ------------------------------------------------------------------------ */

// Event defines the responsibilities of what an event must satisfy.
type Event interface {
	Handle()
}

/* ------------------------------------------------------------------------ */

// EventQ is the structure holding the event q (channel) and stats.
type EventQ struct {
	q     chan Event     `json:"-"`
	Stats map[string]int `json:"stats"`
}

/* ======================================================================== */

// StartEventQ defines the event Q, but does not start processing events.
func (cd *CoreData) StartEventQ() initq.ReqResult {

	if cd.Logr == nil {
		return initq.TryAgain
	}

	if cd.EvtQ != nil {
		return initq.Satisfied
	}

	eq := new(EventQ)
	eq.q = make(chan Event, 3)
	eq.Stats = make(map[string]int)

	cd.EvtQ = eq

	cd.Logr.Verbose.Println("EventQ established.")
	return initq.Satisfied

}

/* ======================================================================== */
/* EVENTS                                                                   */
/* ======================================================================== */

/* ------------------------------------------------------------------------ */

// ShutdownEvent is used to shutdown the service.
type ShutdownEvent struct {
	cd *CoreData
}

/* ======================================================================== */

// NewShutdownEvent creates a new ShutdownEvent and places it on the Q.
func (cd *CoreData) NewShutdownEvent() {
	evt := new(ShutdownEvent)
	evt.cd = cd

	cd.EvtQ.q <- evt
}

/* ======================================================================== */

// Handle calls all shutdown calls, and then closes the EventQ chan.
func (evt *ShutdownEvent) Handle() {

	// To be consistent... but this will be deleted and never seen.
	evt.cd.EvtQ.Stats["Shutdown"]++

	evt.cd.ShutdownClients()
	// Give the clients a bit to time to properly respond to the signal. This
	// involves:
	// - Getting the signal.
	// - Handling a gracefull shutdown.
	// - ...before the Listener exits.
	time.Sleep(time.Microsecond * 500)
	evt.cd.ShutdownListener()
	close(evt.cd.EvtQ.q)
}

/* ------------------------------------------------------------------------ */

// HeartbeatEvent is a periodic check-in from a client.
type HeartbeatEvent struct {
	cd  *CoreData
	cli *ClientConn
}

/* ======================================================================== */

// NewHeartbeatEvent creates a new heartbeat for the specified client.
func (cd *CoreData) NewHeartbeatEvent(cli *ClientConn) {

	evt := new(HeartbeatEvent)
	evt.cd = cd
	evt.cli = cli

	cd.EvtQ.q <- evt
}

/* ======================================================================== */

// Handle registers the hearbeat into the verbose logger.
func (evt *HeartbeatEvent) Handle() {

	evt.cd.EvtQ.Stats["Heartbeat"]++
	evt.cd.Logr.Verbose.Printf("Heart beat from pid %d.\n", evt.cli.PID)
}

/* ------------------------------------------------------------------------ */

// ThresholdEvent is an event that we are watching for.
type ThresholdEvent struct {
	cd  *CoreData
	cli *ClientConn
}

/* ======================================================================== */

// NewThresholdEvent creates a new threshold exceeded for the specified client.
func (cd *CoreData) NewThresholdEvent(cli *ClientConn) {

	evt := new(ThresholdEvent)
	evt.cd = cd
	evt.cli = cli

	cd.EvtQ.q <- evt
}

/* ======================================================================== */

// Handle registers the hearbeat into the verbose logger.
func (evt *ThresholdEvent) Handle() {

	evt.cd.EvtQ.Stats["Threshold"]++
	evt.cd.Logr.Verbose.Printf("Threshold exceeded trigger from pid %d.\n", evt.cli.PID)
}

/* ------------------------------------------------------------------------ */

// ClientRegistrationEvent is when a new client connects to the system.
type ClientRegistrationEvent struct {
	cd  *CoreData
	cli *ClientConn
}

/* ======================================================================== */

// NewClientRegistrationEvent creates and Qs a ClientRegistrationEvent
func (cd *CoreData) NewClientRegistrationEvent(cli *ClientConn) {

	evt := new(ClientRegistrationEvent)
	evt.cd = cd
	evt.cli = cli

	cd.EvtQ.q <- evt
}

/* ======================================================================== */

func (evt *ClientRegistrationEvent) Handle() {

	evt.cd.EvtQ.Stats["ClientRegistration"]++
	evt.cd.Logr.Normal.Printf("Client reports up. Pid %d.\n", evt.cli.PID)
}

/* ------------------------------------------------------------------------ */

type ClientExitsEvent struct {
	cd  *CoreData
	cli *ClientConn
}

/* ======================================================================== */

func (cd *CoreData) NewClientExitsEvent(cli *ClientConn) {
	evt := new(ClientExitsEvent)
	evt.cd = cd
	evt.cli = cli

	cd.EvtQ.q <- evt
}

/* ======================================================================== */

func (evt *ClientExitsEvent) Handle() {

	evt.cd.EvtQ.Stats["ClientExits"]++
	evt.cd.Logr.Normal.Printf("Client %d exits.", evt.cli.PID)

	evt.cd.Cash.Remove(evt.cli.PID)
}

/* ------------------------------------------------------------------------ */

type ObserveEvent struct {
	cd *CoreData
}

/* ======================================================================== */

func (cd *CoreData) NewObserveEvent() {

	evt := new(ObserveEvent)
	evt.cd = cd

	cd.EvtQ.q <- evt
}

/* ======================================================================== */

// Handle dumps the CoreData structure as an observability object.
func (evt *ObserveEvent) Handle() {

	evt.cd.EvtQ.Stats["Observe"]++

	if jdata, jerr := json.MarshalIndent(evt.cd, "", "  "); jerr == nil {
		if werr := os.WriteFile("ppsi.observe.json", jdata, 0644); werr == nil {
			evt.cd.Logr.Verbose.Println("Observability dropped.")
		} else {
			evt.cd.Logr.Normal.Println("Observability file write error:", werr.Error())
		}
	} else {
		evt.cd.Logr.Normal.Println("Observability file jsonify error:", jerr.Error())
	}

}
