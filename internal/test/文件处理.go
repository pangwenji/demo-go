package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 创建文件
func creteFile() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("error creating", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString("hello this is  go")

	if err != nil {
		fmt.Println("error ")
	} else {
		fmt.Println("wrting is succcess")
	}
}

// 读取文件
func readyFile() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("happen error")
	}
}

// 判断文件是否存在
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// 复制文件

func copyFile() {
	srcFile, err := os.Open("source.txt")
	if err != nil {
		fmt.Println("Error opening source")
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create("destination.txt")
	if err != nil {
		fmt.Println("destination.txt Error")
		return
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("复制错误")
	} else {
		fmt.Println("复制成功")
	}
}

func main() {
	//creteFile()
	//readyFile()

	//filename := "example.txt"
	//if fileExists(filename) {
	//	fmt.Println("File exists:", filename)
	//} else {
	//	fmt.Println("File does not exist:", filename)
	//}
	copyFile()
}
