# psi

This is a module and command line utility to read ``/proc/pressure`` [psi](https://www.kernel.org/doc/html/v5.4/accounting/psi.html) stats.

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
