package base

import (
	"fmt"
	"os"
	"testing"
)

func TestOs1(t *testing.T) {

	fmt.Printf("os.ModeDir       :%s\n", os.ModeDir)
	fmt.Printf("os.ModeAppend    :%s\n", os.ModeAppend)
	fmt.Printf("os.ModeExclusive :%s\n", os.ModeExclusive)
	fmt.Printf("os.ModeTemporary :%s\n", os.ModeTemporary)
	fmt.Printf("os.ModeSymlink   :%s\n", os.ModeSymlink)
	fmt.Printf("os.ModeDevice    :%s\n", os.ModeDevice)
	fmt.Printf("os.ModeNamedPipe :%s\n", os.ModeNamedPipe)
	fmt.Printf("os.ModeSocket    :%s\n", os.ModeSocket)
	fmt.Printf("os.ModeSetuid    :%s\n", os.ModeSetuid)
	fmt.Printf("os.ModeSetgid    :%s\n", os.ModeSetgid)
	fmt.Printf("os.ModeCharDevice:%s\n", os.ModeCharDevice)
	fmt.Printf("os.ModeSticky    :%s\n", os.ModeSticky)
	fmt.Printf("os.ModeIrregular :%s\n", os.ModeIrregular)

	fmt.Printf("%s\n", os.ModeDir)
	var f1 os.FileMode = 0777
	fmt.Printf("%s\n", f1)
}

func TestOs2(t *testing.T) {
	ep, _ := os.Executable()
	fmt.Printf("ep:%v\n", ep)

	_, err := os.Stat("/tmp/hhh")
	fmt.Printf("ep:%v\n", err)
	if os.IsExist(err) {
		fmt.Println("=====")
	}
}
