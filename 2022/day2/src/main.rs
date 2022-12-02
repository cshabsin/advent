use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();
    let mut f: fn(&str) -> i32 = score;
    if args.len() > 2 {
        f = score2;
    }
    let lines: i32 = fs::read_to_string(&args[1])
        .expect("should have been able to read the file")
        .split("\n")
        .filter(|x| (*x) != "")
        .map(f)
        .sum();
    println!("{lines}")
}

fn score(x: &str) -> i32 {
    let mut iter = x.chars();
    // let opp = iter.next().unwrap();
    iter.next();
    iter.next();
    let me = iter.next().unwrap();
    let mut score = match me {
        'X' => 1,
        'Y' => 2,
        'Z' => 3,
        _ => {
            panic!("unexpected me {me}");
        }
    };
    if x == "A X" {
        score += 3;
    } else if x == "A Y" {
        score += 6;
    } else if x == "B Y" {
        score += 3
    } else if x == "B Z" {
        score += 6;
    } else if x == "C X" {
        score += 6;
    } else if x == "C Z" {
        score += 3;
    }
    score
}

fn score2(x: &str) -> i32 {
    let mut iter = x.chars();
    let opp = iter.next().unwrap();
    iter.next();
    let me = iter.next().unwrap();

    match me {
        'X' =>
        // lose
        {
            match opp {
                'A' => 3,
                'B' => 1,
                'C' => 2,
                _ => {
                    panic!("unknown opp {opp}");
                }
            }
        }
        'Y' =>
        // draw
        {
            match opp {
                'A' => 1 + 3,
                'B' => 2 + 3,
                'C' => 3 + 3,
                _ => {
                    panic!("unknown opp {opp}");
                }
            }
        }
        'Z' =>
        // win
        {
            match opp {
                'A' => 2 + 6,
                'B' => 3 + 6,
                'C' => 1 + 6,
                _ => {
                    panic!("unknown opp {opp}");
                }
            }
        }
        _ => {
            panic!("unknown me {me}");
        }
    }
}
