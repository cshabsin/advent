use std::collections::HashSet;
use std::collections::VecDeque;
use std::env;
use std::fs;
use std::io;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num(&input);
    println!("{num}");
    Ok(())
}

fn get_num(input: &str) -> usize {
    let mut last_four = VecDeque::new();
    for (i, c) in input.chars().enumerate() {
        if last_four.len() == 4 {
            last_four.pop_front();
        }
        last_four.push_back(c);
        if last_four.len() == 4 && unique(&last_four) {
            return i+1;
        }
    }
    0
}

fn unique(v: &VecDeque<char>) -> bool {
    let mut h = HashSet::new();
    for c in v {
        if h.contains(c) {
            return false;
        }
        h.insert(c);
    }
    return true;
}

pub const TEST_INPUT: &str = "mjqjpqmgbljsphdztnvjfqwrcgsmlb";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT), 7);
    }
}
