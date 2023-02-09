package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("abc")
	stack.Push("aab")
	stack.Push("acd")
	stack.Push("jqwle")
	stack.Push("abc")
	stack.Push("aab")
	stack.Push("acd")
	stack.Push("jqwle")
	fmt.Println(stack.Size())
	_ = stack.Pop()
	_ = stack.Pop()
	fmt.Println(stack.Size())
	for !stack.Empty() {
		res, _ := stack.Top()
		fmt.Println(res)
		_ = stack.Pop()
	}
}
