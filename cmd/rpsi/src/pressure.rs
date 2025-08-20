use chrono::{Local};
use serde::Serialize;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
//use chrono::prelude::*;

/* ------------------------------------------------------------------------ */

#[derive(Serialize)]
pub struct PressureLine {
    avg10: f64,
    avg60: f64,
    avg300: f64,
    total: i64,
}

impl PressureLine {
    fn ingest_line(&mut self, line: &str) {
        let parts: Vec<&str> = line.split_whitespace().collect();

        if parts.len() != 5 {
            self.set_error_values();
            return;
        }

        // Validate that this is a 'some' of 'full' line.

        let type_part = parts[0];

        if type_part != "some" && type_part != "full" {
            self.set_error_values();
            return;
        }

        // Read the first value.
        self.avg10 = parse_float_equality(parts[1]);
        self.avg60 = parse_float_equality(parts[2]);
        self.avg300 = parse_float_equality(parts[3]);
        self.total = parse_integer_equality(parts[4])
    }

    fn set_error_values(&mut self) {
        self.avg10 = -1.0;
        self.avg60 = -1.0;
        self.avg300 = -1.0;
        self.total = -1;
    }
}

/* ========================================================================= */

fn parse_float_equality(item: &str) -> f64 {
    let parts: Vec<&str> = item.split("=").collect();

    if parts.len() != 2 {
        return -1.0;
    }

    let numeric_part = parts[1];

    match numeric_part.parse::<f64>() {
        Ok(num) => return num,
        Err(_) => return -1.1,
    }
}

/* ========================================================================= */

fn parse_integer_equality(item: &str) -> i64 {
    let parts: Vec<&str> = item.split("=").collect();

    if parts.len() != 2 {
        return -1;
    }

    let numeric_part = parts[1];

    match numeric_part.parse::<i64>() {
        Ok(num) => return num,
        Err(_) => return -1,
    }
}

/* ------------------------------------------------------------------------ */

#[derive(Serialize)]
pub struct PressureFile {
    some: PressureLine,
    full: PressureLine,
}

impl PressureFile {
    fn ingest_file(&mut self, file_target: &str) {
        let proc_dir = Path::new("/proc/pressure");
        let filename = proc_dir.join(file_target);

        let file = match File::open(filename) {
            Ok(f) => f,
            Err(_) => {
                self.set_error_values();
                return;
            }
        };

        let reader = io::BufReader::new(file);

        for line in reader.lines() {
            match line {
                Ok(line) => {
                    if line.contains("some") {
                        self.some.ingest_line(&line);
                    }

                    if line.contains("full") {
                        self.full.ingest_line(&line);
                    }
                }
                Err(_) => {}
            }
        }
    }

    fn set_error_values(&mut self) {
        self.some.set_error_values();
        self.full.set_error_values();
    }

    fn data_wide(&self, o: &Options) {
        if o.mono {
            print!(
                "{:>6.2} {:>6.2} {:>6.2} {:>6.2} {:>6.2} {:>6.2} ",
                self.some.avg10,
                self.some.avg60,
                self.some.avg300,
                self.full.avg10,
                self.full.avg60,
                self.full.avg300,
            );
        } else {
            let mut c10 = derive_colour(o, self.some.avg10);
            let mut c60 = derive_colour(o, self.some.avg60);
            let mut c300 = derive_colour(o, self.some.avg300);

            print!(
                "{}{:>6.2}\x1b[0m {}{:>6.2}\x1b[0m {}{:>6.2}\x1b[0m ",
                c10, self.some.avg10, c60, self.some.avg60, c300, self.some.avg300,
            );

            c10 = derive_colour(o, self.full.avg10);
            c60 = derive_colour(o, self.full.avg60);
            c300 = derive_colour(o, self.full.avg300);

            print!(
                "{}{:>6.2}\x1b[0m {}{:>6.2}\x1b[0m {}{:>6.2}\x1b[0m ",
                c10, self.full.avg10, c60, self.full.avg60, c300, self.full.avg300,
            );
        }
    }

