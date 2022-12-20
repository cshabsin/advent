use std::cell::RefCell;
use std::collections::HashMap;
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
    let root = Rc::new(RefCell::new(Directory::new("/", None)));
    parse_input(input, Rc::clone(&root));
    let mut dirs = vec![Rc::clone(&root)];
    let mut total = 0;
    loop {
        match dirs.pop() {
            None => {
                return total;
            }
            Some(dir) => {
                if dir.borrow().size() <= 100000 {
                    total += dir.borrow().size();
                }
                dirs.append(&mut dir.borrow().children());
            }
        }
    }
}

fn parse_input(input: &str, root: Rc<RefCell<Directory>>) {
    let mut cwd = Rc::clone(&root);
    for line in input.split("\n") {
        if line == "" {
            continue;
        }
        match InputLine::new(line) {
            InputLine::Root() => (),
            InputLine::Ls() => (),
            InputLine::Dir(name) => add_dir(Rc::clone(&cwd), &name),
            InputLine::File(name, size) => cwd.borrow_mut().add_file(&name, size),
            InputLine::ChdirUp() => {
                let new_cwd = match cwd.borrow().parent() {
                    Some(new_cwd) => new_cwd,
                    None => panic!("no parent for {}", cwd.borrow().name),
                };
                cwd = Rc::clone(&new_cwd);
            } // crash if cd .. in root
            InputLine::Chdir(dir) => {
                let new_cwd = cwd.borrow().child(&dir);
                cwd = Rc::clone(&new_cwd);
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
        } else if line.starts_with("$ cd ") {
            InputLine::Chdir(line.split_at(5).1.to_string())
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
    subdirs: HashMap<String, Rc<RefCell<Directory>>>,
    files: Vec<File>,
    parent_link: Option<Weak<RefCell<Directory>>>,
}

// impl Rc<Directory> {
//     fn add_dir(self, name: &str) {
//         self.subdirs.push(Directory::new(name, Some<self>))
//     }
// }
// fn add_dir(&mut self, name: &str) {
//     if !self.subdirs.contains_key(name) {
//         self.subdirs.insert(name, Rc::new(RefCell::new(Directory{name: name.to_string(), files: Vec::new(), subdirs: HashMap::new(), parent_link: Some())));
//     }
// }

// TODO(cshabsin): Figure out if this can be done as a method on Directory.
// Maybe using Rc::new_cyclic to hold a weak pointer to self to use in this?
// We can't just construct a new Weak<RefCell<Directory>> of self, can we? That wouldn't have any connection to the existing Rc.
// Or is Rc just That Magic?
fn add_dir(dir: Rc<RefCell<Directory>>, subdir: &str) {
    // println!("add_dir({}, {})", dir.borrow().name, subdir);
    if dir.borrow().subdirs.contains_key(subdir) {
        // println!("already there");
        return; // no need to do anything
    }
    dir.borrow_mut().subdirs.insert(
        subdir.to_string(),
        Rc::new(RefCell::new(Directory {
            name: subdir.to_string(),
            files: Vec::new(),
            subdirs: HashMap::new(),
            parent_link: Some(Rc::downgrade(&dir)),
        })),
    );
}

impl Directory {
    fn new(name: &str, parent: Option<&Rc<RefCell<Directory>>>) -> Directory {
        Directory {
            name: name.to_string(),
            subdirs: HashMap::new(),
            files: Vec::new(),
            parent_link: match parent {
                Some(parent) => Some(Rc::downgrade(parent)),
                None => None,
            },
        }
    }

    fn add_file(&mut self, name: &str, size: usize) {
        // println!("add_file({}, {})", self.name, name);
        self.files.push(File {
            name: name.to_string(),
            size,
        })
    }

    fn size(&self) -> usize {
        let mut size = 0;
        for dir in &self.subdirs {
            size += dir.1.borrow().size();
        }
        for f in &self.files {
            size += f.size;
        }
        // println!("size({}): {}", self.name, size);
        size
    }

    fn child(&self, name: &str) -> Rc<RefCell<Directory>> {
        // println!("child({}, {})", self.name, name);
        match self.subdirs.get(name) {
            None => panic!("no subdir {name} found in {}", self.name),
            Some(foo) => Rc::clone(&foo),
        }
    }

    fn parent(&self) -> Option<Rc<RefCell<Directory>>> {
        // println!("parent({})", self.name);
        match &self.parent_link {
            None => panic!("attempt to get parent of root"),
            Some(parent_link) => parent_link.upgrade(),
        }
    }

    fn children(&self) -> Vec<Rc<RefCell<Directory>>> {
        let mut rc = Vec::new();
        for c in &self.subdirs {
            rc.push(Rc::clone(&c.1));
        }
        rc
    }
}

struct File {
    name: String,
    size: usize,
}

pub const TEST_INPUT: &str = "$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
";

#[cfg(test)]
mod tests {
    use std::cell::RefCell;
    use std::rc::Rc;

    use crate::add_dir;
    use crate::get_num;
    use crate::Directory;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT), 95437);
    }

    #[test]
    fn dir_memory() {
        let root = Rc::new(RefCell::new(Directory::new("/", None)));
        root.borrow_mut().add_file("hi", 3);
        add_dir(Rc::clone(&root), "subdir");
        let sd = root.borrow().child("subdir");
        match sd.borrow().parent() {
            Some(_) => (),
            None => panic!("no parent"),
        };
    }
}
