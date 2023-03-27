package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 **/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	var useCases = []struct {
		l1 []int
		l2 []int
	}{
		{
			l1: []int{2, 4, 3},
			l2: []int{5, 6, 4},
		},
	}

	for _, useCase := range useCases {

		// var ln1 = &ListNode{Val: useCase.l1[0]}
		// for i := 1; i < len(useCase.l1); i++ {
		// 	templn
		// }
		// var ln2 = &ListNode{Val: useCase.l2[0]}
		// for i := 1; i < len(useCase.l2); i++ {

		// }

		fmt.Println(useCase)
	}

}
