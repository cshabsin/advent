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

fn get_num(input: &str) -> i32 {
    let mut c = Computer::new(input);
    let mut total = 0;
    loop {
        let (cycle, x) = c.cycle();
        match cycle {
            20 | 60 | 100 | 140 | 180 | 220 => {
                total += x*cycle;
                if cycle == 220 {
                    break;
                }
            },
            _ => {}
        }
    }
    total
}

struct Computer {
    x: i32,
    instructions: Vec<Instruction>,
    cycle: i32,
    ip: usize,       // index into instructions
    progress: usize, // number of cycles spent on current instruction
}

impl Computer {
    fn new(input: &str) -> Computer {
        Computer {
            x: 1,
            instructions: input
                .split("\n")
                .map(|line| Instruction::new(line))
                .flatten()
                .collect(),
            cycle: 0,
            ip: 0,
            progress: 0,
        }
    }

    fn advance_ip(&mut self) {
        self.ip += 1;
        if self.ip >= self.instructions.len() {
            self.ip = 0;
        }
    }

    // cycle runs one cycle of the computer and returns the current cycle number and the value of x during that cycle
    fn cycle(&mut self) -> (i32, i32) {
        let rx = self.x;
        self.cycle += 1;
        match self.instructions[self.ip] {
            Instruction::Noop() => {
                self.advance_ip();
            },
            Instruction::Addx(v) => {
                self.progress += 1;
                if self.progress == 2 {
                    self.progress = 0;
                    self.advance_ip();
                    self.x += v;
                }
            }
        }
        (self.cycle, rx)
    }
}

enum Instruction {
    Addx(i32),
    Noop(),
}

impl Instruction {
    fn new(line: &str) -> Option<Instruction> {
        if line == "" {
            return None;
        }
        let mut tokens = line.split(" ");
        match tokens.next().unwrap() {
            "noop" => Some(Instruction::Noop()),
            "addx" => Some(Instruction::Addx(tokens.next().unwrap().parse().unwrap())),
            unknown => {
                panic!("unknown token {unknown}")
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT), 13140);
    }
}

pub const TEST_INPUT: &str = "addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop";
