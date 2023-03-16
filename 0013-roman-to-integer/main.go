package main

import (
	"fmt"
)

func romanToInt(s string) (n int) {

	var prev int
	for i := len(s) - 1; i >= 0; i-- {
		curr := value(s[i])

		// if previous value is less than current
		// add to result as the number is not part of
		// a larger number set
		// else, the previous number was part of a larger
		// number, hence reduce the larger number from previous one.
		// eg: s = MCMXCIV	expected = 1994
		// c	prev	curr				n
		// V	0		5		0<5			0+5=5
		// I	5		1		5>1			5-1=4
		// C	4		100		4<100		4+100=104
		// X	104		10		104>10		104-10=94
		// M	94		1000	94<1000		94+1000=1094
		// C	1094	100		1094>100	1094-100=994
		// M	994		1000	994<1000	994+1000=1994
		if prev <= curr {
			n += curr
		} else {
			n -= curr
		}
		fmt.Println(prev, curr, n)
		prev = curr
	}
	return
}

func value(c byte) int {
	switch c {
	case 'M':
		return 1000
	case 'D':
		return 500
	case 'C':
		return 100
	case 'L':
		return 50
	case 'X':
		return 10
	case 'V':
		return 5
	case 'I':
		return 1
	}
	return 0
}

func main() {

	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
