
use std::io::{self, Write};

/* ========================================================================= */

/*
    VOIR: handle_header is one method of writing to stdout and handling
    VOIR: exceptions when doing so. The trailing .ok() effectively
    VOIR: ignores the errors.
*/

fn handle_header() {
    let stdout = io::stdout();
    let mut handle = stdout.lock();

    handle.write(b"rpsi - Rust pressure reporter\n").ok();
}


/* ========================================================================= */

/*
    VOIR: handle_about is the 'idiomatic' way of handling writes to stdout.
    VOIR: This method allows for formatting options and is a simplified 'macro'
    VOIR: for the write. Note that println!() takes a lock on stdout which is a
    VOIR: requirement for writes - even if you are doing something as simple as
    VOIR: this - which means the macro is even more appropriate.
    VOIR:
    VOIR: Personally... i consider such 'bultins' / macros an anti-pattern for
    VOIR: writing to stdout. My bias seems to be getting in the way in at least
    VOIR: this specific case.
*/

pub fn handle_about(version: &str) {
    handle_header();
    println!("  Version: {}", version);    
}

/* ========================================================================= */

/*
    VOIR: handle_help is *another* means of ignoring errors when (explicitly)
    VOIR: writing to stout.
*/

pub fn handle_help() {

    handle_header();

    /* Abandoning this in favor of the simpler option when i copy-pasta'd from the Go version
    let mut stdout = io::stdout();

    _ = stdout.write_all(b"  Usage: rpsi <options> <interval>\n");
    _ = stdout.write_all(b"  Options:\n");
    _ = stdout.write_all(b"    -a     Show about information.\n");
    _ = stdout.write_all(b"    -h     Show this usage information.\n");
    _ = stdout.write_all(b"    -j     Dump a (single) collection as a JSON structure.\n");
    */

    println!("Usage: psi <options> <int>");
    println!("Options:");
    println!("  -a     Show about information.");
    println!("  -h     Show this usage information.");
    println!("  -j     Dump current stats as a JSON structure. This option is incompatible");
    println!("         with other options or interval printing.");
    println!("  -m     Print output in monochrome. Default is ANSI colour.");
    println!("  -t     Print a timestamp on each line of tabular output.");
    println!("  -w     Print in a wide format (potentially beyond 80 chars).");
    println!("  <int>  Print tabular stats on interval. The supplied time is in seconds.")
    // This part differs.
    //intln!("  <int>  Print tabular stats on interval. The supplied time is assumed either");
    //intln!("         seconds (if an integer) or a Golang duration (eg: 500ms, 1s, 2m).");
}
