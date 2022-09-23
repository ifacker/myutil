package myutil

import (
	"fmt"
	"testing"
)

func TestUnicode2String(t *testing.T) {
	fmt.Println("\uf8ff\u0046")
	unicode := `\u60a8\u4f3c\u4e4e\u5df2\u7ecf\u7b7e\u5230\u8fc7\u4e86..`
	str, err := Unicode2String(unicode)
	fmt.Println(str, err)
}
