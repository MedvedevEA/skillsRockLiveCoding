package main

import (
	"fmt"
	"strconv"
)

var (
	numerator   = 1
	denominator = 6
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
			fBase := afterPoint[len(afterPoint)-(i+1):]
			isF := true
			for j := 2; j < 6; j++ {
				fNext := afterPoint[len(afterPoint)-(i+1)*j : len(afterPoint)-(i+1)*j+(i+1)]
				if fBase != fNext {
					isF = false
					break
				}
			}
			if isF {
				afterPoint = afterPoint[:len(afterPoint)-(i+1)*5] + "(" + fBase + ")"
				do = false
				break
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
