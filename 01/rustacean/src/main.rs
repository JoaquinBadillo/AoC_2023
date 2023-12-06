use std::env;
use std::fs;
use std::io::BufRead;
use std::collections::HashMap;

fn main() {
    let args: Vec<String> = env::args().collect();
    let file_path: &str;

    if args.len() < 2 {
        file_path = "./input.txt";
    } else {
        file_path = &args[1];
    }

    let numbers = HashMap::from([
        ("zero", 0),
        ("one", 1),
        ("two", 2),
        ("three", 3),
        ("four", 4),
        ("five", 5),
        ("six", 6),
        ("seven", 7),
        ("eight", 8),
        ("nine", 9)
    ]);

    let file = fs::File::open(file_path)
        .expect("File not found");
    
    let total = std::io::BufReader::new(file).lines().fold(0, |total, line| {
        let mut first = true;
        let line = line.expect("Error reading line");
        
        let mut parsed: String = String::new().to_owned();
        let vals = line.chars().fold((0,0), |mut digits, c| {
            let mut num = 0;
            let mut found = false;

            let data = c.to_digit(10);

            if data == None {
                parsed.push(c);
                for (key, value) in numbers.iter() {
                    match parsed.find(key) {
                        Some(_) => {
                            found = true;
                            num = *value;

                            parsed.clear();
                            parsed.push(c);

                            break;
                        },
                        None => {}
                    }
                }
            } else {
                parsed.clear();
                found = true;
                num = data.unwrap();
            }
            
            if !found {
                return digits;
            }

            if first {
                first = false;
                digits.0 = num;
            }

            digits.1 = num;
            digits

        });
        
        total + vals.0 * 10 + vals.1
    });

    println!("{}", total);
}