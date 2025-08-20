use core::time;


mod cmdline;
mod clbasics;
mod pressure;
mod version;

/* ======================================================================== */

fn main() {
    
    let cl = cmdline::parse_cmd_line();

    // Handle the command line basics.

    if cl.errored() {
        eprintln!("ERROR: {}", cl.error());
        std::process::exit(1);
    }

    if cl.about() {
        clbasics::handle_about(version::VERSION_STRING);
        std::process::exit(0);
    }

    if cl.help() {
        clbasics::handle_help();
        std::process::exit(0);
    }

    // Command line basics / early exits are over. We will be collecting data.

    let mut opts = pressure::new_options();

    // VOIR: This is an ugly hand-off but i think it is indicative of
    // VOIR: cross-cutting concerns in tooling of this size. The goal is to
    // VOIR: separate 'classes' but not incur the overly expensive cost of
    // VOIR: doing so. This is quite possibly a Go-bias on my part. I should
    // VOIR: also admit that the design of representation in the same class as
    // VOIR: the collector is an anti-pattern (discussed in the Go port).
    opts.set_debug_random(cl.rando());
    opts.set_mono(cl.mono());
    opts.set_timestamp(cl.tmstp());
    opts.set_wide(cl.wide());

    let mut data = opts.init_pressure_data();
    data.refresh();

    // This may be a single JSON dump.
    if cl.json() {
        // Serialize to a pretty-printed JSON string
        let json_str = serde_json::to_string_pretty(&data).expect("Failed to serialize to pretty JSON");
        println!("{}", json_str);
        std::process::exit(0);
    }



    data.print_header();
    data.print_line();

    let iteration_time = cl.iter();
    let sleep_time = time::Duration::from_secs(iteration_time);
if iteration_time > 0 {
    loop {
        std::thread::sleep(sleep_time);

        data.refresh();

        data.print_line();

    }
}
  


}





