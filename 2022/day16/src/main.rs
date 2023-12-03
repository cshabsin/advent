use pathfinding::prelude::dijkstra;
use std::collections::HashMap;
use std::collections::HashSet;
use std::env;
use std::fs;
use std::hash::{Hash, Hasher};
use std::io;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num(&input);
    println!("{num}");
    Ok(())
}

fn get_num(input: &str) -> i32 {
    let g = Graph::new(input);
    let start = GameState::start();
    let result = dijkstra(&start, |p| { println!("handling {}", p.pos); g.get_successors(p) }, |p| p.done());
    result.unwrap().1
}

struct Graph {
    nodes: HashMap<String, Node>,
}

impl Graph {
    fn new(input: &str) -> Graph {
        let mut g = Graph {
            nodes: HashMap::new(),
        };
        for line in input.split("\n") {
            if line == "" {
                continue;
            }
            let node = Node::new(line);
            g.nodes.insert(node.name.clone(), node);
        }
        g
    }

    fn get_successors(&self, prev: &GameState) -> Vec<(GameState, i32)> {
        let mut succ = Vec::new();
        let n = self.nodes.get(&prev.pos).unwrap();
        if prev.used_nodes.contains(&prev.pos) && n.rate != 0 {
            let mut used_nodes = prev.used_nodes.clone();
            used_nodes.insert(prev.pos.to_string());
            succ.push((
                GameState {
                    pos: prev.pos.to_string(),
                    used_nodes: used_nodes,
                    moves: prev.moves + 1,
                },
                -n.rate,
            ));
        }
        for tunn in &n.tunnels {
            succ.push((
                GameState {
                    pos: tunn.clone(),
                    used_nodes: prev.used_nodes.clone(),
                    moves: prev.moves + 1,
                },
                0,
            ))
        }
        succ
    }
}

#[derive(Clone, PartialEq, Eq)]
struct GameState {
    pos: String,
    used_nodes: HashSet<String>,
    moves: usize,
}

impl Hash for GameState {
    fn hash<H: Hasher>(&self, state: &mut H) {
        self.pos.hash(state);
        self.moves.hash(state);
        let mut nodes = self.used_nodes.iter().collect::<Vec<&String>>();
        nodes.sort();
        for s in nodes.iter() {
            s.hash(state);
        }
    }
}

impl GameState {
    fn start() -> GameState {
        GameState {
            pos: "AA".to_string(),
            used_nodes: HashSet::new(),
            moves: 0,
        }
    }

    fn done(&self) -> bool {
        self.moves == 30
    }
}

struct Node {
    name: String,
    rate: i32,
    tunnels: Vec<String>,
}

// node

impl Node {
    fn new(line: &str) -> Node {
        let (name, line) = line.strip_prefix("Valve ").unwrap().split_at(2);
        let mut split = line.strip_prefix(" has flow rate=").unwrap().split(";");
        let rate = split.next().unwrap().parse().unwrap();
        
        let tunnels = split
            .next()
            .unwrap()
            .strip_prefix(" tunnels lead to valves ")
            .unwrap()
            .split(", ")
            .map(String::from)
            .collect();
        Node {
            name: name.to_string(),
            rate: rate,
            tunnels: tunnels,
        }
    }
}

pub const TEST_INPUT: &str = "Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valves GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnels lead to valves II";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT), 0);
    }
}
