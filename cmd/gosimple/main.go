// Package main implements gosimple; a simplified Go implementation of the PSI sample code.
package main

/*
#include <poll.h>

typedef struct pollfd pollfd_t;
*/
import "C"
import (
	"fmt"
	"os"
	"syscall"
)

/* ======================================================================== */

func main() {

	gf, err := os.OpenFile("/proc/pressure/io", syscall.O_RDWR|syscall.O_NONBLOCK, 0)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: open -", err.Error())
		os.Exit(1)
	}

	defer gf.Close()

	// I use the same value as the cpsi/test_run.sh test script. It is known to
	// occasionally trigger on my dev system with these values.
	trigger := []byte("some 15000 1000000")

	_, err = gf.Write(trigger)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: write -", err.Error())
		os.Exit(1)
	}

	fd := int(gf.Fd())

	pfd := C.pollfd_t{
		fd:     C.int(fd),
		events: C.POLLPRI,
	}

	for {
		// +----- The C pollfd structure
		// | +--- Only one in the array
		// | | +- Five second timeout
		// | | |
		// | | +---------------+
		// | +--------------+  |
		// +----------V     V  V
		prv := C.poll(&pfd, 1, 5000)

		if prv == 0 {
			fmt.Println("TimeoutInterval")
			continue
		}

		if pfd.revents&C.POLLERR == C.POLLERR {
			fmt.Fprintln(os.Stderr, "ERROR: got POLLERR, event source is gone")
			os.Exit(1)
		}

		if pfd.revents&C.POLLPRI == C.POLLPRI {
			fmt.Println("HitThreshold")
		} else {

			// I get one of these on occasion that was 0. This is likely
			// associated with the (<0) return value that I fail to check for
			// above.
			// Possible approaches:
			// - Simply filter it out?
			// - Bring in errno and strerror() for reporting.

			//fmt.Fprintf(os.Stderr, "ERROR: Unknown event received - 0x%x\n", pfd.revents)
			//os.Exit(1)

			fmt.Printf("UnknownEvent(0x%x)\n", pfd.revents)
		}
	}
}
