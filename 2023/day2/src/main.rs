use std::fs::read_to_string;

struct Game {
    id: u32,
    max_red: u32,
    max_green: u32,
    max_blue: u32,
    possible: bool,
    power: u32,
}

fn main() {
    let lines: Vec<String> = read_lines("input.txt");
    p1_p2(lines);
}

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename)
        .unwrap()
        .lines()
        .map(String::from)
        .collect()
}

fn p1_p2(lines: Vec<String>) {
    let red_count: u32 = 12;
    let green_count: u32 = 13;
    let blue_count: u32 = 14;

    let mut id_sum: u32 = 0;
    let mut games: Vec<Game> = Vec::new();
    let mut all_power: u32 = 0;

    for row in lines {
        let id: u32 = row
            .split(":")
            .next()
            .unwrap()
            .split(" ")
            .last()
            .unwrap()
            .parse()
            .unwrap();

        let mut game = Game {
            id,
            max_red: 0,
            max_green: 0,
            max_blue: 0,
            possible: false,
            power: 0,
        };

        let sets: Vec<&str> = row.split(":").last().unwrap().split(";").collect();
        for set in sets {
            let handful: Vec<&str> = set.split(",").collect();
            for group in handful {
                let group_split: Vec<&str> = group.trim().split(" ").collect();
                let count: u32 = group_split[0].parse().unwrap();
                let color: &str = group_split[1];
                match color {
                    "red" => {
                        if game.max_red < count {
                            game.max_red = count;
                        }
                    }
                    "green" => {
                        if game.max_green < count {
                            game.max_green = count;
                        }
                    }
                    "blue" => {
                        if game.max_blue < count {
                            game.max_blue = count;
                        }
                    }
                    _ => println!("Unknown color"),
                }
            }
        }
        if red_count >= game.max_red && green_count >= game.max_green && blue_count >= game.max_blue
        {
            game.possible = true;
        }
        game.power = game.max_red * game.max_green * game.max_blue;
        games.push(game)
    }

    for game in &mut games {
        if game.possible {
            id_sum += game.id;
        }
        all_power += game.power;
    }

    println!("Part 1: {}", id_sum);
    println!("Part 2: {}", all_power)
}
