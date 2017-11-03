package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Count(s []string) map[string]int {
	var word string
	m := make(map[string]int)
	for i := 0; i < len(s); {
		word = s[i]
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

//数字为key，次数为value

func main() {
	file, err := os.Open("millionRandom.txt")
	if err != nil {
		fmt.Println("failed to open")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	data := []string{}

	for {
		str, err := reader.ReadString('\n') //每次读取一行
		if err != nil {
			break // 读完或发生错误
		}
		str = strings.Replace(str, "\n", "", -1) //去掉每一行的换行符
		str = strings.Replace(str, "\r", "", -1)
		data = append(data, str) //我试图用切片保存
	}

	outputFile, outputError := os.OpenFile("qiurungeng.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	fmt.Fprintln(outputFile, Count(data))
	defer outputFile.Close()
}
