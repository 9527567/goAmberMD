package config

import (
	"github.com/jessevdk/go-flags"
)

type Option struct {
	Parm7 string  `short:"p" long:"parm7" description:"amber parm7 file" required:"true"`
	Rst7  string  `short:"c" long:"rst7" description:"amber srt7 file" required:"true"`
	Temp  float64 `short:"t" long:"temp" description:"Simulated Temperature" required:"true"`
}

func ParseOption() Option {
	var option Option
	flags.Parse(&option)
	return option
}
