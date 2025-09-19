package test

import (
	"design_protoc/util"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/emicklei/proto"
)

func TestEProtoHello(t *testing.T) {
	reader, err := os.Open(filepath.Clean("/home/houzw/document/git-rep/HOUZW/golang/aic-gospring/ag-core/tool/aggen/parser/emicklei_proto/test/hzw1.proto"))
	if err != nil {
		fmt.Errorf("open file failed, err: %v", err)
		return
	}
	defer reader.Close()
	parser := proto.NewParser(reader)
	definiton, err := parser.Parse()
	if err != nil {
		fmt.Errorf("parse file failed, err: %v", err)
		return
	}

	util.TestLog("TestEProtoHello", definiton)

}
