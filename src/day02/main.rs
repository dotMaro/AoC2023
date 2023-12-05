use std::fs;

fn main() {
    let input_bytes = fs::read("src/day02/input.txt").unwrap();
    let input = String::from_utf8(input_bytes);
}

fn filter_games(games: &Vec<Game>, drawn_cubes: &Vec<Cube>) -> u8 {
    games
        .iter()
        .filter(|g| {
            g.draws
                .iter()
                .all(|d| d.is_possible_with_drawn_cubes(drawn_cubes))
        })
        .map(|g| g.id)
        .sum()
}

fn parse_games(s: &str) -> Vec<Game> {
    s.lines().map(|l| parse_game(l)).collect()
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

impl CubeSet {
    fn is_possible_with_drawn_cubes(&self, drawn_cubes: &Vec<Cube>) -> bool {
        self.cubes.iter().all(|c| {
            let matching_cube = drawn_cubes.iter().find(|d| d.color == c.color);
            let res = matching_cube.is_none() || matching_cube.is_some_and(|d| d.count <= c.count);
            println!("{:?} {:?} {}", c, matching_cube, res);
            res
        })
    }
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
    fn part_1_example() {
        let input = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green";
        let games = parse_games(input);
        let sum = filter_games(
            &games,
            &vec![
                Cube {
                    count: 12,
                    color: "red".to_string(),
                },
                Cube {
                    count: 13,
                    color: "green".to_string(),
                },
                Cube {
                    count: 14,
                    color: "blue".to_string(),
                },
            ],
        );
        println!("{:?}", sum);
    }
}
