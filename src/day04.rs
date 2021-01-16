extern crate nom;

use std::collections::HashSet;

use nom::{
    branch::alt,
    bytes::complete::tag,
    character::complete::{alpha1, alphanumeric1, line_ending},
    combinator::{map, opt, recognize},
    multi::many1,
    sequence::{pair, terminated},
    IResult,
};

#[derive(Debug, PartialEq)]
pub struct Entry {
    key: String,
    value: String,
}
#[derive(Debug, PartialEq)]
pub struct Passport {
    passport: Vec<Entry>,
}

fn parse_key(input: &str) -> IResult<&str, &str> {
    terminated(alpha1, tag(":"))(input)
}

fn parse_value(input: &str) -> IResult<&str, &str> {
    recognize(pair(opt(tag("#")), alphanumeric1))(input)
}

fn parse_entry(input: &str) -> IResult<&str, Entry> {
    map(pair(parse_key, parse_value), |(key, value)| Entry {
        key: key.to_string(),
        value: value.to_string(),
    })(input)
}

fn parse_entries(input: &str) -> IResult<&str, Passport> {
    map(
        many1(terminated(parse_entry, alt((tag(" "), line_ending)))),
        |vec| Passport { passport: vec },
    )(input)
}

fn parse_passports(input: &str) -> IResult<&str, Vec<Passport>> {
    map(many1(terminated(parse_entries, line_ending)), |vec| vec)(input)
}

fn valid_passport(p: &Passport) -> bool {
    let mandatory_keys: HashSet<_> = ["eyr", "iyr", "byr", "ecl", "pid", "hcl", "hgt"]
        .iter()
        .cloned()
        .collect();
    let keys: HashSet<_> = p
        .passport
        .iter()
        .map(|Entry { key, value: _ }| &key[..])
        .collect();
    mandatory_keys.is_subset(&keys)
}

pub fn part1(input: String) {
    assert_eq!(parse_key("abc:"), Ok(("", "abc")));
    assert_eq!(parse_value("cfa07d"), Ok(("", "cfa07d")));
    assert_eq!(parse_value("#cfa07d"), Ok(("", "#cfa07d")));
    assert_eq!(
        parse_entry("abc:#cfa07d"),
        Ok((
            "",
            Entry {
                key: "abc".to_string(),
                value: "#cfa07d".to_string()
            }
        ))
    );

    let (_, passports) = parse_passports(&input[..]).unwrap();
    println!(
        "{}",
        passports.iter().filter(|p| valid_passport(&p)).count()
    );
}

pub fn part2(_input: String) {}
