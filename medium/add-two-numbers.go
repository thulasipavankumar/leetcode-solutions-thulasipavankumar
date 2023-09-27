package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	answer := ListNode{}
	answer.Next = nil
	r := &answer

	var carry int
	for result := r; l1 != nil || l2 != nil; {
		var currentSum int
		if l1 != nil {
			currentSum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			currentSum += l2.Val
			l2 = l2.Next
		}
		currentSum += carry
		if currentSum > 9 {
			carry = 1
			currentSum %= 10

		} else {
			carry = 0
		}

		result.Next = &ListNode{Val: currentSum, Next: nil}
		result = result.Next

		fmt.Println(currentSum, carry)
		answer.PrintList()
	}
	if carry != 0 {
		for r.Next != nil {
			r = r.Next
		}
		r.Next = &ListNode{carry, nil}
	}
	answer.PrintList()
	return answer.Next
}

func fillValues(l *ListNode, arr []int) {
	for i, head := 1, l; i < len(arr); i++ {
		temp := ListNode{Val: arr[i], Next: nil}
		head.Next = &temp
		head = head.Next
	}
}
func (list ListNode) PrintList() {
	for l := &list; l != nil; l = l.Next {
		fmt.Print("->", l.Val)
	}
	fmt.Println()
}
func main() {
	arr1, arr2 := []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}
	//arr1,arr2 := []int{2,4,3}, []int{5,6,4}
	l1 := ListNode{Val: arr1[0], Next: nil}
	l2 := ListNode{Val: arr2[0], Next: nil}

	fillValues(&l1, arr1)
	fillValues(&l2, arr2)
	fmt.Println(l1, l2)
	fmt.Println(addTwoNumbers(&l1, &l2))

}
