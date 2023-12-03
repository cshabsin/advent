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
    let mut max: [i32; 3] = [0,0,0];
    for elf in &elves {
        if elf > &max[0] {
            max[2] = max[1];
            max[1] = max[0];
            max[0] = *elf;
        } else if elf > &max[1] {
            max[2] = max[1];
            max[1] = *elf;
        } else if elf > &max[2] {
            max[2] = *elf;
        }
    }
    let total = max[0]+max[1]+max[2];
    println!("{total}");
}
