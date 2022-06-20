package main

import "strings"

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

func main() {
	//fmt.Println(problem.countPrimes(27))

	// var l1, l2 ListNode
	// l1.Val = 1
	// l2.Val = 2

	// mergeTwoLists(&l1, &l2)

	isPalindrome("0P")
}
