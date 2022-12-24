use core::fmt;
use std::collections::HashMap;
use std::env;
use std::fs;
use std::io;
use std::rc::Rc;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num(&input);
    println!("part 1: {num}");
    let num = get_num2(&input);
    println!("part 2: {num}");
    Ok(())
}

fn get_num(input: &str) -> i64 {
    let monkeys = Monkeys::new(input);
    monkeys.eval("root")
}

fn get_num2(input: &str) -> i64 {
    let mut monkeys = Monkeys::new(input);
    monkeys.equalize_root();
    let mut cm = monkeys.to_cellmonkey("root");
    loop {
        println!("cellmonkey: {}", cm.as_ref());
        cm = cm.as_ref().simplify();
        if let CellMonkey::Match((left, right)) = cm.as_ref() {
            if matches!(left.as_ref(), CellMonkey::Human) {
                match right.as_ref() {
                    CellMonkey::Val(v) => return *v,
                    _ => panic!("not a val?")
                }
            }
        }
    }
}

enum CellMonkey {
    Human,
    Val(i64),
    Add((Rc<CellMonkey>, Rc<CellMonkey>)),
    Subtract((Rc<CellMonkey>, Rc<CellMonkey>)),
    Multiply((Rc<CellMonkey>, Rc<CellMonkey>)),
    Divide((Rc<CellMonkey>, Rc<CellMonkey>)),
    Match((Rc<CellMonkey>, Rc<CellMonkey>)),
}

impl CellMonkey {
    fn simplify(&self) -> Rc<CellMonkey> {
        match self {
            CellMonkey::Match((left, right)) => {
                let val = match right.as_ref() {
                    CellMonkey::Val(v) => v,
                    _ => panic!("right should be val"),
                };
                match left.as_ref() {
                    CellMonkey::Human => panic!("don't call me, you're already human. {val}"),
                    CellMonkey::Val(other) => panic!("two vals? {other} = {val}"),
                    CellMonkey::Match(_) => panic!("match???"),
                    CellMonkey::Add((a, b)) => {
                        if let CellMonkey::Val(addend) = a.as_ref() {
                            Rc::new(CellMonkey::Match((
                                b.clone(),
                                CellMonkey::val(*val - *addend),
                            )))
                        } else if let CellMonkey::Val(addend) = b.as_ref() {
                            Rc::new(CellMonkey::Match((
                                a.clone(),
                                CellMonkey::val(*val - *addend),
                            )))
                        } else {
                            panic!("add of two non-vals?")
                        }
                    }
                    CellMonkey::Subtract((a, b)) => {
                        if let CellMonkey::Val(addend) = a.as_ref() {
                            // a-(blahblah) = val -> blahblah = a-val
                            Rc::new(CellMonkey::Match((
                                b.clone(),
                                CellMonkey::val(*addend - *val),
                            )))
                        } else if let CellMonkey::Val(addend) = b.as_ref() {
                            // (blahblah) - b = val -> blahblah = val+b
                            Rc::new(CellMonkey::Match((
                                a.clone(),
                                CellMonkey::val(*val + *addend),
                            )))
                        } else {
                            panic!("add of two non-vals?")
                        }
                    }
                    CellMonkey::Multiply((a, b)) => {
                        if let CellMonkey::Val(mult) = a.as_ref() {
                            Rc::new(CellMonkey::Match((
                                b.clone(),
                                CellMonkey::val(*val / *mult),
                            )))
                        } else if let CellMonkey::Val(mult) = b.as_ref() {
                            Rc::new(CellMonkey::Match((
                                a.clone(),
                                CellMonkey::val(*val / *mult),
                            )))
                        } else {
                            panic!("multiply of two non-vals?")
                        }
                    }
                    CellMonkey::Divide((a, b)) => {
                        if let CellMonkey::Val(mult) = a.as_ref() {
                            Rc::new(CellMonkey::Match((
                                b.clone(),
                                CellMonkey::val(*val * *mult),
                            )))
                        } else if let CellMonkey::Val(mult) = b.as_ref() {
                            Rc::new(CellMonkey::Match((
                                a.clone(),
                                CellMonkey::val(*val * *mult),
                            )))
                        } else {
                            panic!("divide of two non-vals?")
                        }
                    }
                }
            }
            _ => panic!("only simply a match"),
        }
    }

    fn val(v: i64) -> Rc<CellMonkey> {
        Rc::new(CellMonkey::Val(v))
    }
}

impl fmt::Display for CellMonkey {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            CellMonkey::Add((a, b)) => write!(f, "({a} + {b})"),
            CellMonkey::Subtract((a, b)) => write!(f, "({a} - {b})"),
            CellMonkey::Multiply((a, b)) => write!(f, "({a} * {b})"),
            CellMonkey::Divide((a, b)) => write!(f, "({a} / {b})"),
            CellMonkey::Match((a, b)) => write!(f, "({a} = {b})"),
            CellMonkey::Val(v) => write!(f, "{v}"),
            CellMonkey::Human => write!(f, "human"),
        }
    }
}

enum Monkey {
    Val(i64),
    Add((String, String)),
    Subtract((String, String)),
    Multiply((String, String)),
    Divide((String, String)),
    Match((String, String)),
}

