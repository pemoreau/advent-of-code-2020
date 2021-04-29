package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func readFile() []int {
	file, _ := os.Open("../../inputs/day09.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []int{}
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		res = append(res, v)
	}
	return res
}

func check25(list []int, start int, end int, value int) bool {
	for i := start; i < end; i++ {
		for j := i + 1; j < end; j++ {
			if list[i]+list[j] == value {
				return true
			}
		}
	}
	return false
}

func searchSum(list []int, value int) int {
	for i := 0; i < len(list); i++ {
		sum := 0
		min := math.MaxInt64
		max := 0
		for j := i; j < len(list); j++ {
			v := list[j]
			sum += v
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
			if sum == value {
				return min + max
			} else if sum > value {
				break
			}
		}
	}
	return 0
}

func part1(list []int) int {
	for i := 0; i < len(list); i++ {
		if !check25(list, i, i+25, list[i+25]) {
			return list[i+25]
		}
	}
	return 0
}

func main() {
	list := readFile()
	start := time.Now()
	num := part1(list)
	fmt.Println("part1: ", num)
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", searchSum(list, num))
	fmt.Println(time.Since(start))
}
