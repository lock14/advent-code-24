package day2

import (
	"advent/util"
	"bufio"
	"iter"
	"log"
	"os"
	"slices"
	"strings"
)

func Part1(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	numSafe := int64(0)
	for scanner.Scan() {
		report, err := util.MapSliceErr(strings.Fields(scanner.Text()), util.ParseInt64)
		if err != nil {
			log.Fatal(err)
		}
		safe := safe(slices.All(report), -1) || safe(slices.Backward(report), -1)
		if safe {
			numSafe += 1
		}
	}
	return numSafe
}

func Part2(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	numSafe := int64(0)
	for scanner.Scan() {
		report, err := util.MapSliceErr(strings.Fields(scanner.Text()), util.ParseInt64)
		if err != nil {
			log.Fatal(err)
		}
		safe := safeWithRemoval(slices.All(report), len(report)) || safeWithRemoval(slices.Backward(report), len(report))
		if safe {
			numSafe += 1
		}
	}
	return numSafe
}

func safe(itr iter.Seq2[int, int64], skip int) bool {
	next, stop := iter.Pull2(itr)
	defer stop()
	idx, a, _ := next()
	if idx == skip {
		idx, a, _ = next()
	}
	idx, b, ok := next()
	if idx == skip {
		idx, b, ok = next()
	}
	for ok {
		if diff := a - b; a <= b || diff < 1 || diff > 3 {
			return false
		}
		a = b
		idx, b, ok = next()
		if idx == skip {
			idx, b, ok = next()
		}
	}
	return true
}

func safeWithRemoval(itr iter.Seq2[int, int64], max int) bool {
	for i := -1; i < max; i++ {
		if safe(itr, i) {
			return true
		}
	}
	return false
}
