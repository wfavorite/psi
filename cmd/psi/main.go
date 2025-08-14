// Package main implements the psi utility.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"psi/pkg/psi"
)

/* ======================================================================== */

func main() {

	var cmdl *CmdLine
	if cl, clErr := ParseCommandLine(); clErr == nil {
		cmdl = cl
	} else {
		fmt.Fprintln(os.Stderr, "ERROR:", clErr.Error())
		os.Exit(1)
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

	if cmdl.DbgRandData {
		// This is a totally non-standard & hidden option for testing ANSI
		// colour output.
		sio.RandomValues = true
	}

	si := sio.NewStallInfo()

	if err := si.Collect(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to read psi data\n       %s\n", err.Error())
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
				fmt.Fprintf(os.Stderr, "ERROR: Failed to read psi data\n       %s\n", err.Error())
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
	fmt.Println("  Version:", VersionString)
	fmt.Println("  Paul Reynaud <inmate@itter.castle.at>")
	fmt.Println("  William Favorite <wfavorit@gmail.com>")

}

/* ======================================================================== */

// HandleUsage prints the standard help/usage information to stdout.
func HandleUsage() {
	fmt.Println("psi - Pressure reporter")
	fmt.Println("  Usage: psi <options> <int>")
	fmt.Println("  Options:")
	fmt.Println("    -a     Show about information.")
	fmt.Println("    -h     Show this usage information.")
	fmt.Println("    -j     Dump current stats as a JSON structure. This option is incompatible")
	fmt.Println("           with other options or interval printing.")
	fmt.Println("    -m     Print output in monochrome. Default is ANSI colour.")
	fmt.Println("    -t     Print a timestamp on each line of tabular output.")
	fmt.Println("    -w     Print in a wide format (potentially beyond 80 chars).")
	fmt.Println("    <int>  Print tabular stats on interval. The supplied time is assumed either")
	fmt.Println("           seconds (if an integer) or a Golang duration (eg: 500ms, 1s, 2m).")
}
