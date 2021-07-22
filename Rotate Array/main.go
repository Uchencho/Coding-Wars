package main

import "log"

// https://app.codility.com/c/run/trainingPY88JW-VKK/

func Solution(A []int, K int) []int {

	if len(A) == K {
		return A
	}

	if len(A) <= 1 {
		return A
	}

	for i := 0; i < K; i++ {
		A = rotate(A)
	}
	return A
}

func rotate(in []int) []int {
	out := make([]int, 2)
	out[0] = in[len(in)-1]                // <<< set the last value of in to first value of out
	out[1] = in[0]                        // <<< set the first value of in to second value of out
	out = append(out, in[1:len(in)-1]...) // unpack the rest
	return out
}

func main() {
	in := []int{1000}
	log.Println(Solution(in, 5))
}

/*
// Sample input
[3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]

*/
