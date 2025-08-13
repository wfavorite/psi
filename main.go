package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	// STUB: Consider creating and calling on an args struct.
	// STUB: Or something.

	cpuPSI, err := ReadPressureFile("/proc/pressure/cpu")

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to read cpu - %s\n", err.Error())
		os.Exit(1)
	}

	jdata, err := json.MarshalIndent(cpuPSI, "", "  ")

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to marshall data - %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(string(jdata))
	os.Exit(0)
}

/* ======================================================================== */

func HandleAbout() {
	fmt.Println("psi - Pressure reporter")

}

/* ======================================================================== */

func HandleUsage() {
	fmt.Println("psi - Pressure reporter")
	fmt.Println("  Usage: psi <options> <int>")
	fmt.Println("  Options:")
	fmt.Println("    -a     Show about information")
	fmt.Println("    -h     Show this usage information")
	fmt.Println("    -j     Dump current stats in JSON (incompatible with other options or interval)")
	fmt.Println("    <int>  Print tabular stats on interval")
}
