// Package main implements the psi utility.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/wfavorite/psi/pkg/psi"
)

/* ======================================================================== */

func main() {

	var cmdl *CmdLine
	if cl, clErr := ParseCommandLine(); clErr != nil {
		fmt.Fprintln(os.Stderr, "ASSERT:", clErr.Error())
		os.Exit(1)
	} else {
		cmdl = cl
	}

	if len(cmdl.Error) > 0 {
		fmt.Fprintln(os.Stderr, "ERROR:", cmdl.Error)
		os.Exit(1)
	}

	if cmdl.OptAbout {
		HandleAbout()
		os.Exit(0)
	}

	if cmdl.OptUsage {
		HandleUsage()
		os.Exit(0)
	}

	sio := psi.NewOptions()

	if cmdl.OptWide {
		sio.Width = psi.Wide
	}

	if cmdl.OptTMStamp {
		sio.TimeStamp = true
	}

	// STUB: This is for debuggery.
	sio.RandomValues = true

	si := sio.NewStallInfo()

	if err := si.Collect(); err != nil {
		// STUB: This error message is incorrect.
		fmt.Fprintf(os.Stderr, "ERROR: Failed to read cpu - %s\n", err.Error())
		os.Exit(1)
	}

	if cmdl.OptJSON {
		jdata, err := json.MarshalIndent(si, "", "  ")

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to marshall data - %s\n", err.Error())
			os.Exit(1)
		}

		fmt.Println(string(jdata))
		os.Exit(0)
	}

	si.PrintHeader(os.Stdout)
	si.PrintLine(os.Stdout)

	if cmdl.ArgInterval != 0 {

		for {
			time.Sleep(cmdl.ArgInterval)

			if err := si.Collect(); err != nil {
				// STUB: This error message is incorrect.
				fmt.Fprintf(os.Stderr, "ERROR: Failed to read cpu - %s\n", err.Error())
				os.Exit(1)
			}

			si.PrintLine(os.Stdout)
		}
	}

	os.Exit(0)
}

/* ======================================================================== */

// HandleAbout prints the standard about information to stdout.
func HandleAbout() {
	fmt.Println("psi - Pressure reporter")

}

/* ======================================================================== */

// HandleUsage prints the standard help/usage information to stdout.
func HandleUsage() {
	fmt.Println("psi - Pressure reporter")
	fmt.Println("  Usage: psi <options> <int>")
	fmt.Println("  Options:")
	fmt.Println("    -a     Show about information")
	fmt.Println("    -h     Show this usage information")
	fmt.Println("    -j     Dump current stats in JSON (incompatible with other options or interval)")
	fmt.Println("    <int>  Print tabular stats on interval")
}
