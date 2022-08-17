package service

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"sort"
	"strings"
)

// CheckFileExist 检查文件是否存在
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// RunCmd 运行系统命令
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

// 判断字符串是否在字符串数组中,二分查找
func in(target string, strArray []string, ignoreCase bool) bool {
	if ignoreCase {
		for index, s := range strArray {
			if index < len(strArray) && strings.EqualFold(s, target) {
				return true
			}
		}
	} else {
		sort.Strings(strArray)
		index := sort.SearchStrings(strArray, target)
		if index < len(strArray) && strArray[index] == target {
			return true
		}
	}
	return false
}

// ForeachStruct 遍历结构体
func ForeachStruct(obj interface{}) (reflect.Type, reflect.Value) {
	t := reflect.TypeOf(obj) // 注意，obj不能为指针类型，否则会报：panic recovered: reflect: NumField of non-struct type
	v := reflect.ValueOf(obj)
	return t, v
}

//func toString[T float64|float32|int|int8|int16|int32|int64](params T) string {
//	return fmt.Sprintf("%v", params)
//}
func toString[T interface{}](params T) string {
	return fmt.Sprintf("%v", params)
}
