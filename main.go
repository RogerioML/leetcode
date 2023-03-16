package main

import (
	"leetcode/prefix"
	"log"
)

func main() {
	strs := []string{"apple", "apple", "apple"}
	log.Println(prefix.LongestCommonPrefix(strs))
}
