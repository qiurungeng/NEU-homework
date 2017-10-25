// 作业作业 project main.go
package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)

	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}

	return p, err
}

func sampleReadFile() {

	file, _ := os.Open("millionRandom.txt")
	defer file.Close()
	data, _ := ReadFrom(file, 1000000000)
	fmt.Println(string(data))
}
func Count(s string) map[string]int {
	var word string
	m := make(map[string]int)
	for i := 0; i < len(s); {
		word = s[i : i+1]
		v, ok := m[word]
		if ok != false {
			m[word] = v + 1
		} else {
			m[word] = 1
		}
		i += 1
	}

	return m
}
func main() {
	file, _ := os.Open("millionRandom.txt")
	defer file.Close()
	data, _ := ReadFrom(file, 1000000000)
	var Allnum string = string(data)
	//fmt.Println(Allnum)  for text
	//fmt.Println(Count(Allnum))

	outputFile, outputError := os.OpenFile("qiurungeng.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	fmt.Fprintln(outputFile, Count(Allnum))
	defer outputFile.Close()
}
