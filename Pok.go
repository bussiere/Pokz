package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")
	var buffer bytes.Buffer
	color := []string{"H", "S", "D", "C"}
	value := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}
	for i := 0; i < len(color); i++ {
		//fmt.Println(color[i])
		for j := 0; j < len(value); j++ {
			buffer.WriteString("\"" + value[j] + color[i] + "\",")
		}

	}
	fmt.Println(buffer.String())
}
