package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* ------------------------------------------------------------------------ */

type PressureLine struct {
	Avg10  float64
	Avg60  float64
	Avg300 float64
	Total  int64
}

/* ------------------------------------------------------------------------ */

type Pressure struct {
	Some PressureLine
	Full PressureLine
}

/* ------------------------------------------------------------------------ */

// STUB: Rename this.

type Interval int

/* ------------------------------------------------------------------------ */

const (
	AvgUnk Interval = iota
	Avg10
	Avg60
	Avg300
)

/* ======================================================================== */

func (interval Interval) String() (str string) {

	switch interval {
	case Avg10:
		str = "avg10"
	case Avg60:
		str = "avg60"
	case Avg300:
		str = "avg300"
	default:
		str = "avgUnk"
	}

	return
}

/* ======================================================================== */

func ReadPressureFile(filename string) (psi *Pressure, err error) {

	var f *os.File

	f, err = os.Open(filename)

	if err != nil {
		return
	}

	defer f.Close()

	// Create a candidate structure
	psiCandidate := new(Pressure)

	s := bufio.NewScanner(f)

	for s.Scan() {
		// Pull a line
		line := s.Text()
		// Split into parts
		parts := strings.Fields(line)

		if len(parts) != 5 {
			err = fmt.Errorf("unexpected line field count")
			return
		}

		lineType := parts[0]

		var pLine *PressureLine
		switch lineType {
		case "some":
			pLine = &psiCandidate.Some
		case "full":
			pLine = &psiCandidate.Full

		default:
			err = fmt.Errorf("unknown line type")
			return
		}

		var interval Interval
		var value float64

		// avg10
		interval, value, err = parseAvgEntry(parts[1])

		if err != nil {
			return
		}

		if interval != Avg10 {
			err = fmt.Errorf("first item not avg10")
			return
		}

		pLine.Avg10 = value

		// avg60
		interval, value, err = parseAvgEntry(parts[2])

		if err != nil {
			return
		}

		if interval != Avg60 {
			err = fmt.Errorf("second item not avg60")
		}

		pLine.Avg60 = value

		// avg300
		interval, value, err = parseAvgEntry(parts[3])

		if err != nil {
			return
		}

		if interval != Avg300 {
			err = fmt.Errorf("third item not avg300")
		}

		pLine.Avg300 = value

	}

	// All good, no errors so candidate becomes the return value.
	psi = psiCandidate

	return
}

/* ======================================================================== */

// parseAvgEntry converts one of the three known avgNN entries into its
// constituent parts.
func parseAvgEntry(entry string) (interval Interval, value float64, err error) {

	if len(entry) < 10 {
		err = fmt.Errorf("impossibly short average entry")
		return

	}

	switch {
	case strings.HasPrefix(entry, "avg10"):
		interval = Avg10
	case strings.HasPrefix(entry, "avg60"):
		interval = Avg60
	case strings.HasPrefix(entry, "avg300"):
		interval = Avg300
	default:
		err = fmt.Errorf("mislabeled avg entry")
		return
	}

	parts := strings.Split(entry, "=")

	if len(parts) != 2 {
		err = fmt.Errorf("invalid avg split")
		return
	}

	value, err = strconv.ParseFloat(parts[1], 64)

	if err != nil {
		return
	}

	if value < 0 || value > 100 {
		err = fmt.Errorf("parsed value out of bounds")
	}

	return
}
