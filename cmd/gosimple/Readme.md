# gosimple - A simplistic Go PSI implementation

## Overview

This is a simple implementation of the PSI monitor in the [kernel.org PSI documentation](https://www.kernel.org/doc/html/v5.4/accounting/psi.html) that is mostly in Go but with a CGO call to ``poll()``.

## Notes

- The ``os.OpenFile()`` call does not have an explicit ``os.O_NONBLOCK`` option. I omitted it, and it still works. It should be possible to wrap the C implementation of open() - if necessary, but i used the ``syscall.O_NONBLOCK`` value - and that worked as well.
- I created a typedef on the ``struct pollfd`` type. This made declaring it much easier.
- I added a 5 second 'IntervalTimeout` event so that you can tell that it is doing *something*.
- I occasionally get an empty (0) ``revents`` value. I treat this as any other event (I do not exit as the sample code does).

