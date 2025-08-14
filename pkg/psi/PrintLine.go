package psi

import (
	"fmt"
	"io"
)

/* ======================================================================== */

func (si *StallInfo) PrintLine(w io.Writer) {
	if si.opt.Width == Wide {
		si.lineWide(w)
	} else {
		si.lineCondensed(w)
	}
}

/* ======================================================================== */

func (si *StallInfo) lineWide(w io.Writer) {

	fmt.Fprint(w, "#")

	if si.opt.TimeStamp {
		fmt.Fprintf(w, " %-12s", si.Timestamp)
	}

	si.CPU.dataWide(w, si.opt)
	si.IO.dataWide(w, si.opt)
	si.IRQ.dataWide(w, si.opt)
	si.Mem.dataWide(w, si.opt)

	// EOL
	fmt.Fprintln(w)

}

/* ======================================================================== */

func (pf *PressureFile) dataWide(w io.Writer, o *Options) {

	if pf == nil {
		return
	}

	if o.Monochrome {

		oneLineFmt := " %-6.2f %-6.2f %-6.2f  %-6.2f %-6.2f %-6.2f"

		fmt.Fprintf(w, oneLineFmt,
			pf.Some.Avg10,
			pf.Some.Avg60,
			pf.Some.Avg300,
			pf.Full.Avg10,
			pf.Full.Avg60,
			pf.Full.Avg300,
		)
	} else {
		oneLineFmt := " %s%-6.2f\033[0m %s%-6.2f\033[0m %s%-6.2f\033[0m"

		c10 := deriveColour(o, pf.Some.Avg10)
		c60 := deriveColour(o, pf.Some.Avg60)
		c300 := deriveColour(o, pf.Some.Avg300)

		// Two parts - because the arg list is just so long

		fmt.Fprintf(w, oneLineFmt,
			c10,
			pf.Some.Avg10,
			c60,
			pf.Some.Avg60,
			c300,
			pf.Some.Avg300,
		)

		c10 = deriveColour(o, pf.Full.Avg10)
		c60 = deriveColour(o, pf.Full.Avg60)
		c300 = deriveColour(o, pf.Full.Avg300)

		fmt.Fprintf(w, oneLineFmt,
			c10,
			pf.Full.Avg10,
			c60,
			pf.Full.Avg60,
			c300,
			pf.Full.Avg300,
		)

	}
}

const Some bool = true
const Full bool = false

/* ======================================================================== */

func (si *StallInfo) lineCondensed(w io.Writer) {

	fmt.Fprint(w, "#")
	if si.opt.TimeStamp {
		fmt.Fprintf(w, " %-12s", si.Timestamp)
	}
	fmt.Fprint(w, " Some")
	si.CPU.dataCondensed(w, Some, si.opt)
	si.IO.dataCondensed(w, Some, si.opt)
	si.IRQ.dataCondensed(w, Some, si.opt)
	si.Mem.dataCondensed(w, Some, si.opt)
	fmt.Fprintln(w)

	fmt.Fprint(w, "#")
	if si.opt.TimeStamp {
		fmt.Fprintf(w, " %-12s", si.Timestamp)
	}

	fmt.Fprint(w, " Full")
	si.CPU.dataCondensed(w, Full, si.opt)
	si.IO.dataCondensed(w, Full, si.opt)
	si.IRQ.dataCondensed(w, Full, si.opt)
	si.Mem.dataCondensed(w, Full, si.opt)
	fmt.Fprintln(w)

}

/* ======================================================================== */

func (pf *PressureFile) dataCondensed(w io.Writer, some bool, o *Options) {

	if pf == nil {
		return
	}

	if o.Monochrome {
		twoLineFmt := " %3.0f %3.0f %3.0f"

		if some {
			fmt.Fprintf(w, twoLineFmt,
				pf.Some.Avg10,
				pf.Some.Avg60,
				pf.Some.Avg300,
			)
		} else {
			fmt.Fprintf(w, twoLineFmt,
				pf.Full.Avg10,
				pf.Full.Avg60,
				pf.Full.Avg300,
			)
		}
	} else {
		twoLineFmt := " %s%3.0f\033[0m %s%3.0f\033[0m %s%3.0f\033[0m"

		if some {
			c10 := deriveColour(o, pf.Some.Avg10)
			c60 := deriveColour(o, pf.Some.Avg60)
			c300 := deriveColour(o, pf.Some.Avg300)

			fmt.Fprintf(w, twoLineFmt,
				c10,
				pf.Some.Avg10,
				c60,
				pf.Some.Avg60,
				c300,
				pf.Some.Avg300,
			)
		} else {
			c10 := deriveColour(o, pf.Full.Avg10)
			c60 := deriveColour(o, pf.Full.Avg60)
			c300 := deriveColour(o, pf.Full.Avg300)

			fmt.Fprintf(w, twoLineFmt,
				c10,
				pf.Full.Avg10,
				c60,
				pf.Full.Avg60,
				c300,
				pf.Full.Avg300,
			)
		}
	}

}

/* ======================================================================== */

func deriveColour(o *Options, f float64) (c string) {

	if f > o.RedThreshold {
		// Red
		//c = "\033[31m"
		// Bright red
		c = "\033[91m"
	} else {
		if f > o.YellowThreshold {
			// Yellow
			//c = "\033[33m"
			// Bright Yellow
			c = "\033[93m"
		} else {
			// Green
			//c = "\033[32m"
			// Bright Green
			c = "\033[92m"
		}
	}

	return
}
