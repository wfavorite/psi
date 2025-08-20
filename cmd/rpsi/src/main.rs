
mod cmdline;
mod clbasics;
mod pressure;
mod version;


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
    let mut data = pressure::init_pressure_data();
    data.refresh();

    // This may be a single JSON dump.
    if cl.json() {
        // Serialize to a pretty-printed JSON string
        let json_str = serde_json::to_string_pretty(&data).expect("Failed to serialize to pretty JSON");
        println!("{}", json_str);
        std::process::exit(0);
    }

    // STUBbed from here

    println!("Begin debug / non-implementation block");
    println!("  -m(onochrome)    : {}", cl.mono());
    println!("  -t(imestamp)     : {}", cl.tmstp());
    println!("  -w(ide output)   : {}", cl.wide());
    println!("  iteration value  : {}", cl.iter());


}





