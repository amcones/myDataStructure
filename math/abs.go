package math

import "golang.org/x/exp/constraints"

func Abs[T constraints.Signed](num int) int {
	if num >= 0 {
		return num
	}
	return -num
}
