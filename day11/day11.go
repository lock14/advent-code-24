package day11

import (
	. "advent/util"
	"strconv"
	"strings"
)

type stone struct {
	digits string
	next   *stone
	prev   *stone
}

func Part1(filename string) int64 {
	list := sentinel()
	scanner, closeFunc := NewScanner(filename)
	defer closeFunc()
	for scanner.Scan() {
		for _, digits := range strings.Fields(scanner.Text()) {
			insertAfter(list.prev, &stone{digits: digits})
		}
	}
	for i := 0; i < 25; i++ {
		blink(list)
	}
	var count int64
	cur := list.next
	for cur != list {
		count++
		cur = cur.next
	}
	return count
}

func blink(list *stone) {
	cur := list.next
	for cur != list {
		if cur.digits == "0" {
			cur.digits = "1"
		} else if len(cur.digits)%2 == 0 {
			mid := len(cur.digits) / 2
			insertAfter(cur, &stone{digits: removeLeadingZeros(cur.digits[mid:])})
			cur.digits = cur.digits[:mid]
			cur = cur.next
		} else {
			n := Must(ParseInt64(cur.digits))
			cur.digits = strconv.FormatInt(n*2024, 10)
		}
		cur = cur.next
	}
}

func removeLeadingZeros(s string) string {
	i := 0
	for i < len(s)-1 && s[i] == '0' {
		i++
	}
	return s[i:]
}

func Part2(filename string) int64 {
	return 0
}

func sentinel() *stone {
	s := &stone{}
	s.next = s
	s.prev = s
	return s
}

func insertBefore(node *stone, s *stone) {
	s.next = node
	s.prev = node.prev
	node.prev.next = s
	node.prev = s
}

func insertAfter(node *stone, s *stone) {
	s.prev = node
	s.next = node.next
	node.next.prev = s
	node.next = s
}
