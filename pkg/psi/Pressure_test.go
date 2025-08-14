package psi

import "testing"

/* ======================================================================== */

func TestAvgEntry(t *testing.T) {

	var interval Interval
	var value float64

	var err error
	var expected float64

	// -----
	expected = 0.0
	interval, value, err = parseAvgEntry("avg10=0.00")

	if err != nil {
		t.Errorf("Errored on valid pattern - %s", err.Error())
	}

	if interval != Avg10 {
		t.Errorf("Parsed %s; expected %s", interval, Avg10)
	}

	if value != expected {
		t.Errorf("Parsed %f; expected %f", value, expected)
	}

	// -----
	expected = 3.14
	interval, value, err = parseAvgEntry("avg10=3.14")

	if err != nil {
		t.Errorf("Errored on valid pattern - %s", err.Error())
	}

	if interval != Avg10 {
		t.Errorf("Parsed %s; expected %s", interval, Avg10)
	}

	if value != expected {
		t.Errorf("Parsed %f; expected %f", value, expected)
	}

	// -----
	expected = 100.0
	interval, value, err = parseAvgEntry("avg10=100.00")

	if err != nil {
		t.Errorf("Errored on valid pattern - %s", err.Error())
	}

	if interval != Avg10 {
		t.Errorf("Parsed %s; expected %s", interval, Avg10)
	}

	if value != expected {
		t.Errorf("Parsed %f; expected %f", value, expected)
	}

	// -----
	expected = 99.99
	interval, value, err = parseAvgEntry("avg300=99.99")

	if err != nil {
		t.Errorf("Errored on valid pattern - %s", err.Error())
	}

	if interval != Avg300 {
		t.Errorf("Parsed %s; expected %s", interval, Avg10)
	}

	if value != expected {
		t.Errorf("Parsed %f; expected %f", value, expected)
	}

	// -----
	expected = 0.00
	interval, value, err = parseAvgEntry("avg300=-1.00")

	if err == nil {
		t.Errorf("Failed to to error on out of bounds condition.")
	}

	// -----
	expected = 0.00
	interval, value, err = parseAvgEntry("avg300=100.01")

	if err == nil {
		t.Errorf("Failed to to error on out of bounds condition.")
	}

}

/* ======================================================================== */

func TestParsePSIFile(t *testing.T) {

	var psi *PressureFile
	var err error

	psi = NewPressureFile()

	err = psi.ReadFromFile("../../test/cpu.1")

	if err == nil {
		if psi != nil {
			// All as expected so far

			if psi.Some.Avg10 != 0.0 {
				t.Errorf("some avg10 parsed as %f, expected %f", psi.Some.Avg10, 0.00)
			}

		} else {
			t.Errorf("No error, but return value is nil")
		}
	} else {
		t.Errorf("Error not expected - %s", err.Error())
	}

}
