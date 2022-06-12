package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	c, err := StartSampleContainer()
	fmt.Println("Container started")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(c)
}
