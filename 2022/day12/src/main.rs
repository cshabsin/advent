use std::collections::HashSet;
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

fn get_num(input: &str) -> u16 {
    let mut b = Board::new(input);
    let visited = HashSet::new();

    0
}

struct Board {
    board: Vec<Vec<u16>>, // heights. row, col indexed.
    position: (usize, usize),
    target: (usize, usize),
}

impl Board {
    fn new(input: &str) -> Board {
        let mut board = Vec::new();
        let mut position = (0, 0);
        let mut target = (0, 0);
        for (row, line) in input.split("\n").enumerate() {
            let mut rowVec = Vec::new();
            for (col, c) in line.chars().enumerate() {
                rowVec.push(match c {
                    'S' => {
                        position = (row, col);
                        0
                    }
                    'E' => {
                        target = (row, col);
                        25  // it's surrounded by x,y,z so just make it 25.
                    }
                    'a'..='z' => (c as u16) - ('a' as u16),
                    other => panic!("unexpected char {other}"),
                });
            }
            board.push(rowVec);
        }
        Board {
            board: board,
            position: position,
            target: target,
        }
    }

    fn neighbors(&self, position: (usize, usize)) -> Vec<(usize, usize)> {
        Vec::new()
    }
}

pub const TEST_INPUT: &str = "";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT), 0);
    }
}
