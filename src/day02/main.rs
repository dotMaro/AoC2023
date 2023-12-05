use std::fs;

fn main() {
    let input_bytes = fs::read("src/day02/input.txt").unwrap();
    let input = String::from_utf8(input_bytes);
}

fn parse_game(s: &str) -> Game {
    let (game_string, draws) = s.split_once(':').unwrap();
    let id = game_string["Game ".len()..].parse().unwrap();

    let mut cube_sets = vec![];
    for set in draws.split(';') {
        let mut cubes = vec![];
        for cube in set.split(", ") {
            let (count_str, color) = cube.trim_start().split_once(' ').unwrap();
            let count = count_str.parse().unwrap();
            cubes.push(Cube {
                color: color.to_string(),
                count,
            });
        }
        cube_sets.push(CubeSet { cubes });
    }

    Game {
        id,
        draws: cube_sets,
    }
}

#[derive(Debug)]
struct Game {
    id: u8,
    draws: Vec<CubeSet>,
}

#[derive(Debug)]
struct CubeSet {
    cubes: Vec<Cube>,
}

#[derive(Debug)]
struct Cube {
    color: String,
    count: u8,
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn parse_game_examples() {
        let game = parse_game("Game 37: 2 blue, 7 green, 5 red; 5 green, 2 blue; 6 blue, 11 red");
        println!("{:?}", game);
    }
}
