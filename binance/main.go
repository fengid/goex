package binance

import (
	"github.com/fengid/goex/binance/futures/fapi"
	"github.com/fengid/goex/binance/spot"
)

type Binance struct {
	Spot *spot.Spot
	Swap *fapi.FApi
}

func New() *Binance {
	return &Binance{
		Spot: spot.New(),
		Swap: fapi.NewFApi(),
	}
}
