package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_amber/config"
	"os/exec"
	"sort"
	"strings"
)

func RunCmd(cmdStr string) string {
	cmd := exec.Command("bash", "-c", cmdStr)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.String()
	} else {
		return out.String()
	}
}
func in(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

type Name struct {
	Protein []string `json:"Protein"`
	DNA     []string `json:"DNA"`
	RNA     []string `json:"RNA"`
	Lipid   []string `json:"Lipid"`
	Carbo   []string `json:"Carbo"`
	Solvent []string `json:"Solvent"`
}

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

func SystemNumbers(SymtemRes []string) (int, int, int, int, int, int) {
	numProtein := 0
	numDNA := 0
	numRNA := 0
	numLipid := 0
	numCarbo := 0
	numSolvent := 0
	//N
	name := getName()
	for _, s := range SymtemRes {
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
	var result string = RunCmd(tmp)
	s := strings.Split(result, "\n")
	var Res []string
	for i := 1; i < len(s); i++ {
		if len(s[i]) == 47 {
			Res = append(
				Res,
				strings.Trim(s[i][5:10], " "),
			)
		}
	}
	numProtein, numDNA, numRNA, numLipid, numCarbo, numSolvent := SystemNumbers(Res)
	fmt.Println(numProtein, numDNA, numRNA, numLipid, numCarbo, numSolvent)
}
