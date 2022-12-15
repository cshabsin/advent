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

fn get_num(input: &str) -> u16 {
    0
}

pub const TEST_INPUT: &str = "";

#[cfg(test)]
mod tests {
    use crate::TEST_INPUT;
    use crate::get_num;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT), 0);
    }
}