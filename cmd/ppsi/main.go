package main

import (
	"fmt"
	"os"

	"github.com/wfavorite/initq"
)

func main() {

	defer fmt.Println("Exiting main()")

	cd := NewCoreData()

	iq := initq.NewInitQ()

	iq.Add("logger", cd.StartLogger)
	iq.Add("cache", cd.ClientCache)
	iq.Add("eventq", cd.StartEventQ)
	iq.Add("listener", cd.StartListener)
	iq.Add("signals", cd.RegisterSignals)

	// STUB: Check the error here.
	iq.Process()

	// Now fall into the main event loop.

	for {
		select {
		case sig := <-cd.Sigs.S:
			cd.HandleSignal(sig)
		case evt, ok := <-cd.EvtQ.Q:

			if ok {
				evt.Handle()
			} else {
				// When the event Q is shutdown, then we are done.
				os.Exit(0)
			}
		}
	}
}
