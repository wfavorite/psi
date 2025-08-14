package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

/* ------------------------------------------------------------------------ */

// CmdLine is a structure that captures the realm of possible command line
// arguments.
type CmdLine struct {
	OptAbout    bool // -a
	OptUsage    bool // -h
	OptJSON     bool // -j
	OptMono     bool // -m
	DbgRandData bool // -r (hidden)
	OptTMStamp  bool // -t
	OptWide     bool // -w
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
	fs.SetOutput(io.Discard) // Don't write to anything
	fs.Usage = func() {}     // If asked, do a no-op

	fs.BoolVar(&cmdl.OptAbout, "a", false, "")    // -a  About
	fs.BoolVar(&cmdl.OptUsage, "h", false, "")    // -h  Usage
	fs.BoolVar(&cmdl.OptJSON, "j", false, "")     // -h  Dump a JSON structure
	fs.BoolVar(&cmdl.OptMono, "m", false, "")     // -m  Write tabular output in monochrome
	fs.BoolVar(&cmdl.DbgRandData, "r", false, "") // -r  (Hidden) option to run with random data
	fs.BoolVar(&cmdl.OptTMStamp, "t", false, "")  // -t  Include timestamp in tabular output
	fs.BoolVar(&cmdl.OptWide, "w", false, "")     // -w  Write wide tabular output

	err = fs.Parse(args[1:])

	// This is most likely a bad option message.
	if err != nil {
		cmdl.Error = err.Error()
		return
	}

	// A list of what is left. There should be *max* one.
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
				cmdl.Error = fmt.Sprintf("The interval \"%s\" was not parsable", remaining[0])
				return
			}

		}
	default:
		cmdl.Error = "Extra (non-flag) arguments not understood"
		return
	}

	// This is the assert case. The only thing that returns an error here
	// would have been caught by now. So... this is just being pedantic.
	if assertCase := cmdl.ValidateArgs(); assertCase != nil {
		log.Fatal(assertCase.Error())
	}

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
		return
	}

	if cmdl.OptAbout {
		if cmdl.OptUsage ||
			cmdl.OptJSON ||
			cmdl.OptMono ||
			cmdl.OptTMStamp ||
			cmdl.OptWide ||
			cmdl.ArgInterval > 0 {
			cmdl.Error = "The -a option is mutually exclusive of all other options"
			return
		}
	}

	if cmdl.OptUsage {
		if cmdl.OptJSON ||
			cmdl.OptMono ||
			cmdl.OptTMStamp ||
			cmdl.OptWide ||
			cmdl.ArgInterval > 0 {
			cmdl.Error = "The -h option is mutually exclusive of all other options"
			return
		}
	}

	if cmdl.OptJSON {
		if cmdl.OptMono ||
			cmdl.OptTMStamp ||
			cmdl.OptWide ||
			cmdl.ArgInterval > 0 {
			cmdl.Error = "The -j option is mutually exclusive of all other options"
			return
		}
	}

	if cmdl.ArgInterval < 0 {
		cmdl.Error = "The supplied interval is out of bounds (negative)"
	}

	return
}
