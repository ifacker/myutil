package io

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestNewReadAll(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	buf := bufio.NewReader(file)
	result, err := NewReadAll(buf)
	fmt.Println(string(result), err)
}
