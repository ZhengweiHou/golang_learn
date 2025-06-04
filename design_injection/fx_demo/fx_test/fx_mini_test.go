package fx_test

import (
	"context"
	"testing"

	"go.uber.org/fx"
)


func TestMini(t *testing.T) {
	fx.New().Run()
}

func TestMini2(t *testing.T) {
	app := fx.New()
	app.Start(context.Background())
	app.Stop(context.Background())
}
