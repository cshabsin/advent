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
    Ok(())
}

fn get_num(input: &str) -> usize {
    let mut cubes = HashSet::new();
    for line in input.split("\n") {
        cubes.insert(
            line.split(",")
                .map(|d| d.parse().unwrap())
                .next_tuple::<(i32, i32, i32)>()
                .unwrap(),
        );
    }
    let mut cnt = 0;
    for cube in &cubes {
        let (x, y, z) = (cube.0, cube.1, cube.2);
        if !cubes.contains(&(x+1,y,z)) {
            cnt +=1;
        }
        if !cubes.contains(&(x-1,y,z)) {
            cnt +=1;
        }
        if !cubes.contains(&(x,y+1,z)) {
            cnt +=1;
        }
        if !cubes.contains(&(x,y-1,z)) {
            cnt +=1;
        }
        if !cubes.contains(&(x,y,z+1)) {
            cnt +=1;
        }
        if !cubes.contains(&(x,y,z-1)) {
            cnt +=1;
        }
    }
    cnt
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
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num("1,1,1\n2,1,1"), 10);
        assert_eq!(get_num(TEST_INPUT), 64);
    }
}
