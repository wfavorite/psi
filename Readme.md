# psi

## Overview

This repo is a (mostly) proper implementation of a ``/proc/pressure`` [PSI](https://www.kernel.org/doc/html/v5.4/accounting/psi.html) stats reader, and some tooling and design work around the concept.

## Contents

- [psi](cmd/psi/Readme.md) - The closest thing to a 'production' tool in this repo. This relies on the [psi module](pkg/psi/).
- [rpsi](cmd/rpsi/Readme.md) - A Rust port of the [Go psi](cmd/psi/Readme.md) utility.
- [ppsi](cmd/ppsi/Readme.md) - The PoC PSI 'framework' that calls / utilizes the [cpsi](cmd/cpsi/Readme.md) C trigger component.
- [cpsi](cmd/cpsi/Readme.md) - The PoC PSI 'trigger' that is called / utilized by the [ppsi](cmd/ppsi/Readme.md) Go framework.
- [gosimple](cmd/gosimple/Readme.md) - A simplified Go version of the C code found on the [PSI](https://www.kernel.org/doc/html/v5.4/accounting/psi.html) kernel documents page.

## Repo layout

This is a *Go* repo that happens to contain other code - including a Rust crate. It is primarily about the Go implementation and the other code is along for the ride.

## Other things

> Like most things I have shared on github, the ``ppsi``/``cpsi``, ``rpsi``, and ``gosimple`` code is not intended for production use. It is more about potential *designs* of services, not necessarily a Linux PSI monitoring implementation. If you want such tooling for production use - consider *hiring* me, and I will write/maintain/support such tooling for you.

Instead of *watching* pressure information it is more desirable to poll and alert on thresholds. This is discussed in the above kernel.org document. I have created a "client-server" approach to monitoring as a first attempt (aka: spin) at such a tool. The thinking is that such a tool (capable of easily plugging into some framework) should be written in a more 'supportable' language like Go. By "supportable" I am referring to the ability to easily leverage libraries that can do simple things like make RESTful calls or make gRPC calls to higher level reporting tooling/frameworks that would be a bit more *burdensome* in a pure C implementation. I should also volunteer that some higher level frameworks have existing monitors capable of monitoring file output. It would be trivial to add syslog output to a pure-C implementation. All of this said... I should stress that I am not solving for a *specific* case here, but kicking around some design issues that happen to be based on the Linux PSI data.

The first 'spin' (``ppsi``/``cpsi``) involved writing a pure Go 'handler' and a pure C 'trigger' and connecting them via a Unix socket. This is essentially a proof-of-concept for the ability to monitor.

The second 'spin' was to add a C ``poll()`` function in an otherwise pure Go implementation using CGO compiled ``poll()`` call and a bit of help (although possibly unnecessary) from the syscall package.

The third spin would be a properly daemonized & threaded Go/GCO monitor implementation - mixing the ``ppsi`` and ``gosimple`` variants. This would drop the framework-trigger (two binary) implementation and implement all monitors as threads connected to an Event Q that would raise the events into a larger monitoring solution.

- First spin: [``ppsi``](cmd/ppsi/Readme.md) the Go 'framework' that calls the [``cpsi``](cmd/cpsi/Readme.md) C poll implementation. Lessons learned can be found in the ``ppsi`` [Readme](cmd/ppsi/Readme.md) document.
- Second spin: [``gosimple``](cmd/gosimple/Readme.md) is a *very simple* CGO version of the example code in the PSI kernel.org documentation above.
- Third/alternate spin (a): Proposed Go 'daemon' monitor example. Not written at this time.
- Third/alternate spin (b): Pure C monitor that writes to syslog and is monitored by existing alterting tooling. Not written at this time.
