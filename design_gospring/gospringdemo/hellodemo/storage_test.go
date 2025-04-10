package hellodemo

import (
	"log/slog"
	"testing"

	"github.com/go-spring/spring-core/conf/storage"
)

func TestStorageSplitPath(t *testing.T) {
	pstr := "a.b.c"
	path, err := storage.SplitPath(pstr)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("path", "path", path)

	pstr = "a.b.c.d[3]"
	path, err = storage.SplitPath(pstr)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("path", "path", path)

	stor := storage.NewStorage()
	// stor.Set("a.b.c", "1")
	stor.Set("a.b.c.d", "2")
	stor.Set("a.b.c.d[3]", "3")
	stor.Set("a.b.c.d[4]", "4")
	stor.Set("a.b.c.d[5]", "5")

	rawdata := stor.RawData()
	slog.Info("rawdata", "rawdata", rawdata)

	slog.Info("has a.b.c ", "bool:", stor.Has("a.b.c"))
}
