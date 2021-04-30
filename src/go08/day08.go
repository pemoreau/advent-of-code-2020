package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Instr struct {
	op  string
	arg int
}

func parseLine(line string) Instr {
	parts := strings.Split(line, " ")
	op := strings.TrimSpace(parts[0])
	arg, _ := strconv.Atoi(parts[1])
	return Instr{op, arg}
}

func readFile() []Instr {
	file, _ := os.Open("../../inputs/day08.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []Instr{}
	for scanner.Scan() {
		res = append(res, parseLine(scanner.Text()))
	}
	return res
}

type Env struct {
	prg []Instr
	acc int
	pc  int
}

func step(env *Env) {
	if env.pc < len(env.prg) {
		inst := env.prg[env.pc]
		if inst.op == "acc" {
			// fmt.Printf("ACC: pc=%d acc=%d\n", env.pc, env.acc)
			env.acc += inst.arg
			env.pc++
		} else if inst.op == "jmp" {
			// fmt.Printf("JMP: pc=%d arg=%d\n", env.pc, inst.arg)
			env.pc += inst.arg
		} else if inst.op == "nop" {
			// fmt.Printf("NOP: pc=%d\n", env.pc)
			env.pc++
		}
	}
}

type IntSet map[int]struct{}

func (s IntSet) add(v int) {
	s[v] = struct{}{}
}

func (s IntSet) includes(v int) bool {
	_, ok := s[v]
	return ok
}

func terminates(env Env) (bool, int) {
	visited := make(IntSet)
	for !visited.includes(env.pc) {
		if env.pc >= len(env.prg) {
			return true, env.acc
		}
		visited.add(env.pc)
		step(&env)
	}
	return false, env.acc
}

func part1(prg []Instr) int {
	_, acc := terminates(Env{prg: prg})
	return acc
}

func part2(prg []Instr) int {
	env := Env{prg: prg}

	for line, inst := range prg {
		if inst.op == "nop" {
			env.prg[line].op = "jmp"
			if t, acc := terminates(env); t {
				return acc
			}
			env.prg[line].op = "nop"
		}
	}
	for line, inst := range prg {
		if inst.op == "jmp" {
			env.prg[line].op = "nop"
			if t, acc := terminates(env); t {
				return acc
			}
			env.prg[line].op = "jmp"
		}
	}

	return 0
}

func main() {
	list := readFile()
	start := time.Now()
	fmt.Println("part1: ", part1(list))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", part2(list))
	fmt.Println(time.Since(start))
}
