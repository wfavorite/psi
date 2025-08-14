// Package psi is used to pull information from the Pressure Stall Info data
// sources provided in /proc/pressure.
package psi

import (
	"path/filepath"
	"time"
)

/* ------------------------------------------------------------------------ */

// StallInfo models all pressure info found in the Linux /proc filesystem.
type StallInfo struct {
	CPU       *PressureFile `json:"cpu,omitempty"`
	IO        *PressureFile `json:"io,omitempty"`
	IRQ       *PressureFile `json:"irq,omitempty"`
	Mem       *PressureFile `json:"memory,omitempty"`
	Timestamp string        `json:"time"`

	opt *Options

	// These are cached values so we do not continually re-construct otherwise
	// static strings.
	cpuFile string
	ioFile  string
	irqFile string
	memFile string
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

	if si.opt == nil {
		// This should not be a thing. But if it is... handle it gracefully by
		// falling back to default behaviour. About this choice:
		// - It (potentially) hides internal problems. Perhaps there are ways
		//   of capturing this in test.
		// - Falling back to default is a low cost / easy alternative to
		//   punishing the caller.
		si.opt = NewOptions()
	}

	// Grab a timestamp regardless (if we know we are gonna use it).
	si.Timestamp = time.Now().Local().Format("15:04:05.000")

	// Walk through the possible collections.
	// These are standard patterns - all relevant notes are on the first/CPU.

	if si.opt.Collect&CollectCPU == CollectCPU {
		// This may be our first...
		if si.CPU == nil {
			// ...so initialize.
			si.CPU = NewPressureFile()
			si.cpuFile = filepath.Join(si.opt.SourceDirectory, "cpu")
		}

		if si.opt.RandomValues {
			si.CPU.DebugRandom()
		} else {
			err = si.CPU.ReadFromFile(si.cpuFile)
		}

		if err != nil {
			// These are standard things. If we cannot read them, the data
			// cannot be trusted so the best alternative is to fail.
			return
		}

	}

	if si.opt.Collect&CollectIO == CollectIO {

		if si.IO == nil {
			si.IO = NewPressureFile()
			si.ioFile = filepath.Join(si.opt.SourceDirectory, "io")
		}

		if si.opt.RandomValues {
			si.IO.DebugRandom()
		} else {
			err = si.IO.ReadFromFile(si.ioFile)
		}

		if err != nil {
			return
		}

	}

	if si.opt.Collect&CollectIRQ == CollectIRQ {

		if si.IRQ == nil {
			si.IRQ = NewPressureFile()
			si.irqFile = filepath.Join(si.opt.SourceDirectory, "irq")
		}

		if si.opt.RandomValues {
			si.IRQ.DebugRandom()
		} else {
			err = si.IRQ.ReadFromFile(si.irqFile)
		}

		if err != nil {
			return
		}

	}

	if si.opt.Collect&CollectMem == CollectMem {

		if si.Mem == nil {
			si.Mem = NewPressureFile()
			si.memFile = filepath.Join(si.opt.SourceDirectory, "memory")
		}

		if si.opt.RandomValues {
			si.Mem.DebugRandom()
		} else {
			err = si.Mem.ReadFromFile(si.memFile)
		}

		if err != nil {
			return
		}

	}

	return
}
