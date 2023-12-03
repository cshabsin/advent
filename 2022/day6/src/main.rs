use std::collections::HashSet;
use std::collections::VecDeque;
use std::env;
use std::fs;
use std::io;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let (num1, num2) = (get_num(&input, 4), get_num(&input, 14));
    println!("{num1}, {num2}");
    Ok(())
}

fn get_num(input: &str, n: usize) -> usize {
    let mut last_n = VecDeque::new();
    for (i, c) in input.chars().enumerate() {
        if last_n.len() == n {
            last_n.pop_front();
        }
        last_n.push_back(c);
        if last_n.len() == n && unique(&last_n) {
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
        assert_eq!(get_num(TEST_INPUT, 4), 7);
        assert_eq!(get_num(TEST_INPUT, 14), 19);
    }
}
