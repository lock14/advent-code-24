package day5

import (
	. "advent/util"
	"advent/util/set"
	"strings"
)

func Part1(filename string) int64 {
	numbers := make(map[int64]*set.Set[int64])
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		parts := strings.Split(scanner.Text(), "|")
		a := Must(ParseInt64(parts[0]))
		b := Must(ParseInt64(parts[1]))
		if numbers[a] == nil {
			numbers[a] = set.New[int64]()
		}
		numbers[a].Add(b)
	}
	result := int64(0)
	for scanner.Scan() {
		list := Must(MapSliceErr(strings.Split(scanner.Text(), ","), ParseInt64))
		if idx := correctMiddle(numbers, list); idx != -1 {
			result += list[idx]
		}
	}
	return result
}

func correctMiddle(numbers map[int64]*set.Set[int64], list []int64) int {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if numbers[list[i]] == nil || !numbers[list[i]].Contains(list[j]) {
				return -1
			}
		}
	}
	return len(list) / 2
}

func Part2(filename string) int64 {
	return 0
}
