package service

import (
	"encoding/json"
	"fmt"
	"strings"
)

var NtpFlags = 1
var CHARMMWATERFLAG = false

// Name 已知的残基名称
type Name struct {
	Protein     []string `json:"Protein"`
	DNA         []string `json:"DNA"`
	RNA         []string `json:"RNA"`
	Lipid       []string `json:"Lipid"`
	Carbo       []string `json:"Carbo"`
	CharmmWater []string `json:"CharmmWater"`
	Water       []string `json:"Water"`
}

// 从字符串中返回结构体
func getName() Name {
	knowName := `{
  "Protein": [
    "ACE",
    "ALA",
    "ARG",
    "ASH",
    "AS4",
    "ASN",
    "ASP",
    "CALA",
    "CARG",
    "CASN",
    "CASP",
    "CCYS",
    "CCYX",
    "CGLN",
    "CGLU",
    "CGLY",
    "CHID",
    "CHIE",
    "CHIP",
    "CHIS",
    "CHYP",
    "CILE",
    "CLEU",
    "CLYS",
    "CMET",
    "CPHE",
    "CPRO",
    "CSER",
    "CTHR",
    "CTRP",
    "CTYR",
    "CVAL",
    "CYM",
    "CYS",
    "CYX",
    "GLH",
    "GL4",
    "GLN",
    "GLU",
    "GLY",
    "HID",
    "HIE",
    "HIP",
    "HIS",
    "HYP",
    "ILE",
    "LEU",
    "LYN",
    "LYS",
    "MET",
    "NALA",
    "NARG",
    "NASN",
    "NASP",
    "NCYS",
    "NCYX",
    "NGLN",
    "NGLU",
    "NGLY",
    "NHE",
    "NHID",
    "NHIE",
    "NHIP",
    "NHIS",
    "NILE",
    "NLEU",
    "NLYS",
    "NME",
    "NMET",
    "NPHE",
    "NPRO",
    "NSER",
    "NTHR",
    "NTRP",
    "NTYR",
    "NVAL",
    "PHE",
    "PRO",
    "SER",
    "THR",
    "TRP",
    "TYR",
    "VAL"
  ],
  "DNA": [
    "DA",
    "DA3",
    "DA5",
    "DAN",
    "DC",
    "DC3",
    "DC5",
    "DCN",
    "DG",
    "DG3",
    "DG5",
    "DGN",
    "DT",
    "DT3",
    "DT5",
    "DTN"
  ],
  "RNA": [
    "A",
    "A3",
    "A5",
    "AMP",
    "AN",
    "C",
    "C3",
    "C5",
    "CMP",
    "CN",
    "G",
    "G3",
    "G5",
    "GMP",
    "GN",
    "OHE",
    "U",
    "U3",
    "U5",
    "UMP",
    "UN"
  ],
  "Lipid": [
    "POPE",
    "DOPC",
    "AR",
    "CHL",
    "DHA",
    "LAL",
    "MY",
    "OL",
    "PA",
    "PC",
    "PE",
    "PGR",
    "PH-",
    "PS",
    "ST"
  ],
  "Carbo": [
    "0GB",
    "4GB",
    "0YA",
    "4YA",
    "0fA",
    "0YB",
    "2MA",
    "4YB",
    "NLN",
    "UYB",
    "VMB",
    "0SA",
    "6LB",
    "ROH"
  ],
  "CharmmWater": [
    "TIP3"
  ],
	"Water":[
	"WAT"	
  ]
}`
	var name Name
	err := json.Unmarshal([]byte(knowName), &name)
	if err != nil {
		return Name{}
	}
	return name
}

func SystemNumbers(SystemRes []string) (int, int, int, int, int, int, int) {
	numProtein := 0
	numDNA := 0
	numRNA := 0
	numLipid := 0
	numCarbo := 0
	nCharmmWater := 0
	nWater := 0
	//N
	name := getName()
	for _, s := range SystemRes {
		if in(s, name.Protein, false) {
			numProtein++
		} else if in(s, name.DNA, false) {
			numDNA++
		} else if in(s, name.RNA, false) {
			numRNA++
		} else if in(s, name.Lipid, false) {
			numLipid++
		} else if in(s, name.Carbo, false) {
			numCarbo++
		} else if in(s, name.CharmmWater, false) {
			nCharmmWater++
		} else if in(s, name.Water, false) {
			nWater++
		}
	}
	if numLipid > 0 {
		NtpFlags = 2
	}
	if nCharmmWater > 0 {
		CHARMMWATERFLAG = true
	}
	return numProtein, numDNA, numRNA, numLipid, numCarbo, nCharmmWater, nWater
}

