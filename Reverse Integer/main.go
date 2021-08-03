package main

// https://leetcode.com/problems/reverse-integer

import (
	"log"
	"strconv"
)

func reverse(x int) int {
	result := ""
	positiveLimit := 2147483647
	negativeLimit := -2147483647

	var negativeNumber bool
	for _, value := range strconv.Itoa(x) {
		if string(value) == "-" {
			negativeNumber = true
			continue
		}
		result = string(value) + result
	}

	if negativeNumber {
		result = "-" + result
	}
	v, _ := strconv.Atoi(result)
	if v > positiveLimit || v < negativeLimit {
		return 0
	}
	return v
}

func main() {
	log.Println(reverse(-2147483649))
	log.Println(reverse(-2147483600))
}
