use itertools::{Chunk, Itertools};
use std::cell::RefCell;
use std::collections::VecDeque;
use std::env;
use std::fmt::Display;
use std::fs;
use std::io;
use std::str::Split;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num(&input);
    println!("{num}");
    Ok(())
}

fn get_num(input: &str) -> usize {
    let monkeys: Vec<Monkey> = input
        .split("\n")
        .chunks(7)
        .into_iter()
        .map(|chunk| Monkey::new(chunk))
        .collect();

    for _ in 0..10000 {
        for m in &monkeys {
            m.inspect(&monkeys);
        }
        }
    let mut i = monkeys.iter().map(|m| m.get_inspection_count()).sorted().rev();

    i.next().unwrap()*i.next().unwrap()
}

struct Monkey {
    items: RefCell<VecDeque<i64>>,
    op: Operation,
    test_div: i64,
    true_tgt: usize,
    false_tgt: usize,

    inpsection_count: RefCell<usize>,
}

impl Monkey {
    fn new(mut chunks: Chunk<Split<&str>>) -> Monkey {
        chunks.next(); // skip "Monkey 0"
        Monkey {
            items: RefCell::new(Monkey::parse_items(chunks.next().unwrap())),
            op: Monkey::parse_operation(chunks.next().unwrap()),
            test_div: Monkey::parse_last(chunks.next().unwrap()),
            true_tgt: Monkey::parse_last(chunks.next().unwrap()) as usize,
            false_tgt: Monkey::parse_last(chunks.next().unwrap()) as usize,
            inpsection_count: RefCell::new(0),
        }
    }

    fn parse_items(s: &str) -> VecDeque<i64> {
        let mut items = VecDeque::new();
        let (_, s) = s.split_at(s.find(":").unwrap() + 2);
        for n in s.split(",") {
            items.push_back(n.trim().parse().unwrap());
        }
        items
    }

    fn parse_operation(s: &str) -> Operation {
        let s = s.strip_prefix("  Operation: new = old ").unwrap();
        let (op, n) = s.split_at(1);
        if n.trim() == "old" {
            return Operation::Square();
        }
        match (op.chars().nth(0).unwrap(), n.trim().parse().unwrap()) {
            ('*', n) => Operation::Mult(n),
            ('+', n) => Operation::Add(n),
            _ => {
                panic!()
            }
        }
    }

    fn parse_last(s: &str) -> i64 {
        s.split(" ").last().unwrap().parse().unwrap()
    }

    fn has_items(&self) -> bool {
        !self.items.borrow().is_empty()
    }

    fn pull_first_item(&self) -> i64 {
        *self.inpsection_count.borrow_mut() += 1;
        self.items.borrow_mut().pop_front().unwrap()
    }

    fn add_item(&self, item: i64) {
        self.items.borrow_mut().push_back(item)
    }

    fn get_inspection_count(&self) -> usize {
        *self.inpsection_count.borrow()
    }

    fn inspect(&self, monkeys: &Vec<Monkey>) {
        while self.has_items() {
            let mut item = self.pull_first_item();
            // item = self.op.apply(item) / 3;
            item = self.op.apply(item) % (2*7*3*17*11*19*5*13);
            let tgt_monkey = {
                if item % self.test_div == 0 {
                    self.true_tgt
                } else {
                    self.false_tgt
                }
            };
            monkeys[tgt_monkey].add_item(item);
        }
    }
}

impl Display for Monkey {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Items: ")?;
        for i in self.items.borrow().iter() {
            write!(f, "{i}, ")?;
        }
        writeln!(f, "\noperation: {}", self.op)?;
        writeln!(
            f,
            "div {}, true {}, false {}",
            self.test_div, self.true_tgt, self.false_tgt
        )
    }
}

#[derive(PartialEq, Debug)]
enum Operation {
    Mult(i64),
    Add(i64),
    Square(),
}

impl Operation {
    fn apply(&self, item: i64) -> i64 {
        match self {
            Operation::Mult(n) => item*n,
            Operation::Add(n) => item+n,
            Operation::Square() => item*item
        }
    }
}

impl Display for Operation {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::Mult(n) => write!(f, "* {n}"),
            Self::Add(n) => write!(f, "+ {n}"),
            Self::Square() => write!(f, "* old"),
        }
    }
}

pub const TEST_INPUT: &str = "";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::Monkey;
    use crate::Operation;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT), 0);
    }

    #[test]
    fn test_parse_operation() {
        assert_eq!(
            Monkey::parse_operation("  Operation: new = old + 8"),
            Operation::Add(8)
        );
        assert_eq!(
            Monkey::parse_operation("  Operation: new = old * 4"),
            Operation::Mult(4)
        );
    }
}
