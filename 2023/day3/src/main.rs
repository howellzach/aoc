use std::collections::HashMap;
use std::fs::read_to_string;

const SYMBOLS: [char; 29] = [
    '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '\\',
    '|', ';', ':', '"', '\'', '?', '>', '<', ',', '/',
];

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
    let mut engine: Vec<i32> = Vec::new();
    let mut buf = [0; 4];
    let mut num_loc: HashMap<String, i32> = HashMap::new();
    let mut sym_loc: HashMap<i32, Vec<i32>> = HashMap::new();
    let mut gear_loc: HashMap<i32, Vec<i32>> = HashMap::new();
    for (row_idx, row) in lines.iter().enumerate() {
        // Add a '.' to each line to capture numbers on the edge
        let row = &(row.to_owned() + &(".".to_string()));
        let mut num_str: String = "".to_string();
        let mut sym_loc_vec: Vec<i32> = Vec::new();
        let mut gear_loc_vec: Vec<i32> = Vec::new();

        for (char_idx, c) in row.chars().enumerate() {
            if c.is_digit(10) {
                num_str.push_str(c.encode_utf8(&mut buf));
            }
            if SYMBOLS.contains(&c) || c == '.' {
                if !num_str.is_empty() {
                    let num: i32 = num_str.parse().unwrap();
                    let end: i32 = char_idx as i32;
                    let start: i32 = end - num.to_string().len() as i32;
                    let loc: String = format!("{}={}={}", row_idx, start, end);
                    num_loc.insert(loc, num);
                    num_str = "".to_string();
                }
                if c != '.' {
                    sym_loc_vec.push(char_idx as i32)
                }
                if c == '*' {
                    gear_loc_vec.push(char_idx as i32)
                }
            }
        }
        sym_loc.insert(row_idx as i32, sym_loc_vec);
        if gear_loc_vec.len() > 0 {
            gear_loc.insert(row_idx as i32, gear_loc_vec);
        }
    }

    let mut num_loc_hashmap: HashMap<i32, Vec<HashMap<i32, Vec<i32>>>> = HashMap::new();
    for (n_loc, num) in num_loc {
        let mut split_n_loc = n_loc.split("=");
        let row_idx: i32 = split_n_loc.next().unwrap().parse().unwrap();
        let start_idx: i32 = split_n_loc.next().unwrap().parse().unwrap();
        let end_idx: i32 = split_n_loc.next().unwrap().parse().unwrap();
        let row_matches: Vec<i32> = (start_idx - 1..end_idx + 1).collect();

        // Create number hashmap for part 2
        let mut row_map: HashMap<i32, Vec<i32>> = HashMap::new();
        row_map.insert(num, row_matches.clone());
        match num_loc_hashmap.get_mut(&row_idx) {
            Some(row_m) => row_m.push(row_map),
            None => {
                num_loc_hashmap.insert(row_idx, vec![row_map]);
            }
        }

        for i in row_matches {
            // Check every row for adjacent symbols
            if sym_loc.get(&row_idx).unwrap().contains(&i) {
                engine.push(num);
                break;
            }
            // Do not check below last row for adjacent symbols
            if !(row_idx == lines.len() as i32 - 1) {
                if sym_loc.get(&(row_idx + 1)).unwrap().contains(&i) {
                    engine.push(num);
                    break;
                }
            }
            // Do not check above first row for adjacent symbols
            if row_idx != 0 {
                if sym_loc.get(&(row_idx - 1)).unwrap().contains(&i) {
                    engine.push(num);
                    break;
                }
            }
        }
    }
    let res1: i32 = engine.iter().sum();
    println!("Part 1: {}", res1);

    let mut gear_sums: Vec<i32> = Vec::new();
    for (gear_row, gears_in_row) in gear_loc {
        let empty_map = vec![];
        let above_row = num_loc_hashmap.get(&(gear_row - 1)).unwrap_or(&empty_map);
        let same_row = num_loc_hashmap.get(&gear_row).unwrap_or(&empty_map);
        let below_row = num_loc_hashmap.get(&(gear_row + 1)).unwrap_or(&empty_map);

        for gear in gears_in_row {
            let mut gear_mult: Vec<i32> = vec![];
            for a_row in above_row {
                for (n, n_loc) in a_row {
                    if n_loc.contains(&gear) {
                        gear_mult.push(*n);
                    }
                }
            }
            for s_row in same_row {
                for (n, n_loc) in s_row {
                    if n_loc.contains(&gear) {
                        gear_mult.push(*n);
                    }
                }
            }
            for b_row in below_row {
                for (n, n_loc) in b_row {
                    if n_loc.contains(&gear) {
                        gear_mult.push(*n);
                    }
                }
            }
            if gear_mult.len() == 2 {
                gear_sums.push(gear_mult[0] * gear_mult[1]);
            }
        }
    }
    let res2: i32 = gear_sums.iter().sum();
    println!("Part 2: {}", res2);
}