    fn data_condensed(&self, some: bool, o: &Options) {
        if o.mono {
            if some {
                print!(
                    "{:>3.0} {:>3.0} {:>3.0} ",
                    self.some.avg10, self.some.avg60, self.some.avg300,
                )
            } else {
                print!(
                    "{:>3.0} {:>3.0} {:>3.0} ",
                    self.full.avg10, self.full.avg60, self.full.avg300,
                )
            }
        } else {
            if some {
                let c10 = derive_colour(o, self.some.avg10);
                let c60 = derive_colour(o, self.some.avg60);
                let c300 = derive_colour(o, self.some.avg300);

                print!(
                    "{}{:>3.0}\x1b[0m {}{:>3.0}\x1b[0m {}{:>3.0}\x1b[0m ",
                    c10, self.some.avg10, c60, self.some.avg60, c300, self.some.avg300,
                );
            } else {
                let c10 = derive_colour(o, self.full.avg10);
                let c60 = derive_colour(o, self.full.avg60);
                let c300 = derive_colour(o, self.full.avg300);

                print!(
                    "{}{:>3.0}\x1b[0m {}{:>3.0}\x1b[0m {}{:>3.0}\x1b[0m ",
                    c10, self.full.avg10, c60, self.full.avg60, c300, self.full.avg300,
                );
            }
        }
    }
}

/* ========================================================================= */

fn derive_colour(o: &Options, fval: f64) -> String {
    // STUB: Returning heap allocated strings is expensive.
    // STUB: Look for an alternative.
    if fval > o.red_threshold {
        return String::from("\x1b[91m");
    } else {
        if fval > o.yellow_threshold {
            return String::from("\x1b[93m");
        } else {
            return String::from("\x1b[92m");
        }
    }
}

/* ------------------------------------------------------------------------ */

#[derive(Serialize)]
pub struct PressureData {
    cpu: PressureFile,
    io: PressureFile,
    irq: PressureFile,
    memory: PressureFile,

    timestamp: String,

    #[serde(skip)]
    opts: Options,
}

impl PressureData {
    pub fn refresh(&mut self) {
        self.cpu.ingest_file("cpu");
        self.io.ingest_file("io");
        self.irq.ingest_file("irq");
        self.memory.ingest_file("memory");

        let local = Local::now();
        self.timestamp = local.format("%H:%M:%S.%3f").to_string();
    }

    pub fn print_header(&self) {
        if self.opts.wide {
            self.header_wide();
        } else {
            self.header_condensed();
        }
    }

    fn header_wide(&self) {
        // ===== First line =====

        print!("#");

        if self.opts.tmsp {
            print!("             ");
        }

        // There is no collection option - so print all.
        print!(" {:<8}                                 ", "CPU");
        print!(" {:<8}                                 ", "IO");
        print!(" {:<8}                                 ", "IRQ");
        print!(" {:<8}                                 ", "Memory");

        // EOL
        println!();

        // ===== Second line ====

        print!("#");

        if self.opts.tmsp {
            print!("             ");
        }

        print!(" Some________________ Full________________");
        print!(" Some________________ Full________________");
        print!(" Some________________ Full________________");
        print!(" Some________________ Full________________");

        // EOL
        println!();

        // ===== Third line =====

        print!("#");

        if self.opts.tmsp {
            print!(" Timestamp   ");
        }

        print!(
            " {:>6} {:>6} {:>6} {:>6} {:>6} {:>6}",
            "avg10", "avg60", "avg300", "avg10", "avg60", "avg300"
        );
        print!(
            " {:>6} {:>6} {:>6} {:>6} {:>6} {:>6}",
            "avg10", "avg60", "avg300", "avg10", "avg60", "avg300"
        );
        print!(
            " {:>6} {:>6} {:>6} {:>6} {:>6} {:>6}",
            "avg10", "avg60", "avg300", "avg10", "avg60", "avg300"
        );
        print!(
            " {:>6} {:>6} {:>6} {:>6} {:>6} {:>6}",
            "avg10", "avg60", "avg300", "avg10", "avg60", "avg300"
        );

        // EOL
        println!();
    }

