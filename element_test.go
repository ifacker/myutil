package myutil

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicateElement(t *testing.T) {
	as := []string{
		"abc",
		"def",
		"ghi",
		"abc",
	}

	fmt.Println(as)

	bs := RemoveDuplicateElement(as)
	fmt.Println(bs)
}

func TestRemoveDuplicateElementInt(t *testing.T) {
	as := []int{
		1, 2, 3, 4, 5, 1,
	}

	fmt.Println(as)

	bs := RemoveDuplicateElementInt(as)
	fmt.Println(bs)
}
