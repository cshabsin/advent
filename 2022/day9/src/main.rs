use std::collections::HashMap;
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
    let mut b = Board::new();
    for line in input.split('\n') {
        if line == "" {
            continue;
        }
        b.do_move(line);
    }
    b.board.len()
}

struct Board {
    head: (i32, i32),
    tail: (i32, i32),
    board: HashMap<(i32, i32), bool>,
}

impl Board {
    fn new() -> Board {
        let mut board = HashMap::new();
        board.insert((0, 0), true);
        Board {
            head: (0, 0),
            tail: (0, 0),
            board: board,
        }
    }

    fn do_move(&mut self, mv: &str) {
        let mut it = mv.split(' ');
        let (dir, count) = (
            it.next().unwrap(),
            it.next().unwrap().parse::<usize>().unwrap(),
        );
        let d = match dir {
            "U" => (0, -1),
            "D" => (0, 1),
            "R" => (-1, 0),
            "L" => (1, 0),
            _ => {
                panic!("unknown dir {dir}")
            }
        };
        for _ in 0..count {
            self.do_move_internal(d);
        }
    }

    fn do_move_internal(&mut self, d: (i32, i32)) {
        self.head.0 += d.0;
        self.head.1 += d.1;
        if !self.adjacent() {
            self.tail.0 = self.head.0 - d.0;
            self.tail.1 = self.head.1 - d.1;
        }
        self.board.insert(self.tail, true);
    }

    // return true if head is still adjacent to tail (no need to move tail)
    fn adjacent(&self) -> bool {
        (self.head.0 - self.tail.0).abs() <= 1 && (self.head.1 - self.tail.1).abs() <= 1
    }
}

pub const TEST_INPUT: &str = "R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT), 13);
    }
}
