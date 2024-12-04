package day3

import (
	"advent/util"
	"bufio"
	"log"
	"os"
	"regexp"
)

var pattern = regexp.MustCompile("(mul)\\((\\d+),(\\d+)\\)|(do)\\(\\)|(don't)\\(\\)")

func Part1(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	result := int64(0)
	for scanner.Scan() {
		if matches := pattern.FindAllStringSubmatch(scanner.Text(), -1); matches != nil {
			for _, m := range matches {
				if m[1] == "mul" {
					a, err := util.ParseInt64(m[2])
					if err != nil {
						log.Fatal(err)
					}
					b, err := util.ParseInt64(m[3])
					if err != nil {
						log.Fatal(err)
					}
					result += a * b
				}
			}
		}
	}
	return result
}

func Part2(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	bufio.NewReader(f)
	scanner := bufio.NewScanner(f)
	result := int64(0)
	do := true
	for scanner.Scan() {
		if matches := pattern.FindAllStringSubmatch(scanner.Text(), -1); matches != nil {
			for _, m := range matches {
				if m[1] == "mul" && do {
					a, err := util.ParseInt64(m[2])
					if err != nil {
						log.Fatal(err)
					}
					b, err := util.ParseInt64(m[3])
					if err != nil {
						log.Fatal(err)
					}
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
