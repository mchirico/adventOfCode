package main

import (
	"aoc20221201/collect"
	"fmt"
)

func main() {
	data, err := collect.Package()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
