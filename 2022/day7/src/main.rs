use std::cell::RefCell;
use std::env;
use std::fs;
use std::io;
use std::rc::Rc;
use std::rc::Weak;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num(&input);
    println!("{num}");
    Ok(())
}

fn get_num(input: &str) -> usize {
    let mut root = Rc::new(RefCell::new(Directory::new("/", None)));
    parse_input(input, root);
    0
}

fn parse_input(input: &str, root: Rc<RefCell<Directory>>) {
    let mut cwd = root;
    for line in input.split("\n") {
        match InputLine::new(line) {
            InputLine::Root() => (),
            InputLine::Ls() => (),
            InputLine::Dir(name) => cwd.borrow_mut().add_dir(&name),
            InputLine::File(name, size) => cwd.borrow_mut().add_file(&name, size),
            InputLine::ChdirUp() => {
                let new_cwd = cwd.borrow().parent().unwrap();
                cwd = new_cwd
            } // crash if cd .. in root
            InputLine::Chdir(dir) => {
                let new_cwd = cwd.borrow().child(&dir);
                cwd = new_cwd
            }
        }
    }
}

enum InputLine {
    Root(),
    Ls(),
    Dir(String),
    File(String, usize),
    Chdir(String),
    ChdirUp(),
}

impl InputLine {
    fn new(line: &str) -> InputLine {
        if line == "$ cd /" {
            InputLine::Root()
        } else if line == "$ ls" {
            InputLine::Ls()
        } else if line == "$ cd .." {
            InputLine::ChdirUp()
        } else if line.starts_with("cd ") {
            InputLine::Chdir(line.split_at(3).1.to_string())
        } else if line.starts_with("dir ") {
            InputLine::Dir(line.split_at(4).1.to_string())
        } else {
            let mut split = line.split(" ");
            let (size, name) = (
                split.next().unwrap().parse().unwrap(),
                split.next().unwrap(),
            );
            InputLine::File(name.to_string(), size)
        }
    }
}

struct Directory {
    name: String,
    subdirs: Vec<Rc<RefCell<Directory>>>,
    files: Vec<File>,
    parent_link: Option<Weak<RefCell<Directory>>>,
}

// impl Rc<Directory> {
//     fn add_dir(self, name: &str) {
//         self.subdirs.push(Directory::new(name, Some<self>))
//     }
// }

impl Directory {
    fn new(name: &str, parent: Option<&Rc<RefCell<Directory>>>) -> Directory {
        Directory {
            name: name.to_string(),
            subdirs: Vec::new(),
            files: Vec::new(),
            parent_link: match parent {
                Some(parent) => Some(Rc::<RefCell<Directory>>::downgrade(parent)),
                None => None,
            },
        }
    }

    fn add_file(&mut self, name: &str, size: usize) {
        self.files.push(File {
            name: name.to_string(),
            size,
        })
    }

    fn size(&self) -> usize {
        let mut size = 0;
        for dir in &self.subdirs {
            size += dir.borrow().size();
        }
        for f in &self.files {
            size += f.size;
        }
        size
    }

    fn child(&self, name: &str) -> Rc<RefCell<Directory>> {
        for subdir in &self.subdirs {
            if subdir.borrow().name == name {
                return subdir.clone();
            }
        }
        panic!("No directory {name} found in {}", self.name);
    }

    fn parent(&self) -> Option<Rc<RefCell<Directory>>> {
        match &self.parent_link {
            None => None,
            Some(parent_link) => parent_link.upgrade(),
        }
    }

    fn add_dir(&mut self, name: &str) {
        
    }
}

struct File {
    name: String,
    size: usize,
}

pub const TEST_INPUT: &str = "";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT), 0);
    }
}
