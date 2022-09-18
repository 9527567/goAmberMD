package service

import (
	"bytes"
	"fmt"
	"io"
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

func toString[T any](params T) string {
	return fmt.Sprintf("%v", params)
}

func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
func WriteFile(fileName string, content []string) {
	var file *os.File
	var err error
	if exists(fileName) {
		//使用追加模式打开文件
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println("can not remove file:", err)
			return
		}
		file, err = os.Create(fileName)
		if err != nil {
			fmt.Println("create file error:", err)
			return
		}
	} else {
		//不存在创建文件
		file, err = os.Create(fileName)
		if err != nil {
			fmt.Println("create file error:", err)
			return
		}
	}

	defer file.Close()
	//写入文件
	for _, s := range content {
		_, err := io.WriteString(file, s)
		if err != nil {
			fmt.Println("write error:", err)
			return
		}
	}
}