// Input 输入的所有参数
type Input struct {
	Name          string // 任务名称
	Imin          string // 是否执行最小化任务
	Ntmin         string // 最小化任务的标志
	Maxcyc        string // 最小化的最大循环次数
	Ncyc          string // 在ncyc后将最小化方法从最陡下降法切换为共轭梯度法
	Restraintmask string // 指定约束原子的标记
	Restraint_wt  string // 约束原子的力
	Irest         string // 是否重启模拟
	Nstlim        string // 要执行的MD步数
	Ntb           string // 是否执行周期性边界
	Ntc           string // SHAKE 执行键长约束的标志
	Cut           string // 指定非键截断值，8.0通常是一个不错的选择
	Tempi         string // 初始温度
	Tautp         string // 时间常数
	Taup          string // 压力松弛时间
	Mcbarint      string // 作为蒙特卡洛恒压器的一部分执行的体积更改尝试之间的步数 默认值为 100
	Gamma_ln      string // 以ps为单位的碰撞频率
	Dt            string // 时间步长
	Nscm          string // pbc处理选项，对于周期性系统，平移被修正，旋转不会
	Ntwx          string // 每 ntwx 步，坐标将被写入 netcdf 文件 如果ntwx = 0，没有坐标 轨迹文件将被写入 默认 = 0
	Ntpr          string // 每 ntpr 步,能量信息将以人类可读的形式打印到文件mdout和mdinfo.mdinfo每次关闭又重新打开，所以它总是包含最新的能量和温度 默认 50
	Ntwr          string // 每ntwr步，写入重启文件的运动快照
	Ntr           string
	Ntxo          string
	Thermo        string //
	Barostat      string // 用于控制使用哪个恒压器来控制压力的标志。
	Ncsm          string
	Ntf           string
	Ntx           string
	Temp0         string

	//previousref   int
	//heavyrst      int
	//bbrst         int

	// flexiblewat int
}

