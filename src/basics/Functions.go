package main

import (
	"errors"
	"fmt"
)

func PrintMe(a string) {
	fmt.Println(a)
}

func intDiv(n int, d int) (int, int, error) {
	var err error
	if d == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	return n / d, n % d, err
}
