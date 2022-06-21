package main

import (
	"fmt"
	"strings"
)

func reverseString(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var toReturn ListNode
	if list1.Val < list2.Val {
		toReturn.Val = list1.Val
		list1 = list1.Next
	} else {
		toReturn.Val = list2.Val
		list2 = list2.Next
	}
	toReturn.Next = mergeTwoLists(list1, list2)

	return &toReturn
}

func isPalindrome(s string) bool {
	lowerString := strings.ToLower(s)
	maxLen := len(lowerString) - 1
	for i, j := 0, maxLen; i < j; i, j = i+1, j-1 {
		for !(('a' <= lowerString[i] && lowerString[i] <= 'z') || ('0' <= lowerString[i] && lowerString[i] <= '9')) {
			i++
			if i > maxLen {
				return true
			}
		}
		for !(('a' <= lowerString[j] && lowerString[j] <= 'z') || ('0' <= lowerString[j] && lowerString[j] <= '9')) {
			j--
		}

		if lowerString[i] != lowerString[j] {
			return false
		}
	}
	return true

}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	currentDistance := 0
	longestLadders := make([]int, ladders)
	bricksRemaining := bricks

	for currentDistance < len(heights)-1 {
		if heights[currentDistance] >= heights[currentDistance+1] {
			currentDistance++
		} else {
			bricksNeeded := heights[currentDistance+1] - heights[currentDistance]

			if ladders > 0 && bricksNeeded > longestLadders[ladders-1] {
				for ladderIndex := 0; ladderIndex < ladders && bricksNeeded > 0; ladderIndex++ {
					if longestLadders[ladderIndex] < bricksNeeded {
						longestLadders[ladderIndex], bricksNeeded = bricksNeeded, longestLadders[ladderIndex]
					}
				}
			}

			if bricksNeeded > bricksRemaining {
				return currentDistance
			}
			bricksRemaining -= bricksNeeded
			currentDistance++
		}
	}

	return len(heights) - 1
}
func main() {
	//fmt.Println(problem.countPrimes(27))

	// var l1, l2 ListNode
	// l1.Val = 1
	// l2.Val = 2

	// mergeTwoLists(&l1, &l2)

	heights := []int{4, 12, 2, 7, 3, 18, 20, 3, 19}

	fmt.Println(furthestBuilding(heights, 10, 2))

	heights2 := []int{14, 3, 19, 3}
	fmt.Println(furthestBuilding(heights2, 17, 0))
}
