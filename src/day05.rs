fn to_bin(input: &str, lower: char) -> usize {
    input
        .chars()
        .fold(0, |acc, c| if c == lower { 2 * acc } else { 2 * acc + 1 })
}

pub fn part1(input: String) {
    // F:0 B:1
    // L:0 R:1
    let res = input
        .lines()
        .map(|line| to_bin(&line[0..7], 'F') * 8 + to_bin(&line[7..10], 'L'))
        .max()
        .unwrap();
    println!("res= {}", res);
}

pub fn part2(input: String) {
    let mut seats = [false; 912];
    input.lines().for_each(|line| {
        let seat = to_bin(&line[0..7], 'F') * 8 + to_bin(&line[7..10], 'L');
        seats[seat] = true;
    });

    let mut previous = false;
    seats.iter().enumerate().for_each(|(index, element)| {
        if !element && previous {
            println!("empty seat: {}", index);
        }
        previous = *element;
    });
}
