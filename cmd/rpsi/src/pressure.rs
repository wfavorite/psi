
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use serde::Serialize;

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
            }, 
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
                },
                Err(_) => {},
            }
        }
    }

    fn set_error_values(&mut self) {
        self.some.set_error_values();
        self.full.set_error_values();
    }
}


/* ------------------------------------------------------------------------ */

#[derive(Serialize)]
pub struct PressureData {
    cpu: PressureFile,
    io: PressureFile,
    irq: PressureFile,
    memory: PressureFile,
}

impl PressureData {
    pub fn refresh(&mut self) {
        self.cpu.ingest_file("cpu");
        self.io.ingest_file("io");
        self.irq.ingest_file("irq");
        self.memory.ingest_file("memory");
    }
}




/* ========================================================================= */

pub fn init_pressure_data() -> PressureData {

    let cpu = init_pressure_file();
    let io = init_pressure_file();
    let irq = init_pressure_file();
    let memory = init_pressure_file();

    PressureData{
        cpu: cpu,
        io: io,
        irq: irq,
        memory: memory,
    }
}

/* ========================================================================= */

fn init_pressure_file() -> PressureFile {
    
    let some = init_pressure_line();
    let full = init_pressure_line();
    
    PressureFile {
        some: some,
        full:  full,
    }
}

/* ========================================================================= */

fn init_pressure_line() -> PressureLine {
    PressureLine { avg10: 0.0, avg60: 0.0, avg300: 0.0, total: 0 }
}







