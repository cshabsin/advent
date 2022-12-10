use std::env;
use std::fmt;
use std::fs;
use std::io;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let val = get_val(&input);
    println!("{val}");
    Ok(())
}

fn get_val(input: &str) -> String {
    let input: Vec<&str> = input.split("\n").collect();
    let mid = input.iter().position(|x| *x == "").unwrap();
    let (board, moves) = input.split_at(mid);
    let mut board = Board::new(board);
    let moves = Moves::new(moves);
    for mv in moves.moves {
        board.do_move(&mv);
    }
    let mut out = String::new();
    for s in board.stacks {
        out.push(s.0.chars().last().unwrap());
    }
    out
}

struct Board {
    stacks: Vec<Stack>,
}

impl Board {
    // TODO(cshabsin): try to learn how to make this take a generic line iterator?
    fn new(input: &[&str]) -> Board {
        let mut b = Board { stacks: Vec::new() };
        for line in input {
            line.as_bytes()
                .chunks(4)
                .enumerate()
                .map(|it| b.prepend(it))
                .last();
        }
        b
    }

    fn prepend(&mut self, it: (usize, &[u8])) {
        let (i, c) = (it.0, it.1[1]);
        if !c.is_ascii_alphabetic() {
            return;
        }
        while self.stacks.len() <= i {
            self.stacks.push(Stack(String::new()));
        }
        self.stacks[i].push_c(c);
    }

    fn do_move(&mut self, mv: &Move) {
        let hold = self.stacks.get_mut(mv.from).unwrap().pop(mv.count);
        let hold: String = hold.chars().rev().collect();
        self.stacks.get_mut(mv.to).unwrap().push(&hold);
    }
}

impl fmt::Display for Board {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        for s in &self.stacks {
            writeln!(f, "{s}")?;
        }
        Ok(())
    }
}

struct Stack(String);

impl fmt::Display for Stack {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.0)
    }
}

impl Stack {
    fn push_c(&mut self, c: u8) {
        self.0.insert(0, c as char)
    }

    fn push(&mut self, s: &str) {
        self.0 += s;
    }

    fn pop(&mut self, cnt: usize) -> String {
        let (start, end) = self.0.split_at(self.0.len() - cnt);
        let rc = end.to_string();
        self.0 = start.to_string();
        rc
    }
}

pub const TEST_INPUT: &str = "    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";

#[cfg(test)]
mod tests {
    use crate::get_val;
    use crate::Stack;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_val(TEST_INPUT), "CMZ");
    }

    #[test]
    fn stack_pop_works() {
        let mut s = Stack("ABC".to_string());
        assert_eq!(s.pop(1), "C");
        assert_eq!(s.0, "AB");
    }
}
struct Moves {
    moves: Vec<Move>,
}

impl Moves {
    fn new(input: &[&str]) -> Moves {
        let mut v = Vec::new();
        for line in input {
            if *line == "" {
                continue;
            }
            let mut fields = line.split(" ");
            v.push(Move {
                count: fields.nth(1).unwrap().parse::<usize>().unwrap(),
                from: fields.nth(1).unwrap().parse::<usize>().unwrap() - 1,
                to: fields.nth(1).unwrap().parse::<usize>().unwrap() - 1,
            });
        }
        Moves { moves: v }
    }
}

struct Move {
    count: usize,
    from: usize,
    to: usize,
}
