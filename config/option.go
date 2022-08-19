package config

import (
	"github.com/jessevdk/go-flags"
)

type Option struct {
	Parm7       string  `short:"p" long:"parm7" description:"amber parm7 file" required:"true"`
	Rst7        string  `short:"c" long:"rst7" description:"amber srt7 file" required:"true"`
	Temp        float64 `short:"t" long:"temp" description:"Simulated Temperature" required:"true"`
	Thermo      string  `long:"thermo" description:"Thermostat: berendsen, langevin" default:"langevin" choice:"berendsen" choice:"langevin"`
	Baro        string  `long:"baro" description:"Barostat: berendsen, montecarlo " default:"montecarlo" choice:"berendsen" choice:"montecarlo"`
	Finalthermo string  `long:"finalthermo" description:"Barostat: berendsen, langevin " default:"langevin" choice:"berendsen" choice:"langevin"`
	Finalbaro   string  `long:"finalbaro" description:"Barostat: berendsen, montecarlo " default:"montecarlo" choice:"berendsen" choice:"montecarlo"`
	Cutoff      float64 `long:"cutoff" description:"If specified, override default cutoffs with cut" default:"8.0"`
	Ares        string  `long:"ares" description:"Residue name to add to heavy atom masks if present"`
}

func ParseOption() Option {
	var option Option
	flags.Parse(&option)
	return option
}
