package main

import (
	"fmt"
)

type gasengine struct {
	mpg     uint8
	gallons uint8
}

type eletcric struct {
	mpkwh uint8
	kwh   uint8
}

func (e eletcric) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}
func (e gasengine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type engine interface {
	milesLeft() uint8
}

func canMakeit(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("yes")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var myEngine gasengine = gasengine{mpg: 25, gallons: 15} // var myEngine gasengine = gasengine{25, 15}
	var mye eletcric = eletcric{40, 12}
	var myEngine2 = struct {
		mpg     uint
		gallons uint
	}{35, 16}
	canMakeit(myEngine, 80)
	canMakeit(mye, 6)
	fmt.Println(myEngine, myEngine2, mye)
}
