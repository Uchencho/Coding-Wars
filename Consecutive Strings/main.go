package main

// https://www.codewars.com/kata/56a5d994ac971f1ac500003e/train/go

import (
	"log"
	"strings"
)

func LongestConsec(strarr []string, k int) string {
	n := len(strarr)
	if n == 0 || k > n || k <= 0 {
		return ""
	}

	result := ""

	if k == 1 {
		for _, word := range strarr {
			if len(word) > len(result) {
				result = word
			}
		}
		return result
	}

	for i := 0; i <= len(strarr)-k; i++ {
		s := getSubset(strarr, i, k+i)
		if len(s) > len(result) {
			result = s
		}
	}
	return result
}

func getSubset(all []string, start, end int) string {
	return strings.Join(all[start:end], "")
}

func main() {
	r := []string{"zone", "abigail", "theta", "form", "libe", "zas"}
	log.Println(LongestConsec(r, 2))
}
