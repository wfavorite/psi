use std::env;

/* ------------------------------------------------------------------------ */

pub struct CmdLine {
    opt_about: bool,
    opt_usage: bool,
    opt_json: bool,
    opt_mono: bool,
    opt_rando: bool, // Undocumented
    opt_tmstp: bool,
    opt_wide: bool,
    opt_iter: u64,
    opt_error: String,
}

/* ======================================================================== */

impl CmdLine {
    pub fn about(&self) -> bool {
        self.opt_about
    }

    pub fn help(&self) -> bool {
        self.opt_usage
    }

    pub fn json(&self) -> bool {
        self.opt_json
    }

    pub fn mono(&self) -> bool {
        self.opt_mono
    }

    pub fn rando(&self) -> bool {
        self.opt_rando
    }

    pub fn tmstp(&self) -> bool {
        self.opt_tmstp
    }

    pub fn wide(&self) -> bool {
        self.opt_wide
    }

    pub fn iter(&self) -> u64 {
        self.opt_iter
    }

    pub fn errored(&self) -> bool {
        self.opt_error.len() > 0
    }

    pub fn error(&self) -> &str {
        &self.opt_error
    }
}

/* ======================================================================== */

pub fn parse_cmd_line() -> CmdLine {
    let mut opt_about = false;
    let mut opt_usage = false;
    let mut opt_json = false;
    let mut opt_mono = false;
    let mut opt_rando = false;
    let mut opt_tmstp = false;
    let mut opt_wide = false;
    let mut opt_iter: u64 = 0;
    let mut opt_error: String = String::new();

    let mut non_flag_arg: String = String::new();
    let mut non_flag_pos: i32 = -1;
    let mut last_flag_pos: i32 = -1;

    // Pull arguments local
    let args: Vec<String> = env::args().collect();

    for (i, arg) in args.iter().skip(1).enumerate() {
        last_flag_pos = i as i32;

        // Save this... there needs to be a way to tell if an argument was the
        // last argument. To do this... probably assign the value and the
        // position to local variables, and if it was the last then process
        // after the loop has finished.
        //println!("  argument {} is {}", i + 1, arg);

        match arg.as_str() {
            "-a" => opt_about = true,
            "-h" => opt_usage = true,
            "-j" => opt_json = true,
            "-m" => opt_mono = true,
            "-r" => opt_rando = true,
            "-t" => opt_tmstp = true,
            "-w" => opt_wide = true,
            &_ => {
                non_flag_arg = String::from(arg);
                non_flag_pos = i as i32;
            }
        }
    }

    if last_flag_pos >= 0 {
        // Arguments were processed.

        if non_flag_arg.len() > 0 {
            // One was not a standard / expected argument...
            // but it may be a number.

            /* Not using this
            if non_flag_arg.starts_with("-") {
                opt_error = format!("The {} argument was not understood.", non_flag_arg);
            }
            */

            if last_flag_pos == non_flag_pos {
                // It was the last argument - so test for a number

                // This handles negative numbers (that look like flags) as well.
                match non_flag_arg.parse::<u64>() {
                    Ok(num) => opt_iter = num,
                    Err(_) => {
                        opt_error = format!("The {} argument was not understood.", non_flag_arg)
                    }
                }
            } else {
                // It was not the last - so it is an error, regardless.
                opt_error = format!("The {} argument was not understood.", non_flag_arg);
            }
        }
    }

    CmdLine {
        opt_about: opt_about,
        opt_usage: opt_usage,
        opt_json: opt_json,
        opt_mono: opt_mono,
        opt_rando: opt_rando,
        opt_tmstp: opt_tmstp,
        opt_wide: opt_wide,
        opt_iter: opt_iter,
        opt_error: opt_error,
    }
}