// CreateMinInput 如何优雅的创建输入文件？回调函数吗？或者是json
func (input Input) CreateMinInput(option ...string) []string {
	input.Imin = "1"
	input.Ntmin = "2"
	input.Maxcyc = "1000"
	input.Ncyc = "10"
	input.Ntwx = "500"
	input.Ntpr = "50"
	input.Ntwr = "500"
	input.Cut = "8.0"
	input.Ntxo = "2"
	input.Restraintmask = ""
	input.Restraint_wt = "0"
	input.Dt = toString(0.001)
	for i := 0; i < len(option)-1; i += 2 {
		if strings.EqualFold(option[i], "name") {
			input.Name = option[i+1]
		}
		if strings.EqualFold(option[i], "imin") {
			input.Imin = option[i+1]
		}
		if strings.EqualFold(option[i], "ntmin") {
			input.Ntmin = option[i+1]
		}
		if strings.EqualFold(option[i], "maxcyc") {
			fmt.Println(input.Maxcyc)
			input.Maxcyc = option[i+1]
		}
		if strings.EqualFold(option[i], "ncyc") {
			input.Ncyc = option[i+1]
		}
		if strings.EqualFold(option[i], "irest") {
			input.Irest = option[i+1]
		}
		if strings.EqualFold(option[i], "nstlim") {
			input.Nstlim = option[i+1]
		}
		if strings.EqualFold(option[i], "ntb") {
			input.Ntb = option[i+1]
		}
		if strings.EqualFold(option[i], "ntc") {
			input.Ntc = option[i+1]
		}
		if strings.EqualFold(option[i], "cut") {
			input.Cut = option[i+1]
		}
		if strings.EqualFold(option[i], "tempi") {
			input.Tempi = option[i+1]
		}
		if strings.EqualFold(option[i], "taup") {
			input.Tautp = option[i+1]
		}
		if strings.EqualFold(option[i], "mcbarnt") {
			input.Mcbarint = option[i+1]
		}
		if strings.EqualFold(option[i], "gamma_ln") {
			input.Gamma_ln = option[i+1]
		}
		if strings.EqualFold(option[i], "dt") {
			input.Dt = option[i+1]
		}
		if strings.EqualFold(option[i], "nscm") {
			input.Nscm = option[i+1]
		}
		if strings.EqualFold(option[i], "ntwx") {
			input.Ntwx = option[i+1]
		}
		if strings.EqualFold(option[i], "ntwr") {
			input.Ntwr = option[i+1]
		}
		if strings.EqualFold(option[i], "thermo") {
			input.Thermo = option[i+1]
		}
		if strings.EqualFold(option[i], "barostat") {
			input.Barostat = option[i+1]
		}
		if strings.EqualFold(option[i], "restraintmask") {
			input.Ntr = "1"
			input.Restraintmask = option[i+1]
		}
		if strings.EqualFold(option[i], "restraint_wt") {
			input.Restraint_wt = option[i+1]
		}
	}
	var result []string
	result = append(result, "Minimization: "+input.Name+"\n")
	result = append(result, " &cntrl\n")
	result = append(result, "imin = 1, ioutfm = 1, ntxo = 2,ntc = 1, ntf = 1, ntb = 1,"+"\n")
	result = append(result, "maxcyc = "+input.Maxcyc+",")
	result = append(result, "dt = "+input.Dt+",")
	result = append(result, "ntxo = "+input.Ntxo+",")
	result = append(result, "ntwx = "+input.Ntwx+",")
	result = append(result, "ntpr = "+input.Ntpr+",")
	result = append(result, "ntwr = "+input.Ntwr+",")
	result = append(result, "cut = "+input.Cut+",")
	result = append(result, "\n")
	result = append(result, "ntr = "+input.Ntr+",")
	result = append(result, "restraintmask = "+input.Restraintmask+",")
	result = append(result, "restraint_wt = "+input.Restraint_wt+",")
	result = append(result, "\n")
	result = append(result, "&end\n")
	return result

}
func (input Input) Restraint(option ...string) []string {
	for i := 0; i < len(option)-1; i += 2 {
		if strings.EqualFold(option[i], "restraintmask") {
			input.Ntr = "1"
			input.Restraintmask = option[i+1]
		}
		if strings.EqualFold(option[i], "restraint_wt") {
			input.Restraint_wt = option[i+1]
		}
	}
	var result []string
	result = append(result, "ntr = "+input.Ntr+",")
	result = append(result, "restraintmask = "+input.Restraintmask+",")
	result = append(result, "restraint_wt = "+input.Restraint_wt+",")
	result = append(result, "\n")
	return result
}

