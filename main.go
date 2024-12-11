package main

import (
	"advent/day1"
	"advent/day10"
	"advent/day11"
	"advent/day2"
	"advent/day3"
	"advent/day4"
	"advent/day5"
	"advent/day6"
	"advent/day7"
	"advent/day8"
	"advent/day9"
	"fmt"
)

const (
	Day1Sample  = "day1/sample.txt"
	Day2Sample  = "day2/sample.txt"
	Day3Sample  = "day3/sample.txt"
	Day4Sample  = "day4/sample.txt"
	Day5Sample  = "day5/sample.txt"
	Day6Sample  = "day6/sample.txt"
	Day7Sample  = "day7/sample.txt"
	Day8Sample  = "day8/sample.txt"
	Day9Sample  = "day9/sample.txt"
	Day10Sample = "day10/sample.txt"
	Day11Sample = "day11/sample.txt"
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
		input: Day4Sample,
		part1: day4.Part1,
		part2: day4.Part2,
	},
	{
		input: Day5Sample,
		part1: day5.Part1,
		part2: day5.Part2,
	},
	{
		input: Day6Sample,
		part1: day6.Part1,
		part2: day6.Part2,
	},
	{
		input: Day7Sample,
		part1: day7.Part1,
		part2: day7.Part2,
	},
	{
		input: Day8Sample,
		part1: day8.Part1,
		part2: day8.Part2,
	},
	{
		input: Day9Sample,
		part1: day9.Part1,
		part2: day9.Part2,
	},
	{
		input: Day10Sample,
		part1: day10.Part1,
		part2: day10.Part2,
	},
	{
		input: Day11Sample,
		part1: day11.Part1,
		part2: day11.Part2,
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
