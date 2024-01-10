package template

import (
	"fmt"
	"testing"
)

func TestNewBinaryIndexedTree(t *testing.T) {
	arr := []int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9}
	bit := NewBinaryIndexedTree(arr)
	fmt.Printf("arr = %d\n", arr)
	bit.Show()
}
