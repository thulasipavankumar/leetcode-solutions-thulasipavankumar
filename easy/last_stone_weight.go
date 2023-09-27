package main

import (
	"fmt"
	//    "sort"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
	size int
}

func printList(l *List) {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println()
	return
}

func (l *List) increment() {
	l.size += 1
}
func (l *List) insert(val int) {

	newNode := &Node{data: val}
	defer l.increment()
	if l.head == nil {
		l.head = newNode
		return
	}
	if l.head.data < val {
		temp := l.head
		newNode.next = temp
		l.head = newNode
		return
	}
	curr := l.head
	for curr.next != nil && val < curr.next.data {
		curr = curr.next
	}
	temp := curr.next
	curr.next = newNode
	newNode.next = temp

}
func (l *List) decrement() {
	l.size -= 1
}
func (l *List) pop() int {
	if l.head == nil {
		return 0
	}
	l.decrement()
	temp := l.head.data
	l.head = l.head.next
	return temp
}
func lastStoneWeight(stones []int) int {
	//sort.Ints(stones)
	l := &List{}
	for i := len(stones) - 1; i >= 0; i-- {
		l.insert(stones[i])
	}
	for l.size >= 2 {
		fmt.Println("before modyfying")
		printList(l)
		one := l.pop()
		two := l.pop()
		if one-two > 0 {
			l.insert(one - two)
		}

		fmt.Println("After modyfying")
		printList(l)
		fmt.Println()
	}
	fmt.Println("Left out with")
	printList(l)
	if l.size < 1 {
		return 0
	} else {
		return l.pop()
	}

}

func main() {
	intSlice := []int{2, 7, 4, 1, 8, 1}
	//intSlice := []int{5, 2, 6, 3, 1, 4}

	//intSlice := []int{11,7,6,2,1}
	fmt.Println(lastStoneWeight(intSlice))
}
