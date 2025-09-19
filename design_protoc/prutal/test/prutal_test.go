package test

import (
	"design_protoc/util"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cloudwego/prutal/prutalgen/pkg/prutalgen"
)

func TestPrutalHello(t *testing.T) {
	fmt.Print("hello world")

	// x := prutalgen.NewLoader([]string{"/home/houzw/document/git-rep/HOUZW/golang/aic-gospring/ag-core/tool/aggen/parser/prutal/test/third_party"}, nil)
	x := prutalgen.NewLoader([]string{"/home/houzw/document/git-rep/HOUZW/golang/aic-gospring/ag-core/tool/aggen/parser/prutal/test/third_party"}, nil)

	ff := x.LoadProto(filepath.Clean("/home/houzw/document/git-rep/HOUZW/golang/aic-gospring/ag-core/tool/aggen/parser/prutal/test/hzw1.proto"))

	for _, f := range ff {
		if !strings.HasPrefix(f.GoImport, "github.hzw") {
			continue
		}

		fmt.Printf("======pkg: %s, GoImport:%s\n", f.Package, f.GoImport)

		fmt.Printf("PkgName: %s\n", f.Package)
		fmt.Printf("PkgRefName: %s\n", f.GoPackage)
		fmt.Printf("ImportPath: %s\n", f.GoImport)
		// util.TestLog("xxxxxx", f)

		svc0 := f.Services[0]
		m0 := svc0.Methods[0]

		util.TestLog("xxxxxx", m0.Options)

	}
}
