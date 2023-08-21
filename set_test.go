package myutil

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	// add
	mySet := make(Set)
	mySet.Add("apple")
	mySet.Add("banana")
	mySet.Add("apple") // 这个不会重复添加

	fmt.Println(mySet.Contains("apple"))  // 输出 true
	fmt.Println(mySet.Contains("orange")) // 输出 false

	// addAll
	stringSet := make(Set)
	stringSet.AddAll([]string{"apple", "banana", "apple", "orange"})

	fmt.Println("Set contains 'apple':", stringSet.Contains("apple"))   // 输出 "Set contains 'apple': true"
	fmt.Println("Set contains 'orange':", stringSet.Contains("orange")) // 输出 "Set contains 'orange': true"
	fmt.Println("Set contains 'grape':", stringSet.Contains("grape"))   // 输出 "Set contains 'grape': false"

	// Contains
	fmt.Println(stringSet.Contains("apple"))
	fmt.Println(stringSet.Contains("applePen"))

}
