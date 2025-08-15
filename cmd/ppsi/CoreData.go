package main

/* ------------------------------------------------------------------------ */

type CoreData struct {
	EvtQ *EventQ
	List *Listener
	Cash *ClientCache
	Sigs *Signals
	Logr *Logger
}

/* ======================================================================== */

func NewCoreData() (cd *CoreData) {
	cd = new(CoreData)
	return
}
