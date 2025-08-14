package main

import "testing"

/* ======================================================================== */

func TestCmdLineParse(t *testing.T) {

	var cmdl *CmdLine
	var err error

	cmdl, err = parseCommandLine([]string{"cmd", "-a"})

	if err == nil {
		if !cmdl.OptAbout {
			t.Errorf("The OptAbout flag was not detected")
		}

		if cmdl.OptUsage {
			t.Errorf("The OptUsage flas was set (and should not be)")
		}
	}

}
