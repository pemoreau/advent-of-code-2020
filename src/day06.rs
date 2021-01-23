pub fn part1(input: String) {
    let mut letters = [false; 27];
    let mut scores: Vec<usize> = Vec::new();

    input.lines().for_each(|line| {
        if line.is_empty() {
            scores.push(letters.iter().filter(|&letter| *letter).count());
            letters = [false; 27];
        } else {
            line.chars()
                .for_each(|c| letters[1 + c as usize - 'a' as usize] = true);
        }
    });
    scores.push(letters.iter().filter(|&letter| *letter).count()); // for last entry
    println!("res= {}", scores.iter().sum::<usize>());
}

pub fn part2(input: String) {}
