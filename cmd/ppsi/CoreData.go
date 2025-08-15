package main

/* ------------------------------------------------------------------------ */

// CoreData is the central data structure for all daemon needs.
type CoreData struct {
	EvtQ *EventQ       `json:"event_q"`
	List *Listener     `json:"-"`
	Cash *ClientCache  `json:"active_clients"`
	Sigs *Signals      `json:"signals"`
	Logr *Logger       `json:"-"`
	Clil *ClientLaunch `json:"launched_clients"`
}

/* ======================================================================== */

// NewCoreData creates and initializes a new CoreData structure.
func NewCoreData() (cd *CoreData) {
	cd = new(CoreData)
	return
}
