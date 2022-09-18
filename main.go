package main

import (
	"fmt"
	"go_amber/config"
	"go_amber/service"
)

func StandardEq(option config.Option) {
	systemInfo := service.GetSystemInfo(&option)
	fmt.Println(service.CreateRestraintMask(systemInfo, &option))
}

func main() {
	option := config.ParseOption()
	if !option.Norestart {
		StandardEq(option)
	}
}
