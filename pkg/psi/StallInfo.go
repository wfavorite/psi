// Package psi is used to pull information from the Pressure Stall Info data
// sources provided in /proc/pressure.
package psi

/* ------------------------------------------------------------------------ */

// StallInfo models all pressure info found in the Linux /proc filesystem.
type StallInfo struct {
	CPU    *Pressure
	IO     *Pressure
	IRQ    *Pressure
	Memory *Pressure

	opt *Options
}

/* ======================================================================== */

// NewStallInfo creates a new StallInfo instance based on the supplied Options
// structure.
func (opt *Options) NewStallInfo() (si *StallInfo) {

	si = new(StallInfo)

	// Gracefully handle the method-on-nil problem. (Default to standard
	// behaviours.)
	if opt == nil {
		opt = NewOptions()
	}

	si.opt = opt

	return
}

/* ======================================================================== */

// NewStallInfo creates a new StallInfo instance based on default arguments.
func NewStallInfo() (si *StallInfo) {

	si = new(StallInfo)
	// This is the case were Options are not 'passed'. We rely upon defaults.
	si.opt = NewOptions()

	return
}

/* ======================================================================== */

// Collect is called upon a StallInfo instance to update the data.
func (si *StallInfo) Collect() (err error) {

	if si == nil {
		// Don't panic, just error.
		err = ErrMethodOnNil
		return
	}

	return
}
