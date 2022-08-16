package service

import (
	"encoding/json"
	"fmt"
	"go_amber/config"
	"strings"
)

// 已知的残基名称
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

// 输入的所有参数
type Input struct {
	imin          bool    // 是否执行最小化任务
	ntmin         int     // 最小化任务的标志
	maxcyc        int     // 最小化的最大循环次数
	ncyc          int     // 在ncyc后将最小化方法从最陡下降法切换为共轭梯度法
	restraintmask string  // 指定约束原子的标记
	restraint_wt  float64 // 约束原子的力
	irest         bool    // 是否重启模拟
	nstlim        int     // 要执行的MD步数
	ntb           int     // 是否执行周期性边界
	ntc           int     // SHAKE 执行键长约束的标志
	cut           float64 // 指定非键截断值，8.0通常是一个不错的选择
	tempi         float64 // 初始温度
	tautp         float64 // 时间常数
	taup          float64 // 压力松弛时间
	mcbarint      int     // 作为蒙特卡洛恒压器的一部分执行的体积更改尝试之间的步数 默认值为 100
	gamma_ln      float64 // 以ps为单位的碰撞频率
	dt            int     // 时间步长
	nscm          int     // pbc处理选项，对于周期性系统，平移被修正，旋转不会
	ntwx          int     // 每 ntwx 步，坐标将被写入 netcdf 文件 如果ntwx = 0，没有坐标 轨迹文件将被写入 默认 = 0
	ntpr          int     // 每 ntpr 步,能量信息将以人类可读的形式打印到文件mdout和mdinfo.mdinfo每次关闭又重新打开，所以它总是包含最新的能量和温度 默认 50
	ntwr          int     // 每ntwr步，写入重启文件的运动快照
	//previousref   int
	//heavyrst      int
	//bbrst         int
	thermo   string //
	barostat int    // 用于控制使用哪个恒压器来控制压力的标志。
	// flexiblewat int
}

// 如何优雅的创建输入文件？
func CreateMinInput(task string, ntmin int, restraintmask string, restrain_wt float64) []string {

	var result []string
	//result = append(result, "Minimization: "+task+"\n")
	//result = append(result, " &cntrl\n")
	//result = append(result, "imin = 0,ig = -1,ntwv = -1, ioutfm = 1, ntxo = 2, iwrap = 0,")
	//
	//result = append(result, "\n&end")

	return result

}
