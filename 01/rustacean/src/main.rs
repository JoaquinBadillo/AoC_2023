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
    
    let total = std::io::BufReader::new(file).lines().fold(0, |total, line| {
        let mut first = true;
        let line = line.expect("Error reading line");

        let vals = line.chars().fold((0,0), |mut digits, c| {
            match c.to_digit(10) {
                Some(num) => {
                    if first {
                        first = false;
                        digits.0 = num;
                    }

                    digits.1 = num;
                    digits
                },
                None => digits
            }
        });
        
        total + vals.0 * 10 + vals.1
    });

    println!("{}", total);
}