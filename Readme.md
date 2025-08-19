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

The ``psi`` tool is a simple monitor with basic ANSI colour representation of thresholds. Such tooling is appropriate for *investigating* a 'troubled' system. Such usage is an edge case and not representative of how best to watch these stats. The included "Other things" PoC code are such implementations - although skeletal in nature.

The "full" values are far more important than the "some". *Someone* is likely to block frequently. It would be nice to filter out the some results or just ANSI threshold-colour just the full results. If ``psi`` were to get regular use (and not re-invent a wheel somewhere) this functionality could be added.

> full vs some
>> "full" is when all processes feel pressure. This is indicative of a system under stress. "some" is when *some* processes are blocking. Most of the sample code mentioned below (in Other things) triggers on "some" as this threshold is far easier to hit on a lightly stressed development system.

The [``psi`` module](pkg/psi/) implements the output and ANSI colouring. This is probably not a proper approach (mixing output/representation with collection). There is some discussion about this in the top of the [PrintHeader.go](pkg/psi/PrintHeader.go) file.

## Other things

> Like most things I have shared on github, the ``ppsi``/``cpsi`` and ``gosimple`` code is not intended for production use. It is more about potential *designs* of services, not necessarily a Linux PSI monitoring implementation. If you want such tooling for production use - consider *hiring* me, and I will write/maintain/support such tooling for you.

Instead of *watching* pressure information it is more desirable to poll and alert on thresholds. This is discussed in the above kernel.org document. I have created a "client-server" approach to monitoring as a first attempt (aka: spin) at such a tool. The thinking is that such a tool (capable of easily plugging into some framework) should be written in a more 'supportable' language like Go. By "supportable" I am referring to the ability to easily leverage libraries that can do simple things like make RESTful calls or make gRPC calls to higher level reporting tooling/frameworks that would be a bit more *burdensome* in a pure C implementation. I should also volunteer that some higher level frameworks have existing monitors capable of monitoring file output. It would be trivial to add syslog output to a pure-C implementation. All of this said... I should stress that I am not solving for a *specific* case here, but kicking around some design issues that happen to be based on the Linux PSI data.

The first 'spin' (``ppsi``/``cpsi``) involved writing a pure Go 'handler' and a pure C 'trigger' and connecting them via a Unix socket. This is essentially a proof-of-concept for the ability to monitor.

The second 'spin' was to add a C ``poll()`` function in an otherwise pure Go implementation using CGO compiled ``poll()`` call and a bit of help (although possibly unnecessary) from the syscall package.

The third spin would be a properly daemonized & threaded Go/GCO monitor implementation - mixing the ``ppsi`` and ``gosimple`` variants. This would drop the framework-trigger (two binary) implementation and implement all monitors as threads connected to an Event Q that would raise the events into a larger monitoring solution.

- First spin: [``ppsi``](cmd/ppsi/Readme.md) the Go 'framework' that calls the [``cpsi``](cmd/cpsi/Readme.md) C poll implementation. Lessons learned can be found in the ``ppsi`` [Readme](cmd/ppsi/Readme.md) document.
- Second spin: [``gosimple``](cmd/gosimple/Readme.md) is a *very simple* CGO version of the example code in the PSI kernel.org documentation above.
- Third/alternate spin (a): Proposed Go 'daemon' monitor example. Not written at this time.
- Third/alternate spin (b): Pure C monitor that writes to syslog and is monitored by existing alterting tooling. Not written at this time.
