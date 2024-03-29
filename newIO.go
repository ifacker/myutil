package myutil

import (
	"bufio"
	"errors"
	"io"
)

// 新的 readAll，用于替换以前的 ioutil.ReadAll() 函数，解决 ioutil.ReadAll() 函数在遇到超大 reader 流时的效率低下以及报错问题
func NewReadAll(reader io.Reader) ([]byte, error) {
	if reader == nil {
		return nil, errors.New("reader 为空")
	}
	bufReader := bufio.NewReader(reader)
	var result []byte
	var buf [1024]byte
	var i = 0
	for {
		i++
		//fmt.Println(i)
		n, err := bufReader.Read(buf[:])
		if err == io.EOF {
			//fmt.Println("read the file finished")
			break
		}
		if err != nil {
			return result, err
		}
		result = append(result, buf[:n]...)
	}
	return result, nil
}
