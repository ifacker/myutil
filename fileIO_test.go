package myutil

import (
	"fmt"
	"testing"
)

func TestReadTxt(t *testing.T) {
	filepath := "test.txt"
	result, err := ReadTxt(filepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func TestReadTxt2Byte(t *testing.T) {
	filepath := "../config.cfg"
	body, err := ReadTxt2Byte(filepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

func TestWriteCsv(t *testing.T) {
	filePath := "test.csv"
	title := []string{"ip", "port", "cve", "status"}
	WriteCsv(filePath, title)
}

func TestWriteCsv2(t *testing.T) {
	filePath := "test.csv"
	title := []string{"ip", "port", "cve", "status"}
	data := [][]string{{"192.168.1.1", "8080", "cve-2021-22205", "false"}, {"192.168.31.33", "8080", "cve-2021-22205", "true"}}
	WriteCsv2(filePath, title, data)
}

func TestName(t *testing.T) {
	a := "11"
	b := ""
	b += a
	fmt.Println(b)
}

func TestIsFile(t *testing.T) {
	filepath := "../Poc"
	fmt.Println(IsFile(filepath))
}

func TestWriteTxt(t *testing.T) {
	filepath := "test.txt"
	data := "hello world \n fuck\nfff\n"
	err := WriteTxt(filepath, data)
	if err != nil {
		fmt.Println(err)
	}
}

func TestWriteTxt2(t *testing.T) {
	filepath := "test.txt"
	data := "hello world \n fuck\nfff\n"
	err := WriteTxt2(filepath, data)
	if err != nil {
		fmt.Println(err)
	}
}
