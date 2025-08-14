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

// Options is a structure of possible arguments when creating a new StallInfo
// instance. These options may be used to modify the behaviour of collections
// both for production and testing purposes.
type Options struct {
	SourceDirectory string
	Collect         CollectMap
}

/* ======================================================================== */

// NewOptions creates a new Options structure with default values that may be
// modified as desired.
func NewOptions() (psia *Options) {

	psia = new(Options)

	// Set defaults
	psia.SourceDirectory = "/proc/pressure"
	psia.Collect = CollectAll

	return
}
