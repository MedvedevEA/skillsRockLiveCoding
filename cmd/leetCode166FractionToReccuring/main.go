package main

import (
	"fmt"
	"strconv"
)

var (
	numerator   = 1
	denominator = 214748364
)

func fractionToDecimal(numerator int, denominator int) string {

	beforPoint := ""
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		beforPoint = beforPoint + "-"
	}
	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}
	if numerator >= denominator {
		beforPoint = beforPoint + strconv.Itoa(numerator/denominator)
		numerator = numerator % denominator
	} else {
		beforPoint = beforPoint + "0"
	}
	afterPoint := ""
	do := true

	for numerator%denominator != 0 && do {
		check := numerator*10/denominator > 0
		afterPoint = afterPoint + strconv.Itoa(numerator*10/denominator)
		numerator = numerator * 10 % denominator
		if len(afterPoint) > 10000 {
			do = false
		}

		for i := 0; i < len(afterPoint)/5 && check; i++ {

			f1 := afterPoint[len(afterPoint)-(i+1):]
			f2 := afterPoint[len(afterPoint)-(i+1)*2 : len(afterPoint)-(i+1)*2+(i+1)]
			f3 := afterPoint[len(afterPoint)-(i+1)*3 : len(afterPoint)-(i+1)*3+(i+1)]
			f4 := afterPoint[len(afterPoint)-(i+1)*4 : len(afterPoint)-(i+1)*4+(i+1)]
			f5 := afterPoint[len(afterPoint)-(i+1)*5 : len(afterPoint)-(i+1)*5+(i+1)]

			if f1 == f2 && f2 == f3 && f3 == f4 && f4 == f5 {
				afterPoint = afterPoint[:len(afterPoint)-(i+1)*5] + "(" + f1 + ")"
				do = false
			}
		}

	}
	if len(afterPoint) > 0 {
		beforPoint = beforPoint + "."
	}

	return beforPoint + afterPoint
}

func main() {
	result := fractionToDecimal(numerator, denominator)
	fmt.Println(result)

}
