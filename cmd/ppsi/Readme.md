# ppsi - Parent PSI

## Overview

``ppsi`` is a *skeletal* Go "framework" for launching and monitoring [``cpsi``](../cpsi/Readme.md) instances.

It is considered skeletal because...

- It is PoC code.
- What to launch/monitor would ideally be passed as a config, not hard-coded into the implementation.
- The whole point of a Go calling framework is that Go can more easily plug into some sort of monitoring / alerting interface. This is not implemented (as, again, this is PoC code).

## Usage

> The ``cpsi`` monitor must exist (in the peer "cpsi" directory). This binary will attempt to launch client-monitors from that directory.

Just call it:

``./ppsi``

## (Suggestions) Making a production instance

This is PoC code that is designed to show the feasibility of such a design. Some notes:

First, the ideal design does NOT use client-monitors. A more appropriate design would / should launch monitors in go threads. The approach would likely be a monitor function (in Go) that reports back to the event Q via a chan. The monitor function would either use the syscall module, or more likely a CGO call to the ``poll()`` function. All other functionality should be fine from the Go standard library.

*If* a parent-monitor design is used, then...

- The ``cpsi`` binary should be embedded in the ``ppsi`` binary. It would be copied out to the local file system on invocation. (My build of ``cpsi`` is 23K).
- The config of what to monitor would (ideally) be from a config file rather than compiled in. Even if the config was embedded, it is better than *coding* it in as is done here.
- Threshold events would be raised to some sort of larger monitoring framework. (This is the point of using Go for reporting.)
- The ``ppsi`` instance would need to be launched as / check for root or privilege elevate when calling ``cpsi`` as root-level permissions are required to monitor this interface.

## Findings / lessons learned

The current two-part design is rather trivial, but still overkill. This has way more moving parts than necessary. As a first stab approach to groking the problem, it is ok - perhaps.

Both ``cpsi`` and ``ppsi`` are decent skeletal launch points from which to build similar services. So they might be good to steal from, even if they are not *proper* implementations in themselves. The reader is free to do so - i claim no license on the PoC implementations (while the ``psi`` tool uses the "BSD 3-Clause" license).
