use std::env;
use std::fs;
use std::io::BufRead;

fn main() {
    let args: Vec<String> = env::args().collect();
    let file_path: &str;

    if args.len() < 2 {
        file_path = "./input.txt";
    } else {
        file_path = &args[1];
    }

    let file = fs::File::open(file_path)
        .expect("File not found");

    let mut total = 0;

    for line in std::io::BufReader::new(file).lines() {
        let mut vals = (0, 0);
        let mut first = true;

        let line = line.expect("Error reading line");

        for c in line.chars() {
            let num = match c.to_digit(10) {
                Some(n) => n,
                None => continue,
            };

            if first {
                vals.0 = num;
                first = false
            }

            vals.1 = num;
        }
        
        total += vals.0 * 10 + vals.1;
    }

    println!("{}", total);
}