# rpsi - Rust psi

## Overview

This is Rust port of the [Go psi](../psi/) utility. This port is not *definitive* as it is more about comparing a Go and Rust port(ing experience).

## Usage

> Most of these options are non-functional at this time (version 0.1.0).

```text
$ target/release/rpsi -h
rpsi - Rust pressure reporter
Usage: psi <options> <int>
Options:
  -a     Show about information.
  -h     Show this usage information.
  -j     Dump current stats as a JSON structure. This option is incompatible
         with other options or interval printing.
  -m     Print output in monochrome. Default is ANSI colour.
  -t     Print a timestamp on each line of tabular output.
  -w     Print in a wide format (potentially beyond 80 chars).
  <int>  Print tabular stats on interval. The supplied time is in seconds.
```

## Notes

- This is not a full functional port (of the Go version) so comparing binary sizes and performance is not appropriate. (Although... the Rust version is considerably smaller - with *most* of the implementation complete.)
- Build with ``cargo build --release`` from this directory.
- You have to go external to get timestamps and JSON. Hmmmm.
