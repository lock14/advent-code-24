package day3

import (
	. "advent/util"
	"regexp"
)

var pattern = regexp.MustCompile("(mul)\\((\\d+),(\\d+)\\)|(do)\\(\\)|(don't)\\(\\)")

func Part1(filename string) int64 {
	result := int64(0)
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		if matches := pattern.FindAllStringSubmatch(scanner.Text(), -1); matches != nil {
			for _, m := range matches {
				if m[1] == "mul" {
					a := Must(ParseInt64(m[2]))
					b := Must(ParseInt64(m[3]))
					result += a * b
				}
			}
		}
	}
	return result
}

func Part2(filename string) int64 {
	result := int64(0)
	do := true
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		if matches := pattern.FindAllStringSubmatch(scanner.Text(), -1); matches != nil {
			for _, m := range matches {
				if m[1] == "mul" && do {
					a := Must(ParseInt64(m[2]))
					b := Must(ParseInt64(m[3]))
					result += a * b
				} else if m[4] == "do" {
					do = true
				} else if m[5] == "don't" {
					do = false
				}
			}
		}
	}
	return result
}
