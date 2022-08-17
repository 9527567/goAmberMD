package main

import (
	"fmt"
	"go_amber/service"
)

func main() {
	var input service.Input
	var inputS []string
	inputS = append(inputS, "Mini")
	fmt.Println(input.CreateMinInput(inputS))

}