func (input Input) CreateMdInput(option ...string) []string {
	input.Imin = "0"
	input.Ntmin = "2"
	input.Nstlim = "5000"
	input.Irest = "0"
	input.Ntb = "1"
	input.Tautp = "1.0"
	input.Taup = "1.0"
	input.Mcbarint = "100"
	input.Gamma_ln = "5"
	input.Dt = "0.001"
	input.Ncsm = "0"
	input.Ntwx = "500"
	input.Ntpr = "50"
	input.Ntwr = "500"
	input.Tempi = "303.15"
	input.Ntc = "2"
	input.Ntf = "2"
	input.Cut = "8.0"
	input.Dt = toString(0.001)
	for i := 0; i < len(option)-1; i += 2 {
		if strings.EqualFold(option[i], "name") {
			input.Name = option[i+1]
		}
		if strings.EqualFold(option[i], "imin") {
			input.Imin = option[i+1]
		}
		if strings.EqualFold(option[i], "ntmin") {
			input.Ntmin = option[i+1]
		}
		if strings.EqualFold(option[i], "maxcyc") {
			fmt.Println(input.Maxcyc)
			input.Maxcyc = option[i+1]
		}
		if strings.EqualFold(option[i], "ncyc") {
			input.Ncyc = option[i+1]
		}
		if strings.EqualFold(option[i], "irest") {
			input.Irest = option[i+1]
		}
		if strings.EqualFold(option[i], "nstlim") {
			input.Nstlim = option[i+1]
		}
		if strings.EqualFold(option[i], "ntb") {
			input.Ntb = option[i+1]
		}
		if strings.EqualFold(option[i], "ntc") {
			input.Ntc = option[i+1]
		}
		if strings.EqualFold(option[i], "cut") {
			input.Cut = option[i+1]
		}
		if strings.EqualFold(option[i], "tempi") {
			input.Tempi = option[i+1]
			input.Temp0 = input.Tempi
		}
		if strings.EqualFold(option[i], "taup") {
			input.Tautp = option[i+1]
		}
		if strings.EqualFold(option[i], "mcbarnt") {
			input.Mcbarint = option[i+1]
		}
		if strings.EqualFold(option[i], "gamma_ln") {
			input.Gamma_ln = option[i+1]
		}
		if strings.EqualFold(option[i], "dt") {
			input.Dt = option[i+1]
		}
		if strings.EqualFold(option[i], "nscm") {
			input.Nscm = option[i+1]
		}
		if strings.EqualFold(option[i], "ntwx") {
			input.Ntwx = option[i+1]
		}
		if strings.EqualFold(option[i], "ntwr") {
			input.Ntwr = option[i+1]
		}
		if strings.EqualFold(option[i], "thermo") {
			input.Thermo = option[i+1]
		}
		if strings.EqualFold(option[i], "barostat") {
			input.Barostat = option[i+1]
		}
		if strings.EqualFold(option[i], "ntc") {
			input.Ntc = option[i+1]
			input.Ntf = option[i+1]
		}
		if strings.EqualFold(option[i], "restraintmask") {
			input.Ntr = "1"
			input.Restraintmask = option[i+1]
		}
		if strings.EqualFold(option[i], "restraint_wt") {
			input.Restraint_wt = option[i+1]
		}

	}
	if input.Irest != "0" {
		input.Ntx = "1"
	} else {
		input.Ntx = "5"
	}
	var result []string
	result = append(result, "Minimization: "+input.Name+"\n")
	result = append(result, " &cntrl\n")
	result = append(result, "imin = 0,ig = -1,ntwv = -1, ioutfm = 1, ntxo = 2, iwrap = 0,"+"\n")
	result = append(result, "ntmin = "+input.Ntmin+",")
	result = append(result, "nstlim = "+input.Nstlim+",")
	result = append(result, "ntx = "+input.Ntx+",")
	result = append(result, "irest = "+input.Irest+",")
	result = append(result, "ntwx = "+input.Ntwx+",")
	result = append(result, "ntpr = "+input.Ntpr+",")
	result = append(result, "ntwr = "+input.Ntwr+",")
	result = append(result, "nscm = "+input.Nscm+",")
	result = append(result, "ntc = "+input.Ntc+",")
	result = append(result, "ntf = "+input.Ntf+",")
	result = append(result, "ntb = "+input.Ntb+",")
	result = append(result, "cut = "+input.Cut+",")
	result = append(result, "\n")
	result = append(result, "ntr = "+input.Ntr+",")
	result = append(result, "restraintmask = "+input.Restraintmask+",")
	result = append(result, "restraint_wt = "+input.Restraint_wt+",")
	result = append(result, "\n")
	func() {
		if strings.EqualFold(input.Thermo, "berendsen") {
			result = append(result, "ntt = 1,")
			result = append(result, "tautp = "+input.Tautp+",")
			result = append(result, "temp0"+input.Temp0+",")
			result = append(result, "tempi"+input.Tempi+",")
			result = append(result, "\n")
		} else if strings.EqualFold(input.Thermo, "langevin") {
			result = append(result, "ntt = 3,")
			result = append(result, "gamma_ln = "+input.Gamma_ln+",")
			result = append(result, "temp0"+input.Temp0+",")
			result = append(result, "tempi"+input.Tempi+",")
			result = append(result, "\n")
		}
	}()
	if input.Ntb == "2" {
		func() {
			if input.Barostat == "berendsen" {
				result = append(result, "ntp = "+toString(NtpFlags)+",")
				result = append(result, "taup = "+input.Taup+",")
				result = append(result, "pes0 = 1.0")
				result = append(result, "\n")
			} else if input.Barostat == "montecarlo" {
				result = append(result, "ntp = "+toString(NtpFlags)+",")
				result = append(result, "barostat = 2 ,")
				result = append(result, "pes0 = 1.0")
				result = append(result, "mcbarint = "+input.Mcbarint+",")
				result = append(result, "\n")
			}
		}()
	}
	result = append(result, "&end\n")
	return result
}

func CharmmWater() []string {
	var result []string
	if CHARMMWATERFLAG {
		result = append(result, "  WATNAM = 'TIP3', OWTNM = 'OH2',")
		result = append(result, "\n")
	}
	return result
}
