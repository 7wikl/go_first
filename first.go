package main

import "fmt"

func main() {
	sequence := "hello world"
	for index, value := range sequence {
		fmt.Println(index, string(value))
	}
}
