package main

import (
	"advent/day1"
	"advent/day2"
	"advent/day3"
	"advent/day4"
	"fmt"
)

const (
	Day1Input  = "day1/input.txt"
	Day1Sample = "day1/sample.txt"
	Day2Input  = "day2/input.txt"
	Day2Sample = "day2/sample.txt"
	Day3Input  = "day3/input.txt"
	Day3Sample = "day3/sample.txt"
	Day4Input  = "day4/input.txt"
	Day4Sample = "day4/sample.txt"
)

var days = []struct {
	input string
	part1 func(string) int64
	part2 func(string) int64
}{
	{
		input: Day1Sample,
		part1: day1.Part1,
		part2: day1.Part2,
	},
	{
		input: Day2Sample,
		part1: day2.Part1,
		part2: day2.Part2,
	},
	{
		input: Day3Sample,
		part1: day3.Part1,
		part2: day3.Part2,
	},
	{
		input: Day4Input,
		part1: day4.Part1,
		part2: day4.Part2,
	},
}

func main() {
	for idx, day := range days {
		fmt.Printf("Day%d:\n", idx+1)
		fmt.Println("  Part 1:")
		fmt.Printf("  %d\n", day.part1(day.input))
		fmt.Println("  Part 2:")
		fmt.Printf("  %d\n", day.part2(day.input))
		fmt.Println()
	}
}
