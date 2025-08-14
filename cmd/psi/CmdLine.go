package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

/* ------------------------------------------------------------------------ */

// CmdLine is a structure that captures the realm of possible command line
// arguments.
type CmdLine struct {
	OptAbout    bool
	OptUsage    bool
	OptJSON     bool
	OptWide     bool
	OptTMStamp  bool
	OptMono     bool
	ArgInterval time.Duration
	Error       string
}

/* ======================================================================== */

// ParseCommandLine is the public interface to the parseCommandLine (testable)
// interface.
func ParseCommandLine() (cmdl *CmdLine, err error) {

	return parseCommandLine(os.Args)

}

/* ======================================================================== */

// parseCommandLine is the actual (testable) command line parser. The o
func parseCommandLine(args []string) (cmdl *CmdLine, err error) {

	cmdl = new(CmdLine)

	fs := flag.NewFlagSet("psi", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Usage = func() {}

	fs.BoolVar(&cmdl.OptAbout, "a", false, "")
	fs.BoolVar(&cmdl.OptUsage, "h", false, "")
	fs.BoolVar(&cmdl.OptJSON, "j", false, "")
	fs.BoolVar(&cmdl.OptWide, "w", false, "")
	fs.BoolVar(&cmdl.OptTMStamp, "t", false, "")
	fs.BoolVar(&cmdl.OptMono, "m", false, "")

	err = fs.Parse(args[1:])

	if err != nil {
		cmdl.Error = err.Error()
		return
	}

	remaining := fs.Args()

	switch len(remaining) {
	case 0:
	case 1:
		// First, attempt to parse as an integer (seconds).
		if i64, ipErr := strconv.ParseInt(remaining[0], 10, 32); ipErr == nil {
			// This is not really a proper Duration --------+
			// but multiplying it by a Duration fixes that. |
			//                                              V
			cmdl.ArgInterval = time.Second * time.Duration(i64)
		} else {
			// ...if that fails, parse as a Duration.

			if dur, dpErr := time.ParseDuration(remaining[0]); dpErr == nil {
				cmdl.ArgInterval = dur
			} else {
				cmdl.Error = "The interval was not parsable"
				return
			}

		}
	default:
		cmdl.Error = "Extra (non-flag) arguments not understood"
		return
	}

	err = cmdl.ValidateArgs()

	return
}

/* ======================================================================== */

// ValidateArgs is used to check for invalid combinations of command line
// arguments.
func (cmdl *CmdLine) ValidateArgs() (err error) {

	if cmdl == nil {
		// This is essentially an assert()ion.
		err = fmt.Errorf("nil CmdLine passed to ValidateArgs")
		return
	}

	if len(cmdl.Error) > 0 {
		// No point in processing if we already have an error.
		// Note: This should not be set as we arrive here. STUB

		return
	}

	if cmdl.OptAbout {
		if cmdl.OptUsage || cmdl.OptJSON || cmdl.OptWide || cmdl.ArgInterval > 0 {
			cmdl.Error = "The -a option is mutually exclusive of all other options"
			return
		}
	}

	if cmdl.OptUsage {
		if cmdl.OptJSON || cmdl.OptWide || cmdl.ArgInterval > 0 {
			cmdl.Error = "The -h option is mutually exclusive of all other options"
			return
		}
	}

	if cmdl.OptJSON {
		if cmdl.OptWide || cmdl.ArgInterval > 0 {
			cmdl.Error = "The -j option is mutually exclusive of all other options"
			return
		}
	}

	return
}
