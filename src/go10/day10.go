package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func readFile() []int {
	file, _ := os.Open("../../inputs/day10.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []int{}
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		res = append(res, v)
	}
	return res
}

func part1(list []int) int {
	sort.Ints(list)
	diff := map[int]int{1: 0, 2: 0, 3: 1}
	diff[list[0]]++
	for i, v := range list {
		if i > 0 {
			diff[v-list[i-1]] += 1
		}
	}
	fmt.Println(diff)
	return diff[1] * diff[3]
}

func part2(list []int) int {
	sort.Ints(list)
	list = append(list, list[len(list)-1]+3)
	list = append([]int{0}, list...)
	res := 1
	nb1 := 0
	// incr := []int{}
	for i, v := range list {
		if i > 0 {
			diff := v - list[i-1]
			if diff == 1 {
				nb1++
			}
			if diff == 3 {
				switch nb1 {
				// could have use tribonacci numbers
				case 4:
					res *= 7
				case 3:
					res *= 4
				case 2:
					res *= 2
				}
				nb1 = 0
			}
			// incr = append(incr, diff)
		}
	}
	// fmt.Println(list)
	// fmt.Println(incr)
	return res
}

func main() {
	list := readFile()
	start := time.Now()
	num := part1(list)
	fmt.Println("part1: ", num)
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", part2(list))
	fmt.Println(time.Since(start))
}
