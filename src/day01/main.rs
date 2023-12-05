use std::fs;

fn main() {
    let input_vec = fs::read("src/day01/input.txt").unwrap();
    let input = String::from_utf8(input_vec).unwrap();
    println!("Part 1. {}", sum_of_calibration_values(&input));
    println!(
        "Part 2. {}",
        sum_of_calibration_values_including_written_out_digits(&input)
    );
}

fn sum_of_calibration_values(s: &str) -> u32 {
    s.lines().map(|l| get_calibration_value(l)).sum()
}

fn sum_of_calibration_values_including_written_out_digits(s: &str) -> u32 {
    s.lines()
        .map(|l| get_calibration_value_including_written_out_digits(l))
        .sum()
}

fn get_calibration_value(s: &str) -> u32 {
    let first_digit = s
        .chars()
        .find(|c| c.is_ascii_digit())
        .unwrap()
        .to_digit(10)
        .unwrap();
    let last_digit = s
        .chars()
        .rev()
        .find(|c| c.is_ascii_digit())
        .unwrap()
        .to_digit(10)
        .unwrap();
    first_digit * 10 + last_digit
}

fn get_calibration_value_including_written_out_digits(s: &str) -> u32 {
    let mut first_digit = 0;
    let mut last_digit = 0;
    for i in 0..s.len() {
        match starts_with_digit(&s[i..]) {
            Some(n) => {
                first_digit = n;
                break;
            }
            None => (),
        }
    }
    for i in (0..s.len()).rev() {
        match starts_with_digit(&s[i..]) {
            Some(n) => {
                last_digit = n;
                break;
            }
            None => (),
        }
    }
    first_digit * 10 + last_digit
}

fn starts_with_digit(s: &str) -> Option<u32> {
    let c = s.chars().next().unwrap();
    if c.is_ascii_digit() {
        return Some(c.to_digit(10).unwrap());
    }

    starts_with_written_out_digit(s)
}

fn starts_with_written_out_digit(s: &str) -> Option<u32> {
    for (digit_str, digit) in vec![
        ("one", 1),
        ("two", 2),
        ("three", 3),
        ("four", 4),
        ("five", 5),
        ("six", 6),
        ("seven", 7),
        ("eight", 8),
        ("nine", 9),
    ] {
        if s.starts_with(digit_str) {
            return Some(digit);
        }
    }
    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn get_calibration_value_examples() {
        assert_eq!(get_calibration_value("asfkjh3i2a"), 32);
        assert_eq!(get_calibration_value("jhnsapfaskö9kns7af"), 97);
    }

    #[test]
    fn part1_example() {
        let input = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";
        assert_eq!(sum_of_calibration_values(input), 142);
    }

    #[test]
    fn part2_example() {
        let input = "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen";
        assert_eq!(
            sum_of_calibration_values_including_written_out_digits(input),
            281
        );
    }
}
