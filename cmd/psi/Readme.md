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

The [``psi`` module](../../pkg/psi) implements the output and ANSI colouring. This is probably not a proper approach (mixing output/representation with collection). There is some discussion about this in the top of the [PrintHeader.go](../../pkg/psi/PrintHeader.go) file.
