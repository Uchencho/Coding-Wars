package main

import (
	"fmt"
	"log"
	"strconv"
)

// https://leetcode.com/problems/add-two-numbers/ <<< Unsolved

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	sumOne := getSum(l1)
	sumTwo := getSum(l2)

	r := sumOne + sumTwo

	w := &ListNode{}
	for i, v := range reverseString(strconv.Itoa(r)) {
		val, _ := strconv.Atoi(string(v))
		if i == 0 {
			w = &ListNode{Val: val}
			continue
		}
		w = createListNode(val, w)
	}

	log.Printf("%v", getSum(w))
}

func getSum(l *ListNode) int {
	s := ""
	for {
		s = fmt.Sprintf("%d", l.Val) + s
		if l.Next == nil || (l == &ListNode{}) {
			break
		}
		l = l.Next
	}
	count, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func createListNode(n int, l *ListNode) *ListNode {
	if l.Next != nil {
		createListNode(n, l.Next)
	}
	l.Next = &ListNode{Val: n}
	return l
}

func reverseString(s string) string {
	var r string
	for _, v := range s {
		r = string(v) + r
	}
	return r
}
