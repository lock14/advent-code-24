package day6

import (
	. "advent/util"
	"advent/util/set"
)

const (
	Up = iota
	Down
	Left
	Right
)

func Part1(filename string) int64 {
	var table [][]uint8
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		s := scanner.Text()
		chars := make([]uint8, 0, len(s))
		for i := 0; i < len(s); i++ {
			chars = append(chars, s[i])
		}
		table = append(table, chars)
	}
	row, col, direction := findStart(table)
	return countCells(table, row, col, direction)
}

func Part2(filename string) int64 {
	var table [][]uint8
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		s := scanner.Text()
		chars := make([]uint8, 0, len(s))
		for i := 0; i < len(s); i++ {
			chars = append(chars, s[i])
		}
		table = append(table, chars)
	}
	row, col, direction := findStart(table)
	return countCycles(table, row, col, direction)
}

func countCells(table [][]uint8, row, col, direction int) int64 {
	var count int64
	for 0 <= row && row < len(table) && 0 <= col && col < len(table[row]) {
		if table[row][col] == '#' {
			row, col = reverse(row, col, direction)
			direction = turnRight(direction)
		} else if table[row][col] != 'X' {
			count++
			table[row][col] = 'X'
		}
		row, col = advance(row, col, direction)
	}
	return count
}

func countCycles(table [][]uint8, startRow, startCol, startDirection int) int64 {
	countCells(table, startRow, startCol, startDirection)
	count := int64(0)
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == 'X' && !(i == startRow && j == startCol) {
				table[i][j] = '#'
				if isCycle(table, startRow, startCol, startDirection) {
					count++
				}
				table[i][j] = 'X'
			}
		}
	}
	return count
}

type pos struct {
	row, col, direction int
}

func isCycle(table [][]uint8, row int, col int, direction int) bool {
	positions := set.New[pos]()
	for 0 <= row && row < len(table) && 0 <= col && col < len(table[row]) {
		if table[row][col] == '#' {
			row, col = reverse(row, col, direction)
			direction = turnRight(direction)
		} else {
			pos := pos{row, col, direction}
			if positions.Contains(pos) {
				return true
			}
			positions.Add(pos)
		}
		row, col = advance(row, col, direction)
	}
	return false
}

func turnRight(direction int) int {
	switch direction {
	case Up:
		return Right
	case Down:
		return Left
	case Left:
		return Up
	case Right:
		return Down
	}
	panic("invalid direction")
}

func reverse(row, col, direction int) (int, int) {
	switch direction {
	case Up:
		return row + 1, col
	case Down:
		return row - 1, col
	case Left:
		return row, col + 1
	case Right:
		return row, col - 1
	}
	panic("invalid direction")
}

func advance(row, col, direction int) (int, int) {
	switch direction {
	case Up:
		return row - 1, col
	case Down:
		return row + 1, col
	case Left:
		return row, col - 1
	case Right:
		return row, col + 1
	}
	panic("invalid direction")
}

func findStart(table [][]uint8) (int, int, int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == '^' {
				return i, j, Up
			} else if table[i][j] == '>' {
				return i, j, Right
			} else if table[i][j] == 'v' {
				return i, j, Down
			} else if table[i][j] == '<' {
				return i, j, Left
			}
		}
	}
	panic("no start position found")
}
