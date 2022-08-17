package service

import (
	"encoding/json"
	"fmt"
	"go_amber/config"
	"strings"
)

// Name 已知的残基名称
type Name struct {
	Protein []string `json:"Protein"`
	DNA     []string `json:"DNA"`
	RNA     []string `json:"RNA"`
	Lipid   []string `json:"Lipid"`
	Carbo   []string `json:"Carbo"`
	Solvent []string `json:"Solvent"`
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
  "Solvent": [
    "TIP3",
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

func SystemNumbers(SystemRes []string) (int, int, int, int, int, int) {
	numProtein := 0
	numDNA := 0
	numRNA := 0
	numLipid := 0
	numCarbo := 0
	numSolvent := 0
	//N
	name := getName()
	for _, s := range SystemRes {
		if in(s, name.Protein) {
			numProtein++
		} else if in(s, name.DNA) {
			numDNA++
		} else if in(s, name.RNA) {
			numRNA++
		} else if in(s, name.Lipid) {
			numLipid++
		} else if in(s, name.Carbo) {
			numCarbo++
		} else if in(s, name.Solvent) {
			numSolvent++
		}
	}
	return numProtein, numDNA, numRNA, numLipid, numCarbo, numSolvent
}

func Temp() {
	opt := config.ParseOption()
	tmp := "cpptraj " + "-p " + opt.Parm7 + " --resmask " + " \\* "
	result := RunCmd(tmp)
	s := strings.Split(result, "\n")
	var Res []string
	for i := 1; i < len(s); i++ {
		if len(s[i]) == 47 {
			// Res = append(Res, strings.Trim(s[i][6:10], " "))
			Res = append(Res, strings.Fields(s[i])[1])
		}
	}
	numProtein, numDNA, numRNA, numLipid, numCarbo, numSolvent := SystemNumbers(Res)
	fmt.Println(numProtein, numDNA, numRNA, numLipid, numCarbo, numSolvent)
}

// Input 输入的所有参数
type Input struct {
	Name          string  // 任务名称
	Imin          int8    // 是否执行最小化任务
	Ntmin         int     // 最小化任务的标志
	Maxcyc        int     // 最小化的最大循环次数
	Ncyc          int     // 在ncyc后将最小化方法从最陡下降法切换为共轭梯度法
	Restraintmask string  // 指定约束原子的标记
	Restraint_wt  float64 // 约束原子的力
	Irest         bool    // 是否重启模拟
	Nstlim        int     // 要执行的MD步数
	Ntb           int     // 是否执行周期性边界
	Ntc           int     // SHAKE 执行键长约束的标志
	Cut           float64 // 指定非键截断值，8.0通常是一个不错的选择
	Tempi         float64 // 初始温度
	Tautp         float64 // 时间常数
	Taup          float64 // 压力松弛时间
	Mcbarint      int     // 作为蒙特卡洛恒压器的一部分执行的体积更改尝试之间的步数 默认值为 100
	Gamma_ln      float64 // 以ps为单位的碰撞频率
	Dt            float64 // 时间步长
	Nscm          int     // pbc处理选项，对于周期性系统，平移被修正，旋转不会
	Ntwx          int     // 每 ntwx 步，坐标将被写入 netcdf 文件 如果ntwx = 0，没有坐标 轨迹文件将被写入 默认 = 0
	Ntpr          int     // 每 ntpr 步,能量信息将以人类可读的形式打印到文件mdout和mdinfo.mdinfo每次关闭又重新打开，所以它总是包含最新的能量和温度 默认 50
	Ntwr          int     // 每ntwr步，写入重启文件的运动快照
	//previousref   int
	//heavyrst      int
	//bbrst         int
	Thermo   string //
	Barostat int    // 用于控制使用哪个恒压器来控制压力的标志。
	// flexiblewat int
}

// CreateMinInput 如何优雅的创建输入文件？回调函数吗？或者是json
func (input Input) CreateMinInput(i []string) []string {
	input.Name = i[0]
	input.Imin = 1
	input.Ntmin = 2
	input.Maxcyc = 1000
	input.Ncyc = 10
	input.Ntwx = 500
	input.Ntpr = 50
	input.Ntwr = 500
	input.Restraintmask = ""
	input.Restraint_wt = 0
	input.Dt = 0.001

	var result []string
	result = append(result, "Minimization: "+input.Name+"\n")
	result = append(result, " &cntrl\n")
	result = append(result, "imin = 0,ig = -1,ntwv = -1, ioutfm = 1, ntxo = 2, iwrap = 0,")
	result = append(result, "\n&end")
	return result

}

//func CreateInput(input Input) {
//	t, v := ForeachStruct(input)
//	for i := 0; i < t.NumField(); i++ {
//		fmt.Println(t.Field(i).Name, v.Field(i).Interface())
//	}
//}
