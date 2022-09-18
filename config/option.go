package config

import (
	"github.com/jessevdk/go-flags"
)

// 用于生产过程中的限制，特殊需求
type ProductionMask struct {
	Pmask string `long:"pmask" description:"Restraint mask to use during production (steps 9 and above)."`
	Pwt   string `long:"pwt" description:"Restraint weight to use for '--pmask'; required if '--pmask' specified."`
	Pref  string `long:"pref" description:" Optional reference structure to use for '--pmask'."`
}

type Optional struct {
	Thermo         string  `long:"thermo" description:"Thermostat: berendsen, langevin" default:"langevin" choice:"berendsen" choice:"langevin"`
	Baro           string  `long:"baro" description:"Barostat: berendsen, montecarlo " default:"montecarlo" choice:"berendsen" choice:"montecarlo"`
	Finalthermo    string  `long:"finalthermo" description:"Barostat: berendsen, langevin " default:"langevin" choice:"berendsen" choice:"langevin"`
	Finalbaro      string  `long:"finalbaro" description:"Barostat: berendsen, montecarlo " default:"montecarlo" choice:"berendsen" choice:"montecarlo"`
	Cutoff         float64 `long:"cutoff" description:"If specified, override default cutoffs with cut" default:"8.0"`
	Ares           string  `long:"ares" description:"Residue name to add to heavy atom masks if present"`
	ProductionMask `group:"production Restraint"`
	Overwrite      bool `short:"O" long:"overwrite" description:"Overwrite existing files, otherwise skip."`
	Charmmwater    bool `long:"charmmwater" description:"If specified assume CHARMM water (i.e. 'TIP3')."`
	Norestart      bool `long:"norestart" description:"Do standard Eq with no restarts."`
	Skipfinaleq    bool `long:"skipfinaleq" description:"If specified, skip final eq. (step 10)."`
}
type Necessary struct {
	Parm7 string  `short:"p" long:"parm7" description:"amber parm7 file" required:"true"`
	Rst7  string  `short:"c" long:"rst7" description:"amber srt7 file" required:"true"`
	Temp  float64 `short:"t" long:"temp" description:"Simulated Temperature" required:"true"`
}
type Option struct {
	Necessary `group:"necessary"`
	Optional  `group:"optional"`
}

func ParseOption() Option {
	var option Option
	flags.Parse(&option)
	return option
}
