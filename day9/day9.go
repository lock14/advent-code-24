package day9

import (
	. "advent/util"
	"bufio"
)

func Part1(filename string) int64 {
	ID := 0
	freeID := -1
	file := true
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	scanner.Split(bufio.ScanBytes)
	var inflate []int
	for scanner.Scan() {
		n := int(Must(ParseInt64(scanner.Text())))
		if file {
			for i := 0; i < n; i++ {
				inflate = append(inflate, ID)
			}
			ID++
		} else {
			for i := 0; i < n; i++ {
				inflate = append(inflate, freeID)
			}
			freeID--
		}
		file = !file
	}
	i := nextFree(inflate, 0)
	j := nextFile(inflate, len(inflate)-1)
	for i < j {
		if inflate[i] < 0 && inflate[j] >= 0 {
			inflate[i], inflate[j] = inflate[j], inflate[i]
		}
		i = nextFree(inflate, i)
		j = nextFile(inflate, j)
	}

	var checksum int64
	for i := 0; i < len(inflate); i++ {
		if inflate[i] >= 0 {
			checksum += int64(i * inflate[i])
		}
	}

	return checksum
}

func Part2(filename string) int64 {
	ID := 0
	freeID := -1
	file := true
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	scanner.Split(bufio.ScanBytes)
	var inflate []int
	blockLengths := make(map[int]int)
	for scanner.Scan() {
		n := int(Must(ParseInt64(scanner.Text())))
		if file {
			for i := 0; i < n; i++ {
				inflate = append(inflate, ID)
			}
			blockLengths[ID] = n
			ID++
		} else {
			for i := 0; i < n; i++ {
				inflate = append(inflate, freeID)
			}
			blockLengths[freeID] = n
			freeID--
		}
		file = !file
	}
	j := nextFile(inflate, len(inflate)-1)
	for j >= 0 {
		jBlockLength := blockLengths[inflate[j]]
		i := nextFree(inflate, 0)
		for i < j {
			iBlockLength := blockLengths[inflate[i]]
			if iBlockLength >= jBlockLength {
				for jBlockLength > 0 {
					inflate[i], inflate[j] = inflate[j], inflate[i]
					i++
					j--
					iBlockLength--
					jBlockLength--
				}
				if iBlockLength > 0 {
					blockLengths[inflate[i]] = iBlockLength
				}
				break
			}
			i = nextFree(inflate, i+blockLengths[inflate[i]])
		}
		j = nextFile(inflate, j-jBlockLength)
	}

	var checksum int64
	for i := 0; i < len(inflate); i++ {
		if inflate[i] >= 0 {
			checksum += int64(i * inflate[i])
		}
	}

	return checksum
}

func nextFree(inflate []int, i int) int {
	for i < len(inflate) && inflate[i] >= 0 {
		i++
	}
	return i
}

func nextFile(inflate []int, j int) int {
	for 0 <= j && inflate[j] < 0 {
		j--
	}
	return j
}
