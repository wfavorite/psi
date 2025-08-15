package psi

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

/* ------------------------------------------------------------------------ */

// PressureLine is one of the two lines in a standard pressure file.
type PressureLine struct {
	Avg10  float64
	Avg60  float64
	Avg300 float64
	Total  int64
}

/* ------------------------------------------------------------------------ */

// PressureFile models all the data in a pressure file.
type PressureFile struct {
	Some PressureLine
	Full PressureLine
}

/* ------------------------------------------------------------------------ */

// Interval is the type backing the enum for which avg entry was parsed.
type Interval int

/* ------------------------------------------------------------------------ */

const (
	// AvgUnk is the uninitialized or error case when parsing an avgNN entry.
	AvgUnk Interval = iota

	// Avg10 is the first average entry on a pressure file line.
	Avg10

	// Avg60 is the second average entry on a pressure file line.
	Avg60

	// Avg300 is the third average entry on a pressure file line.
	Avg300
)

/* ======================================================================== */

// String satisfies the Stringer interface for the Interval type.
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

// NewPressureFile is used to create a PressureFile structure that can be
// used to read data from a file.
func NewPressureFile() (psi *PressureFile) {

	psi = new(PressureFile)
	return
}

/* ======================================================================== */

// ReadFromFile reads the PSI pressure data from the supplied filename. The
// supplied filename should be one of the files found in /proc/pressure.
func (psi *PressureFile) ReadFromFile(filename string) (err error) {

	var f *os.File

	f, err = os.Open(filename)

	if err != nil {
		return
	}

	defer f.Close()

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
			pLine = &psi.Some
		case "full":
			pLine = &psi.Full

		default:
			err = fmt.Errorf("unknown line type")
			return
		}

		var interval Interval
		var fValue float64
		var iValue int64

		// avg10
		interval, fValue, err = parseAvgEntry(parts[1])

		if err != nil {
			return
		}

		if interval != Avg10 {
			err = fmt.Errorf("first item not avg10")
			return
		}

		pLine.Avg10 = fValue

		// avg60
		interval, fValue, err = parseAvgEntry(parts[2])

		if err != nil {
			return
		}

		if interval != Avg60 {
			err = fmt.Errorf("second item not avg60")
		}

		pLine.Avg60 = fValue

		// avg300
		interval, fValue, err = parseAvgEntry(parts[3])

		if err != nil {
			return
		}

		if interval != Avg300 {
			err = fmt.Errorf("third item not avg300")
		}

		pLine.Avg300 = fValue

		// total
		iValue, err = parseTotalEntry(parts[4])

		if err != nil {
			return
		}

		pLine.Total = iValue

	}

	return
}

/* ======================================================================== */

// DebugRandom generates random values for all data collected.
func (psi *PressureFile) DebugRandom() {
	psi.Some.Avg10 = rand.Float64() * float64(100)
	psi.Some.Avg60 = rand.Float64() * float64(100)
	psi.Some.Avg300 = rand.Float64() * float64(100)
	psi.Full.Avg10 = rand.Float64() * float64(100)
	psi.Full.Avg60 = rand.Float64() * float64(100)
	psi.Full.Avg300 = rand.Float64() * float64(100)
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

/* ======================================================================== */

func parseTotalEntry(entry string) (value int64, err error) {

	if len(entry) < 7 {
		err = fmt.Errorf("impossibly short total entry")
		return
	}

	if !strings.HasPrefix(entry, "total") {
		// Far more likely...
		// - Wrong data item was passed to this call, or
		// - Reading a different(ly formatted) file.
		err = fmt.Errorf("mislabeled total entry")
		return
	}

	parts := strings.Split(entry, "=")

	if len(parts) != 2 {
		err = fmt.Errorf("invalid total split")
		return
	}

	value, err = strconv.ParseInt(parts[1], 10, 64)

	return
}
