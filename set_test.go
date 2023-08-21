package myutil

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	mySet := make(Set)
	mySet.Add("apple")
	mySet.Add("banana")
	mySet.Add("apple") // 这个不会重复添加

	fmt.Println(mySet.Contains("apple"))  // 输出 true
	fmt.Println(mySet.Contains("orange")) // 输出 false
}
