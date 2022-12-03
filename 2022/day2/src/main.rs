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

fn pull_chars(line: &str) -> (char, char) {
    match line.chars().collect::<Vec<char>>()[..] {
        [c1, _, c2]  => (c1, c2),
        _ => { panic!("unexpected input {line}"); }
    }
}

fn score(x: &str) -> i32 {
    let (opp, me) = pull_chars(x);
    let opp = RPS::from_opp(opp);
    let me = RPS::from_part1(me);
    let result = GameResult::new(&me, &opp);
    // println!("{x}: opp({opp:?}) vs me({me:?}): {result:?}");
    result.bonus()+me.bonus()
}

fn score2(x: &str) -> i32 {
    let (opp, me) = pull_chars(x);
    let opp = RPS::from_opp(opp);
    let me = RPS::from_part2(&opp, me);
    let result = GameResult::new(&me, &opp);
    // println!("{x}: opp({opp:?}) vs me({me:?}): {result:?}");
    result.bonus()+me.bonus()
}

#[derive(PartialEq, Clone, Copy, Debug)]
enum RPS {
    Rock,
    Paper,
    Scissors,
}

impl RPS {
    fn from_opp(opp: char) -> RPS {
        match opp {
            'A' => RPS::Rock,
            'B' => RPS::Paper,
            'C' => RPS::Scissors,
            _ => {
                panic!("unknown opp {opp}");
            }
        }
    }

    fn from_part1(me: char) -> RPS {
        match me {
            'X' => RPS::Rock,
            'Y' => RPS::Paper,
            'Z' => RPS::Scissors,
            _ => {
                panic!("unknown me {me}");
            }
        }
    }

    fn from_part2(opp: &RPS, me: char) -> RPS {
        match me {
            'X' => {
                // lose
                match opp {
                    RPS::Rock => RPS::Scissors,
                    RPS::Paper => RPS::Rock,
                    RPS::Scissors => RPS::Paper,
                }
            }
            'Y' => *opp, // draw
            'Z' => {
                // win
                match opp {
                    RPS::Rock => RPS::Paper,
                    RPS::Paper => RPS::Scissors,
                    RPS::Scissors => RPS::Rock,
                }
            }
            _ => {
                panic!("unknown me {me}");
            }
        }
    }

    fn bonus(&self) -> i32 {
        match self {
            RPS::Rock => 1,
            RPS::Paper => 2,
            RPS::Scissors => 3,
        }
    }

    fn beats(&self, opp: &RPS) -> bool {
        (*self == RPS::Rock && *opp == RPS::Scissors)
            || (*self == RPS::Paper && *opp == RPS::Rock)
            || (*self == RPS::Scissors && *opp == RPS::Paper)
    }
}

#[derive(Debug)]
enum GameResult {
    Win,
    Draw,
    Loss,
}

impl GameResult {
    fn new(me: &RPS, opp: &RPS) -> GameResult {
        if me == opp {
            GameResult::Draw
        } else if me.beats(opp) {
            GameResult::Win
        } else {
            GameResult::Loss
        }
    }

    fn bonus(&self) -> i32 {
        match self {
            GameResult::Loss => 0,
            GameResult::Draw => 3,
            GameResult::Win => 6,
        }
    }
}
