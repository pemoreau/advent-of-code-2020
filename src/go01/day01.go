package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func readFile() []int {
	file, _ := os.Open("../../inputs/day01.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []int{}
	for scanner.Scan() {
		l := scanner.Text()
		// fmt.Println(l)
		value, _ := strconv.Atoi(l)
		res = append(res, value)
	}
	return res
}

func part1(list []int) int {
	for i1, v1 := range list {
		for i2, v2 := range list {
			if i1 != i2 && v1+v2 == 2020 {
				return v1 * v2
			}
		}
	}
	return 0
}

func part2(list []int) int {
	for i1, v1 := range list {
		for i2, v2 := range list {
			for i3, v3 := range list {
				if i1 != i2 && i2 != i3 && i1 != i3 && v1+v2+v3 == 2020 {
					return v1 * v2 * v3
				}
			}
		}
	}
	return 0
}

func main() {
	file, _ := os.Open("../../inputs/day01.txt")
	defer file.Close()
	list := readFile()
	start := time.Now()
	fmt.Println("part1: ", part1(list))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", part2(list))
	fmt.Println(time.Since(start))
}
