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
			t.Errorf("The OptUsage flag was set (and should not be)")
		}

		if cmdl.OptJSON {
			t.Errorf("The OptJSON flag was set (and should not be)")
		}
	}

	// Things that don't work with -a.

	failWithA := []string{"-h", "-j", "-m", "-t", "-w", "2m", "3"}

	for _, badOpt := range failWithA {
		cmdl, err = parseCommandLine([]string{"cmd", "-a", badOpt})

		if len(cmdl.Error) == 0 {
			t.Errorf("The %s option was not considered an error when with -a", badOpt)
		}
	}

	// Things that don't work with -h.

	failWithH := []string{"-a", "-j", "-m", "-t", "-w", "2m", "3"}

	for _, badOpt := range failWithH {
		cmdl, err = parseCommandLine([]string{"cmd", "-h", badOpt})

		if len(cmdl.Error) == 0 {
			t.Errorf("The %s option was not considered an error when with -h", badOpt)
		}
	}

	// Things that don't work with -j.

	failWithJ := []string{"-a", "-h", "-m", "-t", "-w", "2m", "3"}

	for _, badOpt := range failWithJ {
		cmdl, err = parseCommandLine([]string{"cmd", "-j", badOpt})

		if len(cmdl.Error) == 0 {
			t.Errorf("The %s option was not considered an error when with -j", badOpt)
		}
	}

}
