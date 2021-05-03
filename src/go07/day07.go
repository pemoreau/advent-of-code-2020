package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Color struct {
	name    string
	contain []ColorWithMultiplicity
}

type ColorWithMultiplicity struct {
	name         string
	multiplicity int
}

func parseLine(line string) Color {
	parts := strings.Split(line, "bags")
	color := Color{
		name:    strings.TrimSpace(parts[0]),
		contain: []ColorWithMultiplicity{},
	}
	parts = strings.Split(line, "contain")
	parts = strings.Split(strings.TrimSpace(parts[1]), " ")
	for i := 0; i < len(parts); i = i + 4 {
		multiplicity, err := strconv.Atoi(parts[i])
		if err == nil {
			colorm := ColorWithMultiplicity{
				multiplicity: multiplicity,
				name:         parts[i+1] + " " + parts[i+2],
			}
			color.contain = append(color.contain, colorm)
		}
	}
	return color
}

type Store = map[string][]ColorWithMultiplicity

func readFile() Store {
	file, _ := os.Open("../../inputs/day07.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := make(Store)
	for scanner.Scan() {
		color := parseLine(scanner.Text())
		res[color.name] = color.contain
	}

	return res
}

// returns true if name or a descendent contains colorName
func containsColor(s Store, name string, colorName string) bool {
	// fmt.Printf("<%s> contains? <%s>\n", name, colorName)
	colors := s[name]
	for _, c := range colors {
		if c.name == colorName || containsColor(s, c.name, colorName) {
			return true
		}
	}
	return false
}
func countBags(s Store, name string) int {
	// fmt.Printf("<%s> contains? <%s>\n", name, colorName)
	colors := s[name]
	res := 1
	for _, c := range colors {
		res += c.multiplicity * countBags(s, c.name)
	}
	return res
}

func part1(s Store) (res int) {
	for key := range s {
		b := containsColor(s, key, "shiny gold")
		if b {
			res++
		}
	}
	return
}

func part2(s Store) int {
	return countBags(s, "shiny gold") - 1
}

func main() {
	colors := readFile()

	start := time.Now()
	fmt.Println("part1: ", part1(colors))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", part2(colors))
	fmt.Println(time.Since(start))
}
