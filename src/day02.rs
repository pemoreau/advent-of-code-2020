// use std::env::var;

#[derive(Debug)]

pub struct Policy {
    min: u32,
    max: u32,
    letter: char,
}

fn parse_line1(s: &str) -> (Policy, &str) {
    let re = regex::Regex::new(r"^(\d{1,2})-(\d{1,2}) ([a-z]{1}): ([a-z]+)$").unwrap();
    let cap = re.captures(s).unwrap();
    let password: &str = cap.get(4).unwrap().as_str(); // instead of &cap[4]

    (
        Policy {
            min: cap[1].parse().unwrap(),
            max: cap[2].parse().unwrap(),
            letter: cap[3].chars().next().unwrap(),
        },
        password,
    )
}

fn parse_line2(s: &str) -> (Policy, &str) {
    peg::parser! {
      grammar parser() for str {
        rule number() -> u32
          = n:$(['0'..='9']+) { n.parse().unwrap() }

        rule byte() -> char
          = letter:$(['a'..='z']) { letter.chars().next().unwrap() }

        rule password() -> &'input str
          = letters:$([_]*) { letters }

        pub(crate) rule line() -> (Policy, &'input str)
          = min:number() "-" max:number() " " letter:byte() ": " password:password() {
              (Policy { min,max,letter }, password)
          }
      }
    }

    parser::line(s).unwrap()
}

fn check(policy: &Policy, password: &str) -> bool {
    let c = password.chars().filter(|c| *c == policy.letter).count() as u32;
    c >= policy.min && c <= policy.max
}

pub fn part1(input: String) {
    let values: Vec<_> = input.lines().map(|line| parse_line1(line)).collect();
    let number = values
        .iter()
        .filter(|(policy, password)| check(policy, password))
        .count();
    println!("{}", number)
}
