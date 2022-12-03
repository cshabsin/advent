use std::collections::HashSet;
use std::env;
use std::fs::File;
use std::io;
use std::io::BufRead;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1])?;
    let total: u32 = io::BufReader::new(file).lines().map(priority).sum();
    println!("{total}");
    Ok(())
}

fn priority(line: Result<String, std::io::Error>) -> u32 {
    let line = line.unwrap();
    let line = line.as_bytes();
    let l = line.len() / 2;
    let mut firsts = HashSet::new();
    for (i, c) in line.iter().enumerate() {
        if i < l {
            firsts.insert(c);
        } else {
            if firsts.contains(c) {
                return value(*c);
            }
        }
    }
    panic!("no common entry found!");
}

fn value(c: u8) -> u32 {
    if c <= 64 + 26 {
        return (c - 65 + 27) as u32;
    }
    (c - 97 + 1) as u32
}
