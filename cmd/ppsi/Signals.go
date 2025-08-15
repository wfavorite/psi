package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/wfavorite/initq"
)

type Signals struct {
	S chan os.Signal
}

func (cd *CoreData) RegisterSignals() initq.ReqResult {

	if cd.EvtQ == nil {
		return initq.TryAgain
	}

	sigs := new(Signals)
	sigs.S = make(chan os.Signal, 1)

	// Those to ignore
	signal.Ignore(syscall.SIGHUP)

	// Those to handle
	signal.Notify(sigs.S, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGXCPU)

	cd.Sigs = sigs

	return initq.Satisfied
}

func (cd *CoreData) HandleSignal(sig os.Signal) {

	// STUB: As of now, all events we have registered are about shutting down
	cd.NewShutdownEvent()

	// STUB: The only value here is logging the specific signal.
	/*
		switch sig {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			cd.EvtQ.NewShutdownEvent()
		case syscall.SIGXCPU:

		}
	*/
}
