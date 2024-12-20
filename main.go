package goex

import (
	"github.com/fengid/goex/httpcli"
	"github.com/fengid/goex/logger"
	"github.com/fengid/goex/okx"
	"reflect"
)

var (
	DefaultHttpCli = httpcli.Cli
)

var (
	OKx = okx.New()
)

func SetDefaultHttpCli(cli httpcli.IHttpClient) {
	logger.Infof("use new http client implement: %s", reflect.TypeOf(cli).Elem().String())
	httpcli.Cli = cli
}
