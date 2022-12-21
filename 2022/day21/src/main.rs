use std::collections::HashMap;
use std::env;
use std::fs;
use std::io;
use std::num::ParseIntError;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num2(&input);
    println!("{num}");
    Ok(())
}

enum Monkey {
    Val(i64),
    Add((String, String)),
    Subtract((String, String)),
    Multiply((String, String)),
    Divide((String, String)),
    Match((String, String))
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
                    op => panic!("unrecognized operator {op}")
                }
            }
        }
    }
}

fn eval_monkey(monkeys: &HashMap<&str, Monkey>, name: &str) -> i64 {
    match monkeys.get(name).unwrap() {
        Monkey::Val(v) => *v,
        Monkey::Add((a, b)) => eval_monkey(monkeys, &a) + eval_monkey(monkeys, &b),
        Monkey::Subtract((a, b)) => eval_monkey(monkeys, &a) - eval_monkey(monkeys, &b),
        Monkey::Multiply((a, b)) => eval_monkey(monkeys, &a) * eval_monkey(monkeys, &b),
        Monkey::Divide((a, b)) => eval_monkey(monkeys, &a) / eval_monkey(monkeys, &b),
        Monkey::Match((a, b)) => {
            if eval_monkey(monkeys, a) == eval_monkey(monkeys, b) {
                1
            } else {
                0
            }
        }
    }
}

fn get_num(input: &str) -> i64 {
    let mut monkeys = HashMap::new();
    for line in input.split("\n") {
        if line == "" {
            continue;
        }
        let mut split = line.split(": ");
        let name = split.next().unwrap();
        monkeys.insert(name, Monkey::new(split.next().unwrap()));
    }
    eval_monkey(&monkeys, "root")
}

fn get_num2(input: &str) -> i64 {
    let mut monkeys = HashMap::new();
    for line in input.split("\n") {
        if line == "" {
            continue;
        }
        let mut split = line.split(": ");
        let name = split.next().unwrap();
        monkeys.insert(name, Monkey::new(split.next().unwrap()));
    }
    for i in 0..10000000 {
        monkeys.insert("humn", Monkey::Val(i));
        if eval_monkey(&monkeys, "root") == 1 {
            return i;
        }
    }
    -1
}

pub const TEST_INPUT: &str = "root: pppw = sjmn
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
        assert_eq!(get_num2(TEST_INPUT), 301);
    }
}
