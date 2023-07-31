package math

import "golang.org/x/exp/constraints"

func min[T constraints.Ordered](num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func max[T constraints.Ordered](num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
