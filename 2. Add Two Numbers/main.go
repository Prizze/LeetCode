package main

import "fmt"

// Структура ноды
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0				// Перенос с предыдущего разряда
	head := &ListNode{}		// dummy ptr		
	tail := head

	for l1 != nil || l2 != nil || carry != 0 {
		var n1, n2, n = 0, 0, 0

		// Первая цифра
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		// Вторая цифра
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		n = n1 + n2 + carry
		// Формируем список
		newNode := &ListNode{Val: n % 10}
		tail.Next = newNode
		tail = newNode

		carry = n / 10
	}
	return head.Next
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

	l := addTwoNumbers(l1, l2)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
