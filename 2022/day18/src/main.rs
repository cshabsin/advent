use itertools::Itertools;
use std::collections::HashSet;
use std::env;
use std::fs;
use std::io;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num(&input);
    println!("{num}");
    let num2 = get_num2(&input);
    println!("{num2}");
    Ok(())
}

fn get_num(input: &str) -> usize {
    let mut cubes = HashSet::new();
    for line in input.split("\n") {
        cubes.insert(
            line.split(",")
                .map(|d| d.parse().unwrap())
                .next_tuple::<(usize, usize, usize)>()
                .unwrap(),
        );
    }
    let mut cnt = 0;
    for cube in &cubes {
        for (x, y, z) in neighbors(cube.0, cube.1, cube.2) {
            if !cubes.contains(&(x, y, z)) {
                cnt += 1;
            }
        }
    }
    cnt
}

fn get_num2(input: &str) -> usize {
    let mut cubes = HashSet::new();
    for line in input.split("\n") {
        cubes.insert(
            line.split(",")
                .map(|d| d.parse().unwrap())
                .next_tuple::<(usize, usize, usize)>()
                .unwrap(),
        );
    }
    // 0..21 on three sides. each cell is type "exterior" "lava" or "internal".
    // Paint from 0,0,0 to 21,21,21 all spots that can be reached?
    let mut world: [[[Cell; 31]; 31]; 31] = [[[Cell::Unknown; 31]; 31]; 31];
    for y in 0..31 {
        for z in 0..31 {
            traverse(&cubes, &mut world, 0, y, z);
            traverse(&cubes, &mut world, 30, y, z);
        }
        for x in 0..31 {
            traverse(&cubes, &mut world, x, y, 0);
            traverse(&cubes, &mut world, x, y, 30);
        }
    }
    for x in 0..31 {
        for z in 0..31 {
            traverse(&cubes, &mut world, x, 0, z);
            traverse(&cubes, &mut world, x, 30, z);
        }
    }
    traverse(&cubes, &mut world, 0, 0, 0);
    traverse(&cubes, &mut world, 30, 0, 0);
    traverse(&cubes, &mut world, 0, 30, 0);
    traverse(&cubes, &mut world, 30, 30, 0);
    traverse(&cubes, &mut world, 0, 0, 30);
    traverse(&cubes, &mut world, 30, 0, 30);
    traverse(&cubes, &mut world, 0, 30, 30);
    traverse(&cubes, &mut world, 30, 30, 30);
    // once we have the world figured out, just count "Exterior" neighbors.
    let mut cnt = 0;
    for cube in &cubes {
        if cube.0 == 0 {
            cnt += 1; // neighbors won't return -1
        }
        if cube.1 == 0 {
            cnt += 1; // neighbors won't return -1
        }
        if cube.2 == 0 {
            cnt += 1; // neighbors won't return -1
        }
    for (x, y, z) in neighbors(cube.0, cube.1, cube.2) {
            if world[x][y][z] == Cell::Exterior {
                cnt += 1;
            }
        }
    }
    cnt
}

fn neighbors(x: usize, y: usize, z: usize) -> Vec<(usize, usize, usize)> {
    let mut r = Vec::new();
    if x > 0 {
        r.push((x - 1, y, z));
    }
    if y > 0 {
        r.push((x, y - 1, z));
    }
    if z > 0 {
        r.push((x, y, z - 1));
    }
    if x < 30 {
        r.push((x + 1, y, z));
    }
    if y < 30 {
        r.push((x, y + 1, z));
    }
    if z < 30 {
        r.push((x, y, z + 1));
    }
    r
}

fn traverse(
    cubes: &HashSet<(usize, usize, usize)>,
    world: &mut [[[Cell; 31]; 31]; 31],
    x: usize,
    y: usize,
    z: usize,
) {
    if cubes.contains(&(x, y, z)) {
        world[x][y][z] = Cell::Lava;
        return;
    }
    let mut explore_queue = vec![(x, y, z)];
    loop {
        match explore_queue.pop() {
            Some((x, y, z)) => {
                for (x, y, z) in neighbors(x, y, z) {
                    if world[x][y][z] != Cell::Unknown {
                        continue;
                    }
                    if cubes.contains(&(x, y, z)) {
                        world[x][y][z] = Cell::Lava;
                        continue;
                    }
                    world[x][y][z] = Cell::Exterior;
                    explore_queue.push((x, y, z));
                }
            }
            None => return,
        }
    }
}

#[derive(Copy, Clone, PartialEq)]
enum Cell {
    Unknown,
    Exterior,
    Lava,
}

pub const TEST_INPUT: &str = "2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::get_num2;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num("1,1,1\n2,1,1"), 10);
        assert_eq!(get_num(TEST_INPUT), 64);
        assert_eq!(get_num2(TEST_INPUT), 58);
    }
}
