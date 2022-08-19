package main

import (
	"fmt"
	"go_amber/config"
	"go_amber/service"
)

func main() {
	//var input service.Input
	//fmt.Println(input.CreateMinInput("Name", "step1", "Imin", "1", "ntmin", "2"))
	option := config.ParseOption()
	systemInfo := service.GetSystemInfo(&option)
	fmt.Println(service.CreateRestraintMask(systemInfo, &option))
}
