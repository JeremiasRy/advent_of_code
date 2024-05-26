use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};

fn main() {
    let spelled_to_digit_hash: HashMap<&str, &str> = HashMap::<&str, &str>::from([
        ("one", "1"),
        ("two", "2"),
        ("three", "3"),
        ("four", "4"),
        ("five", "5"),
        ("six", "6"),
        ("seven", "7"),
        ("eight", "8"),
        ("nine", "9")
        ]);

    let file = File::open("input.txt").unwrap();
    let mut results: Vec<u32> = Vec::new();

    let lines = io::BufReader::new(file).lines();
    for result in lines {
        if let Ok(line) = result {
            let mut digits: Vec<String> = Vec::new();
            for (idx, char) in line.chars().enumerate() {
                if char.is_ascii_digit() {
                    digits.push(char.to_string());
                    continue;
                }
                let (_, rest) = line.split_at(idx);
                for (spelled, digit) in spelled_to_digit_hash.clone() {
                    if rest.starts_with(spelled) {
                        digits.push(digit.to_string())
                    }
                }
            }
            let digit_str = [digits.first().unwrap().to_string(), digits.last().unwrap().to_string()].join("");
            results.push(digit_str.parse::<u32>().unwrap());
        }
    }
    println!("Result: {}", results.into_iter().fold(0, |acc, val| acc + val));
}
