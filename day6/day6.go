package day6

import (
	. "advent/util"
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
	row, col := findStart(table)
	return countCells(table, row, col)
}

func countCells(table [][]uint8, row int, col int) int64 {
	var count int64
	direction := Up
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

func findStart(table [][]uint8) (int, int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == '^' {
				return i, j
			}
		}
	}
	panic("no start position found")
}

func Part2(filename string) int64 {
	return 0
}
