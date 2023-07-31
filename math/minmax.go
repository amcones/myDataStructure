package math

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func Max[T constraints.Ordered](num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
