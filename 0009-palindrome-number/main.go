package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {

	var s = strconv.Itoa(x)
	var reversed []byte
	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}
	return (s == string(reversed))
}

/*
x = 101

input = 101	10	1	0
num   = 0	1	10	101
*/
func palindrome(x int) bool {
	var input, num int = x, 0
	for input != 0 {
		num = num*10 + input%10
		input /= 10
	}
	return (num == x)
}

func main() {
	fmt.Println(isPalindrome(101))
	fmt.Println(isPalindrome(10))
}
