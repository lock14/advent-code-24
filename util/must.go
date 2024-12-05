package util

import "log"

func Must[T any](t T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return t
}
