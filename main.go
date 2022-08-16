package main

import (
	"fmt"
	"go_amber/service"
)

func main() {
	input := service.CreateMinInput("step1", 1, "1:2080", 8.0)
	fmt.Println(input)
}