    fn header_condensed(&self) {
        // ===== First line =====

        print!("#");

        if self.opts.tmsp {
            print!("             ");
        }

        print!("    ");

        print!("{:<8}    ", "CPU");
        print!("{:<8}    ", "IO");
        print!("{:<8}    ", "IRQ");
        print!("{:<8}    ", "Memory");

        // EOL
        println!();

        // ===== Second line ====

        print!("#");

        if self.opts.tmsp {
            print!(" Timestamp   ");
        }

        print!("    ");

        print!("{:>3} {:>3} {:>3} ", "10s", "1m", "5m");
        print!("{:>3} {:>3} {:>3} ", "10s", "1m", "5m");
        print!("{:>3} {:>3} {:>3} ", "10s", "1m", "5m");
        print!("{:>3} {:>3} {:>3} ", "10s", "1m", "5m");

        // EOL
        println!();
    }

    pub fn print_line(&self) {
        if self.opts.wide {
            self.line_wide();
        } else {
            self.line_condensed();
        }
    }

    fn line_wide(&self) {
        print!("  ");

        if self.opts.tmsp {
            print!("{:<12} ", self.timestamp);
        }

        self.cpu.data_wide(&self.opts);
        self.io.data_wide(&self.opts);
        self.irq.data_wide(&self.opts);
        self.memory.data_wide(&self.opts);

        // EOL
        println!();
    }

    fn line_condensed(&self) {
        if self.opts.tmsp {
            print!("{:<12} ", self.timestamp);
        }

        print!("Some ");
        self.cpu.data_condensed(true, &self.opts);
        self.io.data_condensed(true, &self.opts);
        self.irq.data_condensed(true, &self.opts);
        self.memory.data_condensed(true, &self.opts);
        println!();

        if self.opts.tmsp {
            print!("{:<12} ", self.timestamp);
        }

        print!("Full ");
        self.cpu.data_condensed(false, &self.opts);
        self.io.data_condensed(false, &self.opts);
        self.irq.data_condensed(false, &self.opts);
        self.memory.data_condensed(false, &self.opts);
        println!();
    }
}

/* ------------------------------------------------------------------------ */

pub struct Options {
    // STUB: Collect map
    wide: bool,
    mono: bool,
    tmsp: bool,
    rando: bool,
    red_threshold: f64,
    yellow_threshold: f64,
}

impl Options {
    pub fn set_wide(&mut self, value: bool) {
        self.wide = value;
    }

    pub fn set_mono(&mut self, value: bool) {
        self.mono = value;
    }

    pub fn set_timestamp(&mut self, value: bool) {
        self.tmsp = value;
    }

    pub fn set_debug_random(&mut self, value: bool) {
        self.rando = value;
    }

    pub fn init_pressure_data(self) -> PressureData {
        let cpu = init_pressure_file();
        let io = init_pressure_file();
        let irq = init_pressure_file();
        let memory = init_pressure_file();

        PressureData {
            cpu: cpu,
            io: io,
            irq: irq,
            memory: memory,
            timestamp: String::new(), // STUB: Fix this
            opts: self,
        }
    }
}

/* ========================================================================= */

pub fn new_options() -> Options {
    Options {
        wide: false,
        mono: false,
        tmsp: false,
        rando: false,
        red_threshold: 90.0,
        yellow_threshold: 10.0,
    }
}

/* ========================================================================= */

fn init_pressure_file() -> PressureFile {
    let some = init_pressure_line();
    let full = init_pressure_line();

    PressureFile {
        some: some,
        full: full,
    }
}

/* ========================================================================= */

fn init_pressure_line() -> PressureLine {
    PressureLine {
        avg10: 0.0,
        avg60: 0.0,
        avg300: 0.0,
        total: 0,
    }
}
