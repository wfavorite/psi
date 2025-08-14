package psi

/* ------------------------------------------------------------------------ */

// CollectMap is the type backing the collection argument of which sources to
// pull data for.
type CollectMap uint16

/* ------------------------------------------------------------------------ */

// The Collect arguments are a bitmap of possible collection sources provided
// in /proc/pressure.
const (
	CollectNone CollectMap = 0x0000
	CollectCPU  CollectMap = 0x0001
	CollectIO   CollectMap = 0x0002
	CollectIRQ  CollectMap = 0x0004
	CollectMem  CollectMap = 0x0008
	CollectAll  CollectMap = CollectCPU | CollectIO | CollectIRQ | CollectMem
)

/* ------------------------------------------------------------------------ */

// WidthOption is the datatype backing the Width argument of the output
// methods.
type WidthOption int

/* ------------------------------------------------------------------------ */

const (
	// Condensed is designed to print a full set of lesser precision data on
	// two lines in 80 character width screens per-iteration.
	Condensed WidthOption = iota

	// Wide will print data at full precision on a single line per-iteration.
	Wide
)

/* ------------------------------------------------------------------------ */

// Options is a structure of possible arguments when creating a new StallInfo
// instance. These options may be used to modify the behaviour of collections
// both for production and testing purposes.
type Options struct {

	// Collect is the bitmap of sources to collect.
	// STUB: Rename sources
	Collect CollectMap

	// Width is an output formatting option. It is either Wide that prints all
	// data (with full precision) on one line per-iteration, or Condensed that
	// prints two-lines per-iteration with lower precision.
	Width WidthOption

	// Monochrome (when true) disables ANSI colour printing.
	Monochrome bool

	// Timestamp (when true) prints the timestamp of the collection time per
	// line.
	TimeStamp bool

	RedThreshold    float64
	YellowThreshold float64

	// The following are debug options.

	// SourceDirectory is the directory from which to read the pressure files.
	// This defaults to /proc/pressure.
	SourceDirectory string

	// RandomValues will generate random values for collected data. This can
	// be used to check ANSI colours in the output options.
	RandomValues bool
}

/* ======================================================================== */

// NewOptions creates a new Options structure with default values that may be
// modified as desired.
func NewOptions() (psia *Options) {

	psia = new(Options)

	// Set defaults
	psia.SourceDirectory = "/proc/pressure"
	psia.Collect = CollectAll
	psia.Width = Condensed
	psia.TimeStamp = false
	psia.Monochrome = false
	psia.RedThreshold = 90
	psia.YellowThreshold = 10

	return
}
