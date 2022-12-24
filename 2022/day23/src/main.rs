use core::fmt;
use std::collections::HashMap;
use std::collections::HashSet;
use std::env;
use std::fs;
use std::io;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num(&input);
    println!("{num}");
    let num = get_num2(&input);
    println!("{num}");
    Ok(())
}

fn get_num(input: &str) -> usize {
    let mut b = Board::new(input);
    let mut dirs = vec![
        Direction::North,
        Direction::South,
        Direction::West,
        Direction::East,
    ];
    println!("== Initial state ==");
    println!("{b}");
    for i in 0..10 {
        (b, _) = b.next(&dirs);
        dirs = vec![
            *dirs.get(1).unwrap(),
            *dirs.get(2).unwrap(),
            *dirs.get(3).unwrap(),
            *dirs.get(0).unwrap(),
        ];
        println!("== End of Round {} ==", i + 1);
        println!("{b}");
    }
    b.ground_tiles()
}

fn get_num2(input: &str) -> usize {
    let mut b = Board::new(input);
    let mut dirs = vec![
        Direction::North,
        Direction::South,
        Direction::West,
        Direction::East,
    ];
    let mut num = 1;
    loop {
        let changed: bool;
        (b, changed) = b.next(&dirs);
        if !changed {
            return num;
        }
        num += 1;
        dirs = vec![
            *dirs.get(1).unwrap(),
            *dirs.get(2).unwrap(),
            *dirs.get(3).unwrap(),
            *dirs.get(0).unwrap(),
        ];
    }
}

struct Board {
    board: HashSet<(i32, i32)>,
}

#[derive(Clone, Copy)]
enum Direction {
    North,
    South,
    West,
    East,
}

impl fmt::Display for Direction {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            Direction::North => write!(f, "N"),
            Direction::South => write!(f, "S"),
            Direction::East => write!(f, "E"),
            Direction::West => write!(f, "W"),
        }
    }
}

impl Board {
    fn new(input: &str) -> Board {
        let mut board = HashSet::new();
        for (row, line) in input.split("\n").enumerate() {
            if line == "" {
                continue;
            }
            for (col, c) in line.chars().enumerate() {
                if c == '#' {
                    board.insert((row as i32, col as i32));
                }
            }
        }
        Board { board }
    }

    fn is_elf(&self, r: i32, c: i32) -> bool {
        self.board.contains(&(r, c))
    }

    fn next(&self, dirs: &Vec<Direction>) -> (Board, bool) {
        let mut proposals = HashMap::new();
        let mut any_changed = false;
        for elf in &self.board {
            let (proposal, changed) = self.proposed_move(elf, dirs);
            any_changed |= changed;
            if !proposals.contains_key(&proposal) {
                proposals.insert(proposal, Vec::new());
            }
            proposals.get_mut(&proposal).unwrap().push(elf);
        }
        let mut board = HashSet::new();
        for (prop, elves) in proposals {
            if elves.len() == 1 {
                board.insert(prop);
            } else {
                for elf in elves {
                    board.insert(*elf); // keep their old position
                }
            }
        }
        (Board { board }, any_changed)
    }

    fn proposed_move(&self, elf: &(i32, i32), dirs: &Vec<Direction>) -> ((i32, i32), bool) {
        let (r, c) = *elf;
        if !(self.is_elf(r - 1, c - 1)
            || self.is_elf(r - 1, c)
            || self.is_elf(r - 1, c + 1)
            || self.is_elf(r, c - 1)
            || self.is_elf(r, c + 1)
            || self.is_elf(r + 1, c - 1)
            || self.is_elf(r + 1, c)
            || self.is_elf(r + 1, c + 1))
        {
            return ((r, c), false);
        }
        for dir in dirs {
            match dir {
                Direction::North => {
                    if !(self.is_elf(r - 1, c - 1)
                        || self.is_elf(r - 1, c)
                        || self.is_elf(r - 1, c + 1))
                    {
                        return ((r - 1, c), true);
                    }
                }
                Direction::South => {
                    if !(self.is_elf(r + 1, c - 1)
                        || self.is_elf(r + 1, c)
                        || self.is_elf(r + 1, c + 1))
                    {
                        return ((r + 1, c), true);
                    }
                }
                Direction::West => {
                    if !(self.is_elf(r - 1, c - 1)
                        || self.is_elf(r, c - 1)
                        || self.is_elf(r + 1, c - 1))
                    {
                        return ((r, c - 1), true);
                    }
                }
                Direction::East => {
                    if !(self.is_elf(r - 1, c + 1)
                        || self.is_elf(r, c + 1)
                        || self.is_elf(r + 1, c + 1))
                    {
                        return ((r, c + 1), true);
                    }
                }
            }
        }
        ((r, c), false)
    }

    fn bounds(&self) -> (i32, i32, i32, i32) {
        let (mut min_r, mut min_c, mut max_r, mut max_c) = (99, 99, 0, 0);
        for (r, c) in &self.board {
            if *r < min_r {
                min_r = *r;
            }
            if *r > max_r {
                max_r = *r;
            }
            if *c < min_c {
                min_c = *c;
            }
            if *c > max_c {
                max_c = *c;
            }
        }
        (min_r, min_c, max_r, max_c)
    }

    fn ground_tiles(&self) -> usize {
        let (min_r, min_c, max_r, max_c) = self.bounds();
        println!("{min_r}-{max_r} x {min_c}-{max_c} | {}", self.board.len());
        ((max_r - min_r + 1) * (max_c - min_c + 1)) as usize - self.board.len()
    }
}

impl fmt::Display for Board {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let (min_r, min_c, max_r, max_c) = self.bounds();
        for r in min_r..=max_r {
            for c in min_c..=max_c {
                if self.is_elf(r, c) {
                    write!(f, "#")?;
                } else {
                    write!(f, ".")?;
                }
            }
            write!(f, "\n")?;
        }
        Ok(())
    }
}

pub const TEST_INPUT: &str = "....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..";

pub const TEST_INPUT2: &str = ".....
..##.
..#..
.....
..##.
.....";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::get_num2;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT), 110);
    }

    #[test]
    fn it_works2() {
        assert_eq!(get_num2(TEST_INPUT), 20);
    }

    #[test]
    fn simple_example() {
        assert_eq!(
            get_num(
                ".....
        ..##.
        ..#..
        .....
        ..##.
        ....."
            ),
            5
        );
    }
}
