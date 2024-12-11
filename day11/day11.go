package day11

import (
	. "advent/util"
	"strconv"
	"strings"
)

func Part1(filename string) int64 {
	var count int64
	memo := make(map[string]map[int]int64)
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		for _, digits := range strings.Fields(scanner.Text()) {
			count += countStone(memo, digits, 25)
		}
	}
	return count
}

func Part2(filename string) int64 {
	var count int64
	memo := make(map[string]map[int]int64)
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		for _, digits := range strings.Fields(scanner.Text()) {
			count += countStone(memo, digits, 75)
		}
	}
	return count
}

func countStone(memo map[string]map[int]int64, digits string, iteration int) int64 {
	if _, ok := memo[digits]; !ok {
		memo[digits] = make(map[int]int64)
	}
	if count, ok := memo[digits][iteration]; ok {
		return count
	} else if iteration == 0 {
		memo[digits][iteration] = 1
		return 1
	} else {
		if digits == "0" {
			count := countStone(memo, "1", iteration-1)
			memo[digits][iteration] = count
			return count
		} else if len(digits)%2 == 0 {
			mid := len(digits) / 2
			count := countStone(memo, digits[:mid], iteration-1) + countStone(memo, removeLeadingZeros(digits[mid:]), iteration-1)
			memo[digits][iteration] = count
			return count
		} else {
			n := Must(ParseInt64(digits))
			count := countStone(memo, strconv.FormatInt(n*2024, 10), iteration-1)
			memo[digits][iteration] = count
			return count
		}
	}

}

func removeLeadingZeros(s string) string {
	i := 0
	for i < len(s)-1 && s[i] == '0' {
		i++
	}
	return s[i:]
}
