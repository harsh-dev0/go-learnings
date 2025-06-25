package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var Num1 int = 2 // int 8 16 32 64 uint only positive
	Num2 := 3
	var float1 float32 = 1.23
	float2 := 1.45 //float64
	fmt.Println(Num1 + Num2)
	fmt.Println((float1 + float32(float2)))
	var myStr string = "hey jey"
	fmt.Print(myStr, Num1, Num2)
	fmt.Println(utf8.RuneCountInString("y"))
	var myRune rune = 'a'

	fmt.Println(myRune)

	var myBool bool = true
	fmt.Println(myBool)

	var var1, var2 = 1, 2
	const pie float64 = 3.14
	fmt.Println(var1, var2, pie)

	var a = "jk"
	PrintMe(a)

	var b, c, e = intDiv(2, 0)
	if e != nil {
		fmt.Println(e.Error())
	} else if c == 0 {
		fmt.Printf("Result is %v", b)
	} else {
		fmt.Println(b, c, e)
	}
}
