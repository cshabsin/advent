use std::collections::HashSet;
use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();

    let file_path = &args[1];
    let contents = fs::read_to_string(file_path).expect("should have been able to read the file");

    let mut total: i32 = 0;
    let mut seen = HashSet::new();
    let mut found = false;
    while !found {
        for line in contents.split("\n") {
            if line == "" {
                continue;
            }
            total = total + line.parse::<i32>().unwrap();
            if seen.contains(&total) {
                println!("{total}");
                found = true;
                break
            }
            seen.insert(total);
            println!("added {total} to set");
        }
    }
}
