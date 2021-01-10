// use std::env::var;

pub fn part1(input: String) {
    let values: Vec<_> = input
        .lines()
        .map(|line| line.chars().as_str().repeat(100))
        .collect();
    let cpt = values.iter().enumerate().fold(0, |acc, (i, line)| {
        if line.chars().nth(3 * i).unwrap() == '#' {
            acc + 1
        } else {
            acc
        }
    });
    println!("cpt={}", cpt);
}

pub fn part2(input: String) {}
