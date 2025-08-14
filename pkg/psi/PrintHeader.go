package psi

import (
	"fmt"
	"io"
)

/*
	About printing

	Normally this would not be a thing for a pure class, but... I think it
	is appropriate as this data is highly likely to be used in some sort of
	logging collector or console use, this class knows the data best, and
	finally these are clean as methods - which must be in the class.
*/

/* ======================================================================== */

// PrintHeader prints a header line for the output data.
func (si *StallInfo) PrintHeader(w io.Writer) {

	if si.opt.Width == Wide {
		si.headerWide(w)
	} else {
		si.headerCondensed(w)
	}
}

/* ======================================================================== */

// headerWide prints three lines of header info.
func (si *StallInfo) headerWide(w io.Writer) {

	// ===== First line =====

	fmt.Fprint(w, "#")

	if si.opt.TimeStamp {
		fmt.Fprint(w, "             ")
	}

	if si.CPU != nil {
		fmt.Fprintf(w, " %-8s                                 ", "CPU")
	}

	if si.CPU != nil {
		fmt.Fprintf(w, " %-8s                                 ", "IO")
	}

	if si.IRQ != nil {
		fmt.Fprintf(w, " %-8s                                 ", "IRQ")
	}

	if si.Mem != nil {
		fmt.Fprintf(w, " %-8s                                 ", "Memory")
	}

	// EOL
	fmt.Fprintln(w)

	// ===== Second line ====

	fmt.Fprint(w, "#")

	if si.opt.TimeStamp {
		fmt.Fprint(w, "             ")
	}

	if si.CPU != nil {
		fmt.Fprint(w, " Some                 Full                ")
	}

	if si.IO != nil {
		fmt.Fprint(w, " Some                 Full                ")
	}

	if si.IRQ != nil {
		fmt.Fprint(w, " Some                 Full                ")
	}

	if si.Mem != nil {
		fmt.Fprint(w, " Some                 Full                ")
	}

	// EOL
	fmt.Fprintln(w)

	// ===== Third line =====

	fmt.Fprint(w, "#")

	if si.opt.TimeStamp {
		fmt.Fprint(w, " Timestamp   ")
	}

	if si.CPU != nil {
		fmt.Fprintf(w, " %-6s %-6s %-6s %-6s %-6s %-6s", "avg10", "avg60", "avg300", "avg10", "avg60", "avg300")
	}

	if si.IO != nil {
		fmt.Fprintf(w, " %-6s %-6s %-6s %-6s %-6s %-6s", "avg10", "avg60", "avg300", "avg10", "avg60", "avg300")
	}

	if si.IRQ != nil {
		fmt.Fprintf(w, " %-6s %-6s %-6s %-6s %-6s %-6s", "avg10", "avg60", "avg300", "avg10", "avg60", "avg300")
	}

	if si.Mem != nil {
		fmt.Fprintf(w, " %-6s %-6s %-6s %-6s %-6s %-6s", "avg10", "avg60", "avg300", "avg10", "avg60", "avg300")
	}

	// EOL
	fmt.Fprintln(w)

}

/* ======================================================================== */

func (si *StallInfo) headerCondensed(w io.Writer) {
	// ===== First line =====

	fmt.Fprint(w, "#")

	if si.opt.TimeStamp {
		fmt.Fprint(w, "             ")
	}

	fmt.Fprint(w, "     ")

	if si.CPU != nil {
		fmt.Fprintf(w, " %-8s   ", "CPU")
	}

	if si.CPU != nil {
		fmt.Fprintf(w, " %-8s   ", "IO")
	}

	if si.IRQ != nil {
		fmt.Fprintf(w, " %-8s   ", "IRQ")
	}

	if si.Mem != nil {
		fmt.Fprintf(w, " %-8s   ", "Memory")
	}

	// EOL
	fmt.Fprintln(w)

	// ===== Second line ====

	fmt.Fprint(w, "#")

	if si.opt.TimeStamp {
		fmt.Fprint(w, " Timestamp   ")
	}

	twoLineFmt := " %3s %3s %3s"

	fmt.Fprint(w, "     ")

	if si.CPU != nil {
		fmt.Fprintf(w, twoLineFmt, "10s", "1m", "5m")
	}

	if si.IO != nil {
		fmt.Fprintf(w, twoLineFmt, "10s", "1m", "5m")
	}

	if si.IRQ != nil {
		fmt.Fprintf(w, twoLineFmt, "10s", "1m", "5m")
	}

	if si.Mem != nil {
		fmt.Fprintf(w, twoLineFmt, "10s", "1m", "5m")
	}

	// EOL
	fmt.Fprintln(w)
}
