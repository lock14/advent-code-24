package day12

import (
	. "advent/util"
	"advent/util/set"
)

type cell struct {
	i, j int
}

var directions = []cell{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func Part1(filename string) int64 {
	var table []string
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		table = append(table, scanner.Text())
	}
	visited := set.New[cell]()
	var cost int64
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			area, perimeter := scanRegion(table, visited, table[i][j], i, j)
			cost += area * perimeter
		}
	}
	return cost
}

func scanRegion(table []string, visited *set.Set[cell], target byte, i int, j int) (int64, int64) {
	if i < 0 || j < 0 || i >= len(table) || j >= len(table[i]) {
		return 0, 1
	}
	if visited.Contains(cell{i, j}) && table[i][j] == target {
		return 0, 0
	} else if visited.Contains(cell{i, j}) {
		return 0, 1
	}
	if table[i][j] == target {
		visited.Add(cell{i, j})
		var area, perimeter int64
		area++
		for _, d := range directions {
			a, p := scanRegion(table, visited, target, i+d.i, j+d.j)
			area += a
			perimeter += p
		}
		return area, perimeter
	} else {
		return 0, 1
	}
}

func Part2(filename string) int64 {
	return 0
}
