package okx

import (
	"github.com/fengid/goex/okx/futures"
	"github.com/fengid/goex/okx/spot"
)

type OKx struct {
	Spot    *spot.Spot
	Futures *futures.Futures
	Swap    *futures.Swap
}

func New() *OKx {
	return &OKx{
		Spot:    spot.New(),
		Futures: futures.New(),
		Swap:    futures.NewSwap(),
	}
}
