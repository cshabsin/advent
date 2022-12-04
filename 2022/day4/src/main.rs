use std::env;
use std::fs::File;
use std::io;
use std::io::BufRead; // make BufReader available

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let file = File::open(&args[1])?;
    let mut filtfn: fn(&Pair) -> bool = contain;
    if args.len() > 2 {
        filtfn = overlap;
    }
    let cnt = io::BufReader::new(file).lines().map(Pair::from_line).filter(filtfn).count();
    println!("{cnt}");
    Ok(())
}

struct Pair {
    first: Range,
    second: Range,
}

impl Pair {
    fn from_line(line: Result<String, io::Error>) -> Pair {
        let line = line.as_ref().unwrap();
        let mut s = line.split(",");
        Pair {
            first: Range::from_string(s.next().unwrap()),
            second: Range::from_string(s.next().unwrap()),
        }
    }
}

fn contain(pair: &Pair) -> bool {
    pair.first.contains(&pair.second) || pair.second.contains(&pair.first)
}

fn overlap(pair: &Pair) -> bool {
    pair.first.overlaps(&pair.second)
}

struct Range {
    first: i32,
    last: i32,
}

impl Range {
    fn from_string(s: &str) -> Range {
        let mut s = s.split("-");
        Range {
            first: s.next().unwrap().parse::<i32>().unwrap(),
            last: s.next().unwrap().parse::<i32>().unwrap(),
        }
    }    

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

