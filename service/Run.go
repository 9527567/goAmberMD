package service

import (
	"go_amber/config"
	"strings"
)

func GetSystemInfo(option *config.Option) []int {
	opt := option.Parm7
	tmp := "cpptraj " + "-p " + opt + " --resmask " + " \\* "
	result := RunCmd(tmp)
	s := strings.Split(result, "\n")
	var Res []string
	for i := 1; i < len(s); i++ {
		if len(s[i]) >= 47 {
			//Res = append(Res, strings.Trim(s[i][6:10], " "))
			Res = append(Res, strings.Fields(s[i])[1])
		}

	}
	numProtein, numDNA, numRNA, numLipid, numCarbo, nCharmmWater, nWater := SystemNumbers(Res)
	var Number []int
	Number = append(Number, numProtein, numDNA, numRNA, numLipid, numCarbo, numCarbo, nCharmmWater, nWater)
	return Number
}
func CreateRestraintMask(n []int, option *config.Option) string {
	var result string
	S := n[0] + n[1] + n[2] + n[3] + n[4]
	if S > 0 {
		result = result + ":1:" + toString(S) + "&!@H="
	}
	if len(option.Ares) != 0 {
		result = result + "|"
		_ares := strings.Fields(option.Ares)
		for _, mask := range _ares {
			result = result + ":" + mask + "&!@H="
		}
	}
	return result
}
