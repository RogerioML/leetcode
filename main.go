package main

import (
	"fmt"
	"leetcode/roman"
)

func main() {
	s := "IV"
	fmt.Println(roman.RomanToInt(s))
	s = "LVIII"
	fmt.Println(roman.RomanToInt(s))
	s = "MCMXCIV"
	fmt.Println(roman.RomanToInt(s))
}
