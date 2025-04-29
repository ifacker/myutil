package myutil

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"

	"github.com/gookit/color"
)

// 读取文件
func ReadTxt(filepath string) (string, error) {
	if !IsFile(filepath) {
		return "", errors.New(color.Warn.Sprintf("注意：将要读取的文件不存在"))
	}
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return "", err
	}
	data := bufio.NewScanner(file)
	result := ""
	for data.Scan() {
		result += data.Text() + "\n"
	}
	return result, nil
}

// 读取文件(无损读取，不会更改文件内的任意字符)
func ReadTxt2Byte(filepath string) ([]byte, error) {
	if !IsFile(filepath) {
		return nil, errors.New(color.Warn.Sprintf("注意：将要读取的文件不存在"))
	}
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	result, err1 := io.ReadAll(file)
	if err1 != nil {
		return nil, err1
	}

	return result, nil
}

// 写入csv文件（追加模式）注：追加模式无标题
func WriteCsv(filepath string, data []string) error {
	if IsFile(filepath) {
		// 如果文件存在，则直接打开并写入
		file2, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			return err
		}
		defer file2.Close()
		file2.WriteString("\xEF\xBB\xBF")
		// 写入
		writer := csv.NewWriter(file2)
		writer.Write(data)
		writer.Flush()
	} else {
		file, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer file.Close()
		file.WriteString("\xEF\xBB\xBF")
		// 写入
		writer := csv.NewWriter(file)
		writer.Write(data)
		writer.Flush()
	}
	return nil
}

// 写入csv文件（覆盖模式）
func WriteCsv2(filePath string, title []string, data [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)

	// 写入
	writer.Write(title)
	for _, dataLine := range data {
		writer.Write(dataLine)
	}
	writer.Flush()

	return nil
}

// 判断文件或文件夹是否存在
func IsFile(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 写入txt文件(追加模式)
func WriteTxt(filepath, data string) error {
	if IsFile(filepath) {
		// 如果文件存在，则直接打开并写入
		color.Warn.Println("注意：文件已存在，该操作将会把内容追加至文件末尾处")
		file2, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR, 0666)
		defer file2.Close()
		if err != nil {
			return err
		}
		_, err1 := file2.WriteString(data)
		if err1 != nil {
			return err1
		}
	} else {
		file, err := os.Create(filepath)
		defer file.Close()
		if err != nil {
			return err
		}
		_, err1 := file.WriteString(data)
		if err1 != nil {
			return err1
		}
	}
	return nil
}

// 写入txt文件(覆盖模式)
func WriteTxt2(filepath, data string) error {
	if IsFile(filepath) {
		color.Error.Println("警告：文件已存在，该操作将会直接覆盖原有文件内的所有内容")
	}
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err1 := file.WriteString(data)
	if err1 != nil {
		return err1
	}
	return nil
}
