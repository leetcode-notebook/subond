package template

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	n1 := &Node{1, nil}
	n2 := &Node{2, nil}
	n3 := &Node{3, nil}
	n4 := &Node{4, nil}
	n5 := &Node{5, nil}

	list := NewLinkedList()
	list.AddRear(n1)
	list.AddRear(n2)
	list.AddRear(n3)
	list.AddRear(n4)
	list.AddRear(n5)
	list.Show()

	if 5 != list.Size() {
		t.Fatalf("fail")
	}
	list.Reverse()
	list.Show()
}

func TestLinkedList_InsertSort(t *testing.T) {
	n1 := &Node{1, nil}
	n2 := &Node{2, nil}
	n3 := &Node{3, nil}
	n4 := &Node{4, nil}
	n5 := &Node{5, nil}

	list := NewLinkedList()
	list.AddRear(n5)
	list.AddRear(n4)
	list.AddRear(n3)
	list.AddRear(n2)
	list.AddRear(n1)
	list.Show()

	if 5 != list.Size() {
		t.Fatalf("fail")
	}
	list.InsertSort()
	list.Show()
}

func TestLinkedList_MoveTheLastToFront(t *testing.T) {
	n1 := &Node{1, nil}
	n2 := &Node{2, nil}
	n3 := &Node{3, nil}
	n4 := &Node{4, nil}
	n5 := &Node{5, nil}

	list := NewLinkedList()
	list.AddRear(n1)
	list.Show()
	fmt.Printf("move the last to front\n")
	list.MoveTheLastToFront()
	list.Show()

	list.AddRear(n2)
	list.Show()
	fmt.Printf("move the last to front\n")
	list.MoveTheLastToFront()
	list.Show()

	list.AddRear(n3)
	list.AddRear(n4)
	list.AddRear(n5)
	list.Show()

	if 5 != list.Size() {
		t.Fatalf("fail")
	}
	fmt.Printf("move the last to front\n")
	list.MoveTheLastToFront()
	list.Show()

	res2 := MoveTheLastToFront(list.head.next)
	pre := res2
	for pre != nil {
		t.Logf("%d\n", pre.val)
		pre = pre.next
	}
}

func TestLinkedList_DeleteEven(t *testing.T) {
	n1 := &Node{1, nil}
	n2 := &Node{2, nil}
	n3 := &Node{3, nil}
	n4 := &Node{4, nil}
	n5 := &Node{5, nil}

	list := NewLinkedList()
	list.AddRear(n1)
	list.AddRear(n2)
	list.AddRear(n3)
	list.AddRear(n4)
	list.AddRear(n5)
	list.Show()

	if 5 != list.Size() {
		t.Fatalf("fail")
	}
	list.DeleteEvenNode()
	list.Show()
}

func TestLinkedList_DeleteSmallerRight(t *testing.T) {
	n1 := &Node{1, nil}
	n2 := &Node{2, nil}
	n3 := &Node{3, nil}
	n4 := &Node{4, nil}
	n5 := &Node{5, nil}
	n6 := &Node{6, nil}

	list := NewLinkedList()
	list.AddRear(n1)
	list.AddRear(n2)
	list.AddRear(n6)
	list.AddRear(n3)
	list.AddRear(n4)
	list.AddRear(n5)
	list.Show()

	list.DeleteSmallerRight()
	list.Show()
	if 2 != list.Size() {
		t.Fatalf("fail")
	}
}
