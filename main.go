// main.go

package main

import (
	"fmt"

	"github.com/janhaj/aoc23/day01"
	"github.com/janhaj/aoc23/day02"
)

func main() {
	fmt.Println("Welcome!")
	fmt.Println("---------")
	fmt.Println("Day 1:")
	fmt.Println("part 1:")
	fmt.Println(day01.Part1())
	fmt.Println("part 2:")
	fmt.Println(day01.Part2())
	fmt.Println("---------")
	fmt.Println("Day 2:")
	fmt.Println("part 1:")
	fmt.Println(day02.Part1())
	fmt.Println("part 2:")
	fmt.Println(day02.Part2())
}
