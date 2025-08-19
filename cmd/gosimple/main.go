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

			// I get one of these on occasion that was 0.
			// It might need to filtered out?

			//fmt.Fprintf(os.Stderr, "ERROR: Unknown event received - 0x%x\n", pfd.revents)
			//os.Exit(1)

			fmt.Printf("UnknownEvent(0x%x)\n", pfd.revents)
		}

	}
}
