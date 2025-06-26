package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var myString = ("rεsumε")
	count := utf8.RuneCountInString(myString)
	var indexed = myString[0]
	fmt.Println(indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	var builder strings.Builder
	builder.WriteString("Hello, ")
	builder.WriteString("world!")
	result := builder.String()
	fmt.Print(result, count, len(myString))
}
