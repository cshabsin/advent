use std::env;
use std::fs;
use std::io;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num(&input, 2000000);
    println!("{num}");
    Ok(())
}

fn get_num(input: &str, row: i32) -> usize {
    let mut sensors = Vec::new();
    for line in input.split("\n") {
        if line == "" {
            continue;
        }
        sensors.push(Sensor::new(line));
    }
    let b = Board::new(sensors);
    b.count_on_row(row)
}

struct Coord {
    x: i32,
    y: i32,
}

impl Coord {
    fn new(line: &str) -> Coord {
        let mut split = line.split(", ").into_iter();
        let (x, y) = (split.next().unwrap(), split.next().unwrap());
        let (x, y) = (
            x.strip_prefix("x=").unwrap().parse().unwrap(),
            y.strip_prefix("y=").unwrap().parse().unwrap(),
        );
        Coord { x: x, y: y }
    }
}

struct Sensor {
    sensor: Coord,
    closest_beacon: Coord,
}

fn distance(c1: &Coord, c2: &Coord) -> usize {
    (c1.x - c2.x).abs() as usize + (c1.y - c2.y).abs() as usize
}

impl Sensor {
    fn new(line: &str) -> Sensor {
        Sensor {
            sensor: Coord::new(
                line.strip_prefix("Sensor at ")
                    .unwrap()
                    .split(":")
                    .nth(0)
                    .unwrap(),
            ),
            closest_beacon: Coord::new(
                line.split(": ")
                    .nth(1)
                    .unwrap()
                    .strip_prefix("closest beacon is at ")
                    .unwrap(),
            ),
        }
    }

    fn excludes(&self, x: i32, y: i32) -> bool {
        if self.closest_beacon.x == x && self.closest_beacon.y == y {
            return false;
        }
        distance(&self.sensor, &self.closest_beacon) >= distance(&self.sensor, &Coord { x, y })
    }
}

struct Board {
    sensors: Vec<Sensor>,
    minx: i32,
    miny: i32,
    maxx: i32,
    maxy: i32,
}

fn prospective_min(beacon: i32, sensor: i32) -> i32 {
    if beacon < sensor {
        beacon
    } else {
        sensor - (beacon - sensor)
    }
}

fn prospective_max(beacon: i32, sensor: i32) -> i32 {
    if beacon > sensor {
        beacon
    } else {
        sensor + (sensor - beacon)
    }
}

impl Board {
    fn new(sensors: Vec<Sensor>) -> Board {
        let (mut minx, mut miny, mut maxx, mut maxy) = (100000000, 100000000, 0, 0);
        for s in &sensors {
            let prospective_minx = prospective_min(s.closest_beacon.x, s.sensor.x);
            let prospective_maxx = prospective_max(s.closest_beacon.x, s.sensor.x);
            let prospective_miny = prospective_min(s.closest_beacon.y, s.sensor.y);
            let prospective_maxy = prospective_max(s.closest_beacon.y, s.sensor.y);
            if prospective_minx < minx {
                minx = prospective_minx;
            }
            if prospective_miny < miny {
                miny = prospective_miny;
            }
            if prospective_maxx > maxx {
                maxx = prospective_maxx;
            }
            if prospective_maxy > maxy {
                maxy = prospective_maxy;
            }
        }
        Board {
            sensors: sensors,
            minx: minx,
            miny: miny,
            maxx: maxx,
            maxy: maxy,
        }
    }

    fn count_on_row(&self, row: i32) -> usize {
        let mut rc = 0;
        for col in self.minx..=self.maxx {
            for s in &self.sensors {
                if s.excludes(col, row) {
                    rc += 1;
                    break;
                }
            }
        }
        rc
    }
}

pub const TEST_INPUT: &str = "Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3";

#[cfg(test)]
mod tests {
    use crate::get_num;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT, 10), 26);
    }
}
