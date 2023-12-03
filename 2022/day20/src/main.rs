use std::env;
use std::fs;
use std::io;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num2(&input);
    println!("{num}");
    Ok(())
}

fn get_num(input: &str) -> i64 {
    let mut cells = parse_cells(input, 1);
    let mut head = 0;
    // print_cells(&cells, head);
    for i in 0..cells.len() {
        // println!("do_move {i}");
        head = do_move(&mut cells, head, i);
        // print_cells(&cells, head);
    }
    // print_cells(&cells, head);
    let mut total = 0;
    for i in vec![1000, 2000, 3000] {
        total += get_cell(&cells, i);
    }
    total
}

fn get_num2(input: &str) -> i64 {
    let mut cells = parse_cells(input, 811589153);
    let mut head = 0;
    // print_cells(&cells, head);
    for _ in 0..10 {
        for i in 0..cells.len() {
            // println!("doing move {i}");
            head = do_move(&mut cells, head, i);
            // print_cells(&cells, head);
        }
        // println!("after mixing");
        // print_cells(&cells, head);
    }
    // print_cells(&cells, head);
    let mut total = 0;
    for i in vec![1000, 2000, 3000] {
        total += get_cell(&cells, i);
    }
    total
}

fn print_cells(cells: &Vec<Cell>, head: usize) {
    println!("head: {head}");
    let mut i = head;
    loop {
        let cell = cells.get(i).unwrap();
        println!(
            "cell: {} (i={i}) (next={}, prev={})",
            cell.value, cell.next, cell.prev
        );
        i = cells.get(i).unwrap().next;
        if i == head {
            break;
        }
    }
}

fn parse_cells(input: &str, key: i64) -> Vec<Cell> {
    let mut cells = Vec::new();
    for line in input.split("\n") {
        if line == "" {
            continue;
        }
        cells.push(Cell {
            value: line.parse::<i64>().unwrap() * key,
            next: 0,
            prev: 0,
        });
    }
    let l = cells.len();
    for i in 0..l {
        let mut c = cells.get_mut(i).unwrap();
        if i == 0 {
            c.prev = l - 1;
        } else {
            c.prev = i - 1;
        }
        if i == l - 1 {
            c.next = 0;
        } else {
            c.next = i + 1;
        }
    }
    cells
}

fn do_move(cells: &mut Vec<Cell>, head: usize, i: usize) -> usize {
    let val = cells.get(i).unwrap().value;
    let mut head = head;
    if val < 0 {
        // println!("val: {val}");
        let val = (-val as usize) % (cells.len()-1);
        for _ in 0..val {
            // println!("before:");
            // print_cells(cells, head);
            let next = cells.get(i).unwrap().next;
            let prev = cells.get(i).unwrap().prev;
            let prevprev = cells.get(prev).unwrap().prev;
            cells.get_mut(prevprev).unwrap().next = i;
            cells.get_mut(i).unwrap().prev = prevprev;
            cells.get_mut(i).unwrap().next = prev;
            cells.get_mut(prev).unwrap().prev = i;
            cells.get_mut(prev).unwrap().next = next;
            cells.get_mut(next).unwrap().prev = prev;

            // println!("about to set head:");
            // print_cells(cells, head);
            if prev == head {
                head = i;
            } else if i == head {
                head = next;
            }
            // println!("new head:");
            // print_cells(cells, head);
            // println!("---");
        }
    } else {
        let val = (val as usize) % (cells.len()-1);
        for _ in 0..val {
            let next = cells.get(i).unwrap().next;
            let nextnext = cells.get(next).unwrap().next;
            let prev = cells.get(i).unwrap().prev;
            cells.get_mut(i).unwrap().next = nextnext;
            cells.get_mut(i).unwrap().prev = next;
            cells.get_mut(next).unwrap().next = i;
            cells.get_mut(next).unwrap().prev = prev;
            cells.get_mut(prev).unwrap().next = next;
            cells.get_mut(nextnext).unwrap().prev = i;

            if next == head {
                head = i;
            } else if i == head {
                head = next;
            }
        }
    }
    head
}

fn get_cell(cells: &Vec<Cell>, i: usize) -> i64 {
    let i = i % cells.len();
    let zero = cells
        .iter()
        .enumerate()
        .filter(|p| p.1.value == 0)
        .nth(0)
        .unwrap()
        .0;
    let mut index = zero;
    for _ in 0..i {
        index = cells.get(index).unwrap().next;
    }
    cells.get(index).unwrap().value
}

struct Cell {
    value: i64,
    next: usize,
    prev: usize,
}

pub const TEST_INPUT: &str = "1
2
-3
3
-2
0
4";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::get_num2;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT), 3);
    }

    #[test]
    fn it_works2() {
        assert_eq!(get_num2(TEST_INPUT), 1623178306);
    }
}
