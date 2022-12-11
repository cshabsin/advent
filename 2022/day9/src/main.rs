use std::collections::HashMap;
use std::env;
use std::fmt::Display;
use std::fs;
use std::io;

const VERBOSE: bool = false;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num(&input, 9);
    println!("{num}");
    Ok(())
}

fn get_num(input: &str, num_tails: i32) -> usize {
    let mut b = Board::new(num_tails);
    for (n, line) in input.split('\n').enumerate() {
        if line == "" {
            continue;
        }
        if VERBOSE {
            println!("--- {n} {line}:");
        }
        b.do_move(line);
    }
    b.board.len()
}

struct Board {
    snake: Vec<(i32, i32)>,
    board: HashMap<(i32, i32), bool>,
    minx: i32,
    miny: i32,
    maxx: i32,
    maxy: i32,
}

impl Board {
    fn new(num_tails: i32) -> Board {
        let mut board = HashMap::new();
        board.insert((0, 0), true);
        let mut snake = Vec::new();
        for _ in 0..(num_tails + 1) {
            snake.push((0, 0));
        }
        Board {
            snake: snake,
            board: board,
            minx: 0,
            miny: 0,
            maxx: 0,
            maxy: 0,
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
            "R" => (1, 0),
            "L" => (-1, 0),
            _ => {
                panic!("unknown dir {dir}")
            }
        };
        for _ in 0..count {
            self.do_move_internal(d);
        }
    }

    fn do_move_internal(&mut self, d: (i32, i32)) {
        self.snake.first_mut().unwrap().0 += d.0;
        self.snake.first_mut().unwrap().1 += d.1;

        let mut iter = self.snake.iter_mut();
        let mut prev = iter.next().unwrap();
        loop {
            let next = iter.next();
            if next.is_none() {
                break;
            }
            let next = next.unwrap();
            if (next.0 - prev.0).abs() <= 1 && (next.1 - prev.1).abs() <= 1 {
                // still next to head, no movement.
            } else if next.0 < prev.0 - 1 && next.1 == prev.1 {
                next.0 = prev.0 - 1;
                next.1 = prev.1;
            } else if next.0 > prev.0 + 1 && next.1 == prev.1 {
                next.0 = prev.0 + 1;
                next.1 = prev.1;
            } else if next.1 < prev.1 - 1 && next.0 == prev.0 {
                next.1 = prev.1 - 1;
            } else if next.1 > prev.1 + 1 && next.0 == prev.0 {
                next.1 = prev.1 + 1;
            } else if next.0 < prev.0 && next.1 < prev.1 {
                // If the head and tail aren't touching and aren't in the same row or column,
                // the tail always moves one step diagonally to keep up.
                next.0 += 1;
                next.1 += 1;
            } else if next.0 < prev.0 && next.1 > prev.1 {
                next.0 += 1;
                next.1 -= 1;
            } else if next.0 > prev.0 && next.1 < prev.1 {
                next.0 -= 1;
                next.1 += 1;
            } else if next.0 > prev.0 && next.1 > prev.1 {
                next.0 -= 1;
                next.1 -= 1;
            }
            prev = next;
        }
        self.board.insert(*prev, true);
        self.expand();
        if VERBOSE {
            println!("{}", self);
        }
    }

    fn expand(&mut self) {
        for knot in &self.snake {
            if knot.0 < self.minx {
                self.minx = knot.0;
            } else if knot.0 > self.maxx {
                self.maxx = knot.0;
            }
            if knot.1 < self.miny {
                self.miny = knot.1;
            } else if knot.1 > self.maxy {
                self.maxy = knot.1;
            }
        }
    }
}

impl Display for Board {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut board = Vec::new();
        for _ in self.miny..(self.maxy + 1) {
            let mut row = Vec::new();
            for _ in self.minx..(self.maxx + 1) {
                row.push('.');
            }
            board.push(row);
        }
        for i in &self.board {
            *board
                .get_mut((i.0 .1 - self.miny) as usize)
                .unwrap()
                .get_mut((i.0 .0 - self.minx) as usize)
                .unwrap() = '#';
        }
        *board
            .get_mut((-self.miny) as usize)
            .unwrap()
            .get_mut((-self.minx) as usize)
            .unwrap() = 's';
        for i in self.snake.iter().enumerate().rev() {
            let mut c = ((48 + i.0) as u8) as char;
            if i.0 == 0 {
                c = 'H';
            }
            *board
                .get_mut((i.1 .1 - self.miny) as usize)
                .unwrap()
                .get_mut((i.1 .0 - self.minx) as usize)
                .unwrap() = c;
        }
        for row in board {
            for col in row {
                write!(f, "{col}")?;
            }
            write!(f, "\n")?;
        }
        Ok(())
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

pub const TEST_INPUT9: &str = "R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::TEST_INPUT;
    use crate::TEST_INPUT9;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT, 1), 13);
    }

    #[test]
    fn it_works9() {
        assert_eq!(get_num(TEST_INPUT9, 9), 36);
    }
}
