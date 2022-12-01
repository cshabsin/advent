use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();
    let contents = fs::read_to_string(&args[1]).expect("should have been able to read the file");

    let mut elves = Vec::new();
    let mut current_elf = 0;
    for line in contents.split("\n") {
        if line == "" {
            elves.push(current_elf);
            current_elf = 0;
            continue;
        }
        current_elf += line.parse::<i32>().unwrap();
    }
    if current_elf != 0 {
        elves.push(current_elf);
    }
    let mut max = 0;
    for elf in &elves {
        if elf > &max {
            max = *elf;
        }
    }
    println!("{max}");
}
