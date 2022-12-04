use std::env;
use std::fs::File;
use std::io;
use std::io::BufRead; // make BufReader available

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1])?;
    let mut filtfn : fn(&Result<String, io::Error>) -> bool = contain;
    if args.len() > 2 {
        filtfn = overlap;
    }
    let cnt = io::BufReader::new(file).lines().filter(filtfn).count();
    println!("{cnt}");
    Ok(())
}

fn contain(line: &Result<String, io::Error>) -> bool {
    let line = line.as_ref().unwrap();
    let mut s = line.split(",");
    let first = parse_range(s.next().unwrap());
    let second = parse_range(s.next().unwrap());

    first.contains(&second) || second.contains(&first)
}

fn overlap(line: &Result<String, io::Error>) -> bool {
    let line = line.as_ref().unwrap();
    let mut s = line.split(",");
    let first = parse_range(s.next().unwrap());
    let second = parse_range(s.next().unwrap());

    first.overlaps(&second)
}

struct Range {
    first: i32,
    last: i32,
}

impl Range {
    fn contains(&self, other: &Range) -> bool {
        self.first <= other.first && self.last >= other.last
    }

    fn overlaps(&self, other: &Range) -> bool {
        if self.first <= other.first {
            return self.last >= other.first;
        } else {
            return other.overlaps(self);
        }
    }
}

fn parse_range(s: &str) -> Range {
    let mut s = s.split("-");
    Range {
        first: s.next().unwrap().parse::<i32>().unwrap(), 
        last: s.next().unwrap().parse::<i32>().unwrap(),
    }
}