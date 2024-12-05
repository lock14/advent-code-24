package day4

import (
	. "advent/util"
)

var XMas = []byte{'X', 'M', 'A', 'S'}

func Part1(filename string) int64 {
	var table []string
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		table = append(table, scanner.Text())
	}
	result := int64(0)
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == 'X' {
				result += search(table, i, j, 0, up)
				result += search(table, i, j, 0, down)
				result += search(table, i, j, 0, left)
				result += search(table, i, j, 0, right)
				result += search(table, i, j, 0, upperLeft)
				result += search(table, i, j, 0, upperRight)
				result += search(table, i, j, 0, lowerLeft)
				result += search(table, i, j, 0, lowerRight)
			}
		}
	}
	return result
}

func Part2(filename string) int64 {
	var table []string
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		table = append(table, scanner.Text())
	}
	result := int64(0)
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if isXMAS(table, i, j) {
				result++
			}
		}
	}
	return result
}

func search(table []string, row, col, idx int, advance func(int, int) (int, int)) int64 {
	if idx == len(XMas) {
		return 1
	}
	if 0 <= row && row < len(table) && 0 <= col && col < len(table[row]) {
		if table[row][col] == XMas[idx] {
			i, j := advance(row, col)
			return search(table, i, j, idx+1, advance)
		}
	}
	return 0
}

func isXMAS(table []string, i int, j int) bool {
	if i == 0 || j == 0 || i == len(table)-1 || j == len(table[i])-1 {
		// we are at an edge
		return false
	}
	firstDiagonal := false
	secondDiagonal := false
	if (table[i-1][j-1] == 'M' && table[i][j] == 'A' && table[i+1][j+1] == 'S') ||
		(table[i-1][j-1] == 'S' && table[i][j] == 'A' && table[i+1][j+1] == 'M') {
		firstDiagonal = true
	}
	if (table[i+1][j-1] == 'M' && table[i][j] == 'A' && table[i-1][j+1] == 'S') ||
		(table[i+1][j-1] == 'S' && table[i][j] == 'A' && table[i-1][j+1] == 'M') {
		secondDiagonal = true
	}
	return firstDiagonal && secondDiagonal
}

func up(r, c int) (int, int) {
	return r - 1, c
}

func down(r, c int) (int, int) {
	return r + 1, c
}

func left(r, c int) (int, int) {
	return r, c - 1
}

func right(r, c int) (int, int) {
	return r, c + 1
}

func upperLeft(r, c int) (int, int) {
	return r - 1, c - 1
}

func upperRight(r, c int) (int, int) {
	return r - 1, c + 1
}

func lowerLeft(r, c int) (int, int) {
	return r + 1, c - 1
}

func lowerRight(r, c int) (int, int) {
	return r + 1, c + 1
}
