package main

import (
	"fmt"
	"go_amber/config"
	"go_amber/service"
)

func Run(input []string, name string, option config.Option) {
	service.WriteFile(name+".in", input)
	fmt.Println(service.RunCmd("echo -O -i " + name + ".in -p " + option.Parm7 + " -c " + option.Rst7 + " -o " + name + ".out -r +" + name + ".rst7 -x " + name + ".nc -inf " + name + ".mdinfo"))
}

func StandardEq(option config.Option) {
	systemInfo := service.GetSystemInfo(&option)
	var input service.Input
	restraintmask := service.CreateRestraintMask(systemInfo, &option)
	temp := input.CreateMinInput("name", "step1", "restraintmask", restraintmask, "restraint_wt", "5.0")
	Run(temp, "step1", option)
	temp = input.CreateMdInput("name", "step2", "restraintmask", restraintmask, "restraint_wt", "5.0", "nstlim", "15000", "tautp", "0.5")
	Run(temp, "step2", option)
}

func main() {
	option := config.ParseOption()
	if !option.Norestart {
		StandardEq(option)
	}
}
