package util

import (
	"bufio"
	"log"
	"os"
)

func NewScanner(filename string) (*bufio.Scanner, func()) {
	file := Must(os.Open(filename))
	return bufio.NewScanner(bufio.NewReader(file)), func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
