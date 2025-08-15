package main

import "github.com/wfavorite/initq"

/* ------------------------------------------------------------------------ */

// Event defines the responsibilities of what an event must satisfy.
type Event interface {
	Handle()
}

/* ------------------------------------------------------------------------ */

// EventQ is the structure holding the event q (channel) and stats.
type EventQ struct {
	Q chan Event
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
	eq.Q = make(chan Event, 3)

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

	cd.EvtQ.Q <- evt
}

/* ======================================================================== */

// Handle calls all shutdown calls, and then closes the EventQ chan.
func (evt *ShutdownEvent) Handle() {

	evt.cd.ShutdownListener()

	close(evt.cd.EvtQ.Q)
}

/* ------------------------------------------------------------------------ */

type HeartbeatEvent struct {
	cd  *CoreData
	cli *ClientConn
}

/* ======================================================================== */

func (cd *CoreData) NewHeartbeatEvent(cli *ClientConn) {

	evt := new(HeartbeatEvent)
	evt.cd = cd
	evt.cli = cli

	cd.EvtQ.Q <- evt

	return
}

/* ======================================================================== */

func (evt *HeartbeatEvent) Handle() {

	evt.cd.Logr.Verbose.Printf("Heart beat from pid %d.\n", evt.cli.PID)
}

/* ------------------------------------------------------------------------ */

type ClientRegistrationEvent struct {
	cd  *CoreData
	cli *ClientConn
}

/* ======================================================================== */

func (cd *CoreData) NewClientRegistrationEvent(cli *ClientConn) {

	evt := new(ClientRegistrationEvent)
	evt.cd = cd
	evt.cli = cli

	cd.EvtQ.Q <- evt
}

/* ======================================================================== */

func (evt *ClientRegistrationEvent) Handle() {

	evt.cd.Logr.Normal.Printf("Client reports up. Pid %d.\n", evt.cli.PID)
}

type ClientExitsEvent struct {
	cd  *CoreData
	cli *ClientConn
}

func (cd *CoreData) NewClientExitsEvent(cli *ClientConn) {
	evt := new(ClientExitsEvent)
	evt.cd = cd
	evt.cli = cli

	cd.EvtQ.Q <- evt
}

func (evt *ClientExitsEvent) Handle() {

	evt.cd.Logr.Normal.Printf("Client %d exits.", evt.cli.PID)

	evt.cd.Cash.Remove(evt.cli.PID)
}
