# cpsi - Client PSI

## Overview

This is a simple C program that monitors (poll()s) the proc filesystem pressure file as passed on the command line. It is designed to be called by the [``ppsi``](../ppsi/Readme.md) 'framework'.

## Usage

``cpsi`` is called with three positional arguments. (It is not meant to be called as a separate tool, although it could be modified to do so.) These arguments are:

1. The destination Unix socket to write events to.
2. The file in ``/proc/pressure`` to write to.
3. The (three) threshold arguments.

The calling convention:

``cpsi <evt_dest> <monitor_file> <monitor_args>``

An example call:

``cpsi /tmp/unix_alarm.s io "some 15000 1000000"``

``cpsi`` can be called in a "debug" mode using various options. They are:

- Pass "stdout" as the Unix socket name. Events will be written to stdout, and not a Unix socket.
- Define the ``CPSI_VERBOSE_BN`` variable. This is the basename of a log file that debug messages will be written to. A PID will be appended to the filename.

I have included a [test_run.sh](test_run.sh) script to run in local test mode with a threshold likely to trigger on a lightly used dev system.

## Notes

The most common failure is typically errno 22 on the write to the ``/proc/pressure/xxx`` file. When input values are out of bounds, the errno "Invalid argument" error is thrown.
