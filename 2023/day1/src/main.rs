use std::collections::HashMap;
use std::fs::read_to_string;

fn main() {
    let lines: Vec<String> = read_lines("input.txt");
    part1(lines.clone());
    part2(lines.clone());
}

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename)
        .unwrap()
        .lines()
        .map(String::from)
        .collect()
}

fn part1(lines: Vec<String>) {
    let mut nums: Vec<u32> = Vec::new();
    for line in lines {
        let mut line_nums: Vec<String> = Vec::new();
        for c in line.chars() {
            if c.is_numeric() {
                line_nums.push(c.to_string());
            }
        }
        let str_num = line_nums.first().unwrap().to_owned() + line_nums.last().unwrap();
        nums.push(str_num.parse().unwrap())
    }
    let p1: u32 = nums.iter().sum();
    println!("Part 1: {}", p1);
}

fn part2(mut lines: Vec<String>) {
    let num_conv = HashMap::from([
        ("one", "1"),
        ("two", "2"),
        ("three", "3"),
        ("four", "4"),
        ("five", "5"),
        ("six", "6"),
        ("seven", "7"),
        ("eight", "8"),
        ("nine", "9"),
    ]);
    let mut nums: Vec<u32> = Vec::new();
    for line in &mut lines {
        for (x, y) in &num_conv {
            let s = format!(
                "{}{}{}",
                x.chars().next().unwrap(),
                y,
                x.chars().last().unwrap()
            );
            *line = line.replace(x, &s);
        }
        let mut line_nums: Vec<String> = Vec::new();
        for c in line.chars() {
            if c.is_numeric() {
                line_nums.push(c.to_string());
            }
        }
        let str_num = line_nums.first().unwrap().to_owned() + line_nums.last().unwrap();
        nums.push(str_num.parse().unwrap())
    }
    let p2: u32 = nums.iter().sum();
    println!("Part 2: {}", p2);
}
