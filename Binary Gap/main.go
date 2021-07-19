package main

import (
	"log"
	"strconv"
)

// https://app.codility.com/programmers/lessons/1-iterations/binary_gap/ <<< codility

func Solution(N int) int {

	result := 0
	count := 0
	foundOne := false

	b := strconv.FormatInt(int64(N), 2)
	for _, val := range b {

		switch string(val) {
		case "1":
			if !foundOne {
				foundOne = true
				continue
			}
			if result < count {
				result = count
			}
			count = 0
		case "0":
			if foundOne {
				count += 1
			}
		}
	}
	return result
}

func main() {
	log.Println(Solution(561892))
}
