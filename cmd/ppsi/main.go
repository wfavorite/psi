// Package main implements the event handling 'server' side of the PSI
// collector.
package main

import (
	"fmt"
	"os"

	"github.com/wfavorite/initq"
)

/* ======================================================================== */

func main() {

	cd := NewCoreData()

	iq := initq.NewInitQ()

	iq.Add("logger", cd.StartLogger)
	iq.Add("cache", cd.ClientCache)
	iq.Add("eventq", cd.StartEventQ)
	iq.Add("listener", cd.StartListener)
	iq.Add("signals", cd.RegisterSignals)
	iq.Add("claunch", cd.ClientLaunch)

	if err := iq.Process(); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err.Error())
		os.Exit(1)
	}

	cd.Logr.Normal.Printf("ppsi (PID=%d) started.", os.Getpid())

	// Now fall into the main event loop.
	for {
		select {
		case sig := <-cd.Sigs.s:
			cd.HandleSignal(sig)
		case evt, ok := <-cd.EvtQ.q:
			if ok {
				evt.Handle()
			} else {
				// When the event Q is shutdown, then we are done.
				os.Exit(0)
			}
		}
	}
}
