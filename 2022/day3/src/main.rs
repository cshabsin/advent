use std::collections::HashSet;
use std::env;
use std::fs::File;
use std::io;
use std::io::BufRead;
use itertools::{Itertools, Chunk};

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1])?;
    if args.len() == 1 {
        let total: u32 = io::BufReader::new(file).lines().map(priority).sum();
        println!("{total}");
    } else {
        let total: u32 = io::BufReader::new(file).lines().into_iter().chunks(3).into_iter().map(priority_chunk).sum();//into_iter().chunks(3).into_iter().filter(priority_chunk).sum();
        println!("{total}");
    }
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
                return value(c);
            }
        }
    }
    panic!("no common entry found!");
}

fn priority_chunk(line_chunk: Chunk<std::io::Lines<std::io::BufReader<File>>>) -> u32 {
    let mut firsts = HashSet::new();
    let mut seconds = HashSet::new();
    for (i, line) in line_chunk.enumerate() {
        let line = line.unwrap();
        if i == 0 {
            line.as_bytes().into_iter().for_each(|c| { firsts.insert(*c); });
        } else if i == 1 {
            line.as_bytes().into_iter().for_each(|c| { seconds.insert(*c); });
        } else {
            return value(line.as_bytes().into_iter().filter(|c| firsts.contains(*c) && seconds.contains(*c)).next().unwrap());
        }
    }
    0
}

fn value(c: &u8) -> u32 {
    if *c <= 64 + 26 {
        return (*c - 65 + 27) as u32;
    }
    (*c - 97 + 1) as u32
}
