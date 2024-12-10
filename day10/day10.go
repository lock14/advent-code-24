package day10

import (
	. "advent/util"
)

func Part1(filename string) int64 {
	var table []string
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		table = append(table, scanner.Text())
	}
	var count int64
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == '0' {
				summitPos := make(map[int]map[int]struct{})
				trailHeadScore(table, i, j, '0', summitPos)
				var score int64
				for _, cols := range summitPos {
					score += int64(len(cols))
				}
				count += score
			}
		}
	}
	return count
}

func trailHeadScore(table []string, i int, j int, target byte, summitPos map[int]map[int]struct{}) {
	if i < 0 || i >= len(table) || j < 0 || j >= len(table[i]) {
		return
	}
	if table[i][j] == target && target == '9' {
		if _, ok := summitPos[i]; !ok {
			summitPos[i] = make(map[int]struct{})
		}
		summitPos[i][j] = struct{}{}
	} else if table[i][j] == target {
		trailHeadScore(table, i-1, j, target+1, summitPos) // up
		trailHeadScore(table, i+1, j, target+1, summitPos) // down
		trailHeadScore(table, i, j-1, target+1, summitPos) // left
		trailHeadScore(table, i, j+1, target+1, summitPos) // right
	}
}

func Part2(filename string) int64 {
	return 0
}
