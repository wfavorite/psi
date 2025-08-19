# psi

This is a module and command line utility to read ``/proc/pressure`` [PSI](https://www.kernel.org/doc/html/v5.4/accounting/psi.html) stats.

## Usage

```text
psi - Pressure reporter
  Usage: psi <options> <int>
  Options:
    -a     Show about information.
    -h     Show this usage information.
    -j     Dump current stats as a JSON structure. This option is incompatible
           with other options or interval printing.
    -m     Print output in monochrome. Default is ANSI colour.
    -t     Print a timestamp on each line of tabular output.
    -w     Print in a wide format (potentially beyond 80 chars).
    <int>  Print tabular stats on interval. The supplied time is assumed either
           seconds (if an integer) or a Golang duration (eg: 500ms, 1s, 2m).
```

## Design / usage notes

The ``psi`` tool is a simple monitor with basic ANSI colour representation of thresholds. Such tooling is appropriate for *investigating* a 'troubled' system. Such usage is an edge case and not representative of how best to watch these stats. The included PoC code is such an implementation - although skeletal in nature.

## Other things

> Like most things I have shared on github, the ``ppsi``/``cpsi`` and ``gosimple`` code is not intended for production use. If you want such tooling for production use - consider *hiring* me, and I will write/maintain/support such tooling for you.

Instead of *watching* pressure information it is more desirable to poll and alert on thresholds. This is discussed in the above kernel.org document. I have created a "client-server" approach to monitoring as a first attempt at such a tool. The thinking is that such a tool (capable of easily plugging into some framework) should be written in a more 'supportable' language like Go. By "supportable" I am referring to the ability to easily leverage libraries that can do simple things like make RESTful calls or make gRPC calls to higher level reporting tooling/frameworks that would be a bit more *burdensome* in a pure C implementation.

The first 'spin' (``ppsi``/``cpsi``) involved writing a pure Go 'handler' and a pure C 'trigger' and connecting them via a Unix socket. This is essentially a proof-of-concept for the ability to monitor.

The second 'spin' was to add a C ``poll()`` function in an otherwise pure Go implementation using CGO compiled ``poll()`` call and a bit of help (although possibly unnecessary) from the syscall package.

The third spin would be a properly daemonized & threaded Go/GCO monitor implementation - mixing the ``ppsi`` and ``gosimple`` variants.

- First spin: [``ppsi``](cmd/ppsi/Readme.md) the Go 'framework' that calls the [``cpsi``](cmd/cpsi/Readme.md) C poll implementation. Lessons learned can be found in the ``ppsi`` [Readme](cmd/ppsi/Readme.md) document.
- Second spin: [``gosimple``](cmd/gosimple/Readme.md) is a *very simple* CGO version of the example code in the PSI kernel.org documentation above.
- Third spin: Proposed Go 'daemon' monitor example. Not written at this time.
