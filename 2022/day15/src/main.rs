use std::env;
use std::fs;
use std::io;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    let input = fs::read_to_string(&args[1])?;
    let num = get_num(&input, 2000000);
    println!("{num}");
    let num = find_it(&input, 4000000);
    println!("{num}");
    Ok(())
}

fn get_num(input: &str, row: i64) -> usize {
    let b = Board::new(input);
    b.count_on_row(row)
}

fn find_it(input: &str, max: i64) -> i64 {
    let b = Board::new(input);
    for row in 0..=max {
        let col = b.row_has_beacon(row, max);
        if col >= 0 {
            return col * 4000000 + row;
        }
    }
    -1
}

struct Coord {
    x: i64,
    y: i64,
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
    distance: usize,
}

fn distance(c1: &Coord, c2: &Coord) -> usize {
    (c1.x - c2.x).abs() as usize + (c1.y - c2.y).abs() as usize
}

impl Sensor {
    fn new(line: &str) -> Sensor {
        let sensor = Coord::new(
            line.strip_prefix("Sensor at ")
                .unwrap()
                .split(":")
                .nth(0)
                .unwrap(),
        );
        let closest_beacon = Coord::new(
            line.split(": ")
                .nth(1)
                .unwrap()
                .strip_prefix("closest beacon is at ")
                .unwrap(),
        );
        let distance = distance(&sensor, &closest_beacon);

        Sensor {
            sensor: sensor,
            closest_beacon: closest_beacon,
            distance: distance,
        }
    }

    fn excludes(&self, x: i64, y: i64) -> bool {
        if self.closest_beacon.x == x && self.closest_beacon.y == y {
            return false;
        }
        self.distance >= distance(&self.sensor, &Coord { x, y })
    }

    // excludes_until returns the next x coordinate that is not excluded by this sensor on row y.
    fn excludes_until(&self, x: i64, y: i64) -> i64 {
        if !self.excludes(x, y) {
            return x;
        }
        self.sensor.x + self.distance as i64 - (self.sensor.y - y).abs()
    }
}

struct Board {
    sensors: Vec<Sensor>,
    minx: i64,
    _miny: i64,
    maxx: i64,
    _maxy: i64,
}

fn prospective_min(beacon: i64, sensor: i64) -> i64 {
    if beacon < sensor {
        beacon
    } else {
        sensor - (beacon - sensor)
    }
}

fn prospective_max(beacon: i64, sensor: i64) -> i64 {
    if beacon > sensor {
        beacon
    } else {
        sensor + (sensor - beacon)
    }
}

impl Board {
    fn new(input: &str) -> Board {
        let mut sensors = Vec::new();
        for line in input.split("\n") {
            if line == "" {
                continue;
            }
            sensors.push(Sensor::new(line));
        }

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
            _miny: miny,
            maxx: maxx,
            _maxy: maxy,
        }
    }

    fn count_on_row(&self, row: i64) -> usize {
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

    fn row_has_beacon(&self, row: i64, max: i64) -> i64 {
        let mut col = 0;
        loop {
            if col > max {
                return -1;
            }
            let mut nextcol = col;
            for s in &self.sensors {
                let possible_nextcol = s.excludes_until(col, row);
                if possible_nextcol > nextcol {
                    // println!(
                    //     "found {}, {} from beacon {}, {}",
                    //     col, row, s.sensor.x, s.sensor.y
                    // );
                    nextcol = possible_nextcol;
                }
            }
            if col == nextcol {
                return col;
            }
            col = nextcol+1;
        }
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
    use crate::distance;
    use crate::find_it;
    use crate::get_num;
    use crate::Board;
    use crate::Coord;
    use crate::Sensor;
    use crate::TEST_INPUT;

    #[test]
    fn it_works() {
        assert_eq!(get_num(TEST_INPUT, 10), 26);
    }

    #[test]
    fn test_thing() {
        let b = Board::new(TEST_INPUT);
        assert_eq!(b.row_has_beacon(10, 20), -1);
        assert_eq!(b.row_has_beacon(11, 20), 14);
    }

    #[test]
    fn find_it_works() {
        assert_eq!(find_it(TEST_INPUT, 20), 56000011);
    }

    #[test]
    fn sensor_works() {
        let sensor = Coord { x: 8, y: 7 };
        let beacon = Coord { x: 2, y: 10 };
        let distance = distance(&sensor, &beacon);
        assert_eq!(distance, 9);
        let s = Sensor {
            sensor: sensor,
            closest_beacon: beacon,
            distance: distance,
        };
        assert_eq!(s.excludes_until(4, 6), 16);
        assert_eq!(s.excludes_until(7, 5), 15);
        assert_eq!(s.excludes_until(1, 4), 1);
    }
}