impl Monkey {
    fn new(line: &str) -> Monkey {
        match line.parse::<i64>() {
            Ok(v) => Monkey::Val(v),
            Err(_) => {
                let mut tokens = line.split(" ");
                let first = tokens.next().unwrap().to_string();
                let op = tokens.next().unwrap();
                let second = tokens.next().unwrap().to_string();
                match op {
                    "+" => Monkey::Add((first, second)),
                    "-" => Monkey::Subtract((first, second)),
                    "*" => Monkey::Multiply((first, second)),
                    "/" => Monkey::Divide((first, second)),
                    "=" => Monkey::Match((first, second)),
                    op => panic!("unrecognized operator {op}"),
                }
            }
        }
    }
}

struct Monkeys {
    monkeys: HashMap<String, Monkey>,
}

impl Monkeys {
    fn new(input: &str) -> Monkeys {
        let mut monkeys = HashMap::new();
        for line in input.split("\n") {
            if line == "" {
                continue;
            }
            let mut split = line.split(": ");
            let name = split.next().unwrap();
            monkeys.insert(name.to_string(), Monkey::new(split.next().unwrap()));
        }
        Monkeys { monkeys }
    }

    fn eval(&self, name: &str) -> i64 {
        match self.monkeys.get(name).unwrap() {
            Monkey::Val(v) => *v,
            Monkey::Add((a, b)) => self.eval(&a) + self.eval(&b),
            Monkey::Subtract((a, b)) => self.eval(&a) - self.eval(&b),
            Monkey::Multiply((a, b)) => self.eval(&a) * self.eval(&b),
            Monkey::Divide((a, b)) => self.eval(&a) / self.eval(&b),
            Monkey::Match((a, b)) => {
                if self.eval(a) == self.eval(b) {
                    1
                } else {
                    0
                }
            }
        }
    }

    fn equalize_root(&mut self) {
        let pair = match self.monkeys.get("root").unwrap() {
            Monkey::Add(p) => p,
            Monkey::Subtract(p) => p,
            Monkey::Multiply(p) => p,
            Monkey::Divide(p) => p,
            Monkey::Match(p) => p,
            Monkey::Val(v) => panic!("unexpected Val node at root ({v})"),
        };
        self.monkeys.insert(
            "root".to_string(),
            Monkey::Match((pair.0.to_string(), pair.1.to_string())),
        );
    }

    fn insert(&mut self, name: &str, val: i64) {
        self.monkeys.insert(name.to_string(), Monkey::Val(val));
    }

    fn left(&self, name: &str) -> &str {
        match self.monkeys.get(name).unwrap() {
            Monkey::Add(p) => &p.0,
            Monkey::Subtract(p) => &p.0,
            Monkey::Multiply(p) => &p.0,
            Monkey::Divide(p) => &p.0,
            Monkey::Match(p) => &p.0,
            Monkey::Val(v) => panic!("unexpected Val node in left ({v})"),
        }
    }

    fn right(&self, name: &str) -> &str {
        match self.monkeys.get(name).unwrap() {
            Monkey::Add(p) => &p.1,
            Monkey::Subtract(p) => &p.1,
            Monkey::Multiply(p) => &p.1,
            Monkey::Divide(p) => &p.1,
            Monkey::Match(p) => &p.1,
            Monkey::Val(v) => panic!("unexpected Val node in right ({v})"),
        }
    }

    fn has_human(&self, name: &str) -> bool {
        if name == "humn" {
            return true;
        }
        match self.monkeys.get(name).unwrap() {
            Monkey::Val(_) => false,
            Monkey::Add(p) => self.has_human(&p.0) || self.has_human(&p.1),
            Monkey::Subtract(p) => self.has_human(&p.0) || self.has_human(&p.1),
            Monkey::Multiply(p) => self.has_human(&p.0) || self.has_human(&p.1),
            Monkey::Divide(p) => self.has_human(&p.0) || self.has_human(&p.1),
            Monkey::Match(p) => self.has_human(&p.0) || self.has_human(&p.1),
        }
    }

    fn to_cellmonkey(&self, name: &str) -> Rc<CellMonkey> {
        if !self.has_human(name) {
            return Rc::new(CellMonkey::Val(self.eval(name)));
        }
        if name == "humn" {
            return Rc::new(CellMonkey::Human);
        }
        match self.monkeys.get(name).unwrap() {
            Monkey::Val(_) => panic!("how?"),
            Monkey::Add(p) => Rc::new(CellMonkey::Add((
                self.to_cellmonkey(&p.0),
                self.to_cellmonkey(&p.1),
            ))),
            Monkey::Subtract(p) => Rc::new(CellMonkey::Subtract((
                self.to_cellmonkey(&p.0),
                self.to_cellmonkey(&p.1),
            ))),
            Monkey::Multiply(p) => Rc::new(CellMonkey::Multiply((
                self.to_cellmonkey(&p.0),
                self.to_cellmonkey(&p.1),
            ))),
            Monkey::Divide(p) => Rc::new(CellMonkey::Divide((
                self.to_cellmonkey(&p.0),
                self.to_cellmonkey(&p.1),
            ))),
            Monkey::Match(p) => Rc::new(CellMonkey::Match((
                self.to_cellmonkey(&p.0),
                self.to_cellmonkey(&p.1),
            ))),
        }
    }
}

pub const TEST_INPUT: &str = "root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::get_num2;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT), 152);
    }

    #[test]
    fn it_works2() {
        assert_eq!(get_num2(TEST_INPUT), 301);
    }
}
