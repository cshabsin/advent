use std::collections::HashMap;
use std::collections::HashSet;
use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();
    let contents = fs::read_to_string(&args[1]).expect("should have been able to read the file");

    let mut twos: i32 = 0;
    let mut threes: i32 = 0;
    let mut counts = HashMap::<String, Counts>::new();
    for line in contents.split("\n") {
        if line == "" {
            continue;
        }
        let c = Counts::new(line);
        if c.has_count(2) {
            twos += 1
        }
        if c.has_count(3) {
            threes += 1
        }
        counts.insert(line.to_string(), c);
    }
    println!("twos: {twos}, threes: {threes}");
    // pare down the list of options to a manageable set.
    let mut pairs_to_consider = HashSet::new();
    for count in &counts {
        for count2 in &counts {
            if count.1.is_oneoff(count2.1) {
                pairs_to_consider.insert((count.0, count2.0));
            }
        }
    }
    for (s1, s2) in pairs_to_consider {
        let mut found = 0;
        for (i, c1) in s1.chars().enumerate() {
            if s2.chars().nth(i) != Some(c1) {
                found += 1;
                if found > 1 {
                    break;
                }
            }
        }
        if found == 1 {
            println!("oneoff: {s1} vs {s2}");
            let mut s = String::new();
            for (i, c1) in s1.chars().enumerate() {
                if s2.chars().nth(i) == Some(c1) {
                    s.push(c1);
                }
            }
            println!("common: {s}");
        }
    }
}

struct Counts {
    m: HashMap<char, i32>,
}

impl Counts {
    fn new(line: &str) -> Counts {
        let mut letters = HashMap::<char, i32>::new();
        for ch in line.chars() {
            let count = letters.entry(ch).or_insert(0);
            *count += 1;
        }
        return Counts { m: letters };
    }

    fn has_count(&self, c: i32) -> bool {
        for entry in &self.m {
            if *entry.1 == c {
                return true;
            }
        }
        return false;
    }

    fn is_oneoff(&self, other: &Counts) -> bool {
        let mut diffs = 0;
        for entry in &self.m {
            let otherval = other.m.get(&entry.0);
            match otherval {
                Some(otherval) => {
                    if *entry.1 == otherval - 1 || *entry.1 == otherval + 1 {
                        diffs += 1;
                    } else if *entry.1 != *otherval {
                        return false;
                    }
                }
                None => {
                    if *entry.1 == 1 {
                        diffs += 1;
                    } else {
                        return false;
                    }
                }
            }
            if diffs > 2 {
                return false;
            }
        }
        return diffs == 2;
    }
}
