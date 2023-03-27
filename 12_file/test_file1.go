package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	var (
		// fileName = "/home/houzw/document/git-rep/HOUZW/golang_learn/hzw_gofiletest.txt"
		fileName = "./hzw_gofiletest.txt"
		content  = "hello golan\n"
		file     *os.File
		err      error
	)

	// fmt.Println("os.ModeDir       :", os.ModeDir.String())
	// fmt.Println("os.ModeAppend    :", os.ModeAppend.String())
	// fmt.Println("os.ModeExclusive :", os.ModeExclusive.String())
	// fmt.Println("os.ModeTemporary :", os.ModeTemporary.String())
	// fmt.Println("os.ModeSymlink   :", os.ModeSymlink.String())
	// fmt.Println("os.ModeDevice    :", os.ModeDevice.String())
	// fmt.Println("os.ModeNamedPipe :", os.ModeNamedPipe.String())
	// fmt.Println("os.ModeSocket    :", os.ModeSocket.String())
	// fmt.Println("os.ModeSetuid    :", os.ModeSetuid.String())
	// fmt.Println("os.ModeSetgid    :", os.ModeSetgid.String())
	// fmt.Println("os.ModeCharDevice:", os.ModeCharDevice.String())
	// fmt.Println("os.ModeSticky    :", os.ModeSticky.String())
	// fmt.Println("os.ModeIrregular :", os.ModeIrregular.String())
	// fmt.Println("os.ModeType      :", os.ModeType.String())
	// fmt.Println("os.ModePerm      :", os.ModePerm.String())

	if Exists(fileName) {
		//使用追加模式打开文件
		// file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
		// file, err = os.OpenFile(fileName, os.O_WRONLY, os.ModeAppend)
		file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModeAppend) // 写+追加 模式
		// file, err = os.OpenFile(fileName, os.O_APPEND, os.ModeAppend)
		if err != nil {
			fmt.Println("打开文件错误：", err)
			return
		}
	} else {
		//不存在创建文件
		file, err = os.Create(fileName)
		if err != nil {
			fmt.Println("创建失败", err)
			return
		}
	}

	defer file.Close()

	for i := 0; i < 10; i++ {
		//写入文件
		n, err := io.WriteString(file, content)

		if err != nil {
			fmt.Println("写入错误：", err)
			panic(err)
			return
		}
		fmt.Println("写入成功：n=", n)
	}

	//读取文件
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("读取错误：", err)
		return
	}
	fmt.Println("读取成功，文件内容：", string(fileContent))
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
