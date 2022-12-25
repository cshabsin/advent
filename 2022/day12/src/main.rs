use std::collections::BinaryHeap;
use std::collections::HashMap;
use std::env;
use std::fs;
use std::io;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num(&input);
    match num {
        Some(num) => println!("{num}"),
        None => println!("no path found"),
    }
    Ok(())
}

fn get_num(input: &str) -> Option<usize> {
    let b = Board::new(input);
    let mut dist = HashMap::new();
    for (r, row) in b.board.iter().enumerate() {
        for (c, _) in row.iter().enumerate() {
            dist.insert((r, c), usize::MAX);
        }
    }

    let mut heap = BinaryHeap::new();
    heap.push(State {
        cost: 0,
        position: b.position,
        history: vec![b.position],
    });
    while let Some(State {
        cost,
        position,
        history,
    }) = heap.pop()
    {
        println!("delving position {}, {} (cost {})", position.0, position.1, cost);
        if position == b.target {
            println!("history: ");
            for p in history {
                println!("{}, {}", p.0, p.1);
            }
            return Some(cost);
        }
        if cost > dist[&position] {
            continue;
        }
        for neighbor in b.reachable_neighbors(position) {
            let mut new_hist: Vec<_> = history.iter().map(|p| (p.0, p.1)).collect();
            new_hist.push(neighbor);

            let next = State {
                cost: cost + 1,
                position: neighbor,
                history: new_hist,
            };
            if next.cost < dist[&neighbor] {
                heap.push(next);
                *dist.get_mut(&neighbor).unwrap() = cost + 1;
            }
        }
    }
    None
}

#[derive(PartialEq, Eq)]
struct State {
    cost: usize,
    position: (usize, usize),
    history: Vec<(usize, usize)>,
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        other
            .cost
            .cmp(&self.cost)
            .then_with(|| self.position.0.cmp(&other.position.0))
            .then_with(|| self.position.1.cmp(&other.position.1))
    }
}

impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(self.cmp(other))
    }
}

struct Board {
    board: Vec<Vec<i16>>, // heights. row, col indexed.
    position: (usize, usize),
    target: (usize, usize),
}

impl Board {
    fn new(input: &str) -> Board {
        let mut board = Vec::new();
        let mut position = None;
        let mut target = None;
        for (row, line) in input.split("\n").enumerate() {
            let mut row_vec = Vec::new();
            for (col, c) in line.chars().enumerate() {
                row_vec.push(match c {
                    'S' => {
                        position = Some((row, col));
                        0
                    }
                    'E' => {
                        target = Some((row, col));
                        25
                    }
                    'a'..='z' => (c as i16) - ('a' as i16),
                    other => panic!("unexpected char {other}"),
                });
            }
            board.push(row_vec);
        }
        Board {
            board: board,
            position: position.unwrap(),
            target: target.unwrap(),
        }
    }

    fn altitude(&self, position: (usize, usize)) -> i16 {
        *self
            .board
            .get(position.0)
            .unwrap_or(&Vec::new())
            .get(position.1)
            .unwrap_or(&1000)
    }

    fn reachable_neighbors(&self, position: (usize, usize)) -> Vec<(usize, usize)> {
        let mut res = Vec::new();
        let current_altitude = self.altitude(position);
        if position.0 != 0 {
            self.maybe_add(&mut res, (position.0 - 1, position.1), current_altitude);
        }
        if position.1 != 0 {
            self.maybe_add(&mut res, (position.0, position.1 - 1), current_altitude);
        }
        self.maybe_add(&mut res, (position.0 + 1, position.1), current_altitude);
        self.maybe_add(&mut res, (position.0, position.1 + 1), current_altitude);
        res
    }

    fn maybe_add(&self, res: &mut Vec<(usize, usize)>, target_pos: (usize, usize), current_altitude: i16) {
        if self.altitude(target_pos) <= current_altitude + 1 {
            res.push(target_pos);
        }
    }
}

pub const TEST_INPUT: &str = "Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT).unwrap(), 31);
    }
}
