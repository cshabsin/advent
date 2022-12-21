use std::collections::HashMap;
use std::env;
use std::fs;
use std::io;
use std::time::SystemTime;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num2(&input);
    println!("{num}");
    Ok(())
}

fn get_num(input: &str) -> i64 {
    let monkeys = Monkeys::new(input);
    monkeys.eval("root")
}

fn get_num2(input: &str) -> i64 {
    let now = SystemTime::now();
    let mut monkeys = Monkeys::new(input);
    monkeys.equalize_root();
    for i in 0..10000000 {
        monkeys.insert("humn", i);
        if monkeys.eval("root") == 1 {
            return i;
        }
        if i % 10000 == 0 {
            println!("{i} : {}", now.elapsed().unwrap().as_secs());
        }
    }
    -1
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
