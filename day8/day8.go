package day8

import (
	. "advent/util"
)

type pos struct{ x, y int }

func Part1(filename string) int64 {
	var table [][]byte
	locations := make(map[byte][]pos)
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	row := 0
	for scanner.Scan() {
		var rowData []byte
		text := scanner.Text()
		for col := 0; col < len(text); col++ {
			rowData = append(rowData, text[col])
			if text[col] != '.' && text[col] != '#' {
				locations[text[col]] = append(locations[text[col]], pos{row, col})
			}
		}
		table = append(table, rowData)
		row++
	}
	antinodeLocs := make(map[int]map[int]struct{})
	for _, coords := range locations {
		for i := 0; i < len(coords); i++ {
			pos1 := coords[i]
			for j := i + 1; j < len(coords); j++ {
				pos2 := coords[j]
				deltaX := pos2.x - pos1.x
				deltaY := pos2.y - pos1.y

				// step forward
				r := pos2.x + deltaX
				c := pos2.y + deltaY
				if 0 <= r && r < len(table) && 0 <= c && c < len(table[r]) {
					if _, ok := antinodeLocs[r]; !ok {
						antinodeLocs[r] = make(map[int]struct{})
					}
					antinodeLocs[r][c] = struct{}{}
				}
				// step backward
				r = pos1.x - deltaX
				c = pos1.y - deltaY
				if 0 <= r && r < len(table) && 0 <= c && c < len(table[r]) {
					if _, ok := antinodeLocs[r]; !ok {
						antinodeLocs[r] = make(map[int]struct{})
					}
					antinodeLocs[r][c] = struct{}{}
				}
			}
		}
	}
	var count int64
	for row := range antinodeLocs {
		count += int64(len(antinodeLocs[row]))
	}
	return count
}

func Part2(filename string) int64 {
	return 0
}
