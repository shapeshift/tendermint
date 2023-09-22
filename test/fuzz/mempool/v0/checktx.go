package v0

import (
	"github.com/shapeshift/tendermint/abci/example/kvstore"
	"github.com/shapeshift/tendermint/config"
	mempl "github.com/shapeshift/tendermint/mempool"
	mempoolv0 "github.com/shapeshift/tendermint/mempool/v0"
	"github.com/shapeshift/tendermint/proxy"
)

var mempool mempl.Mempool

func init() {
	app := kvstore.NewApplication()
	cc := proxy.NewLocalClientCreator(app)
	appConnMem, _ := cc.NewABCIClient()
	err := appConnMem.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false
	mempool = mempoolv0.NewCListMempool(cfg, appConnMem, 0)
}

func Fuzz(data []byte) int {
	err := mempool.CheckTx(data, nil, mempl.TxInfo{})
	if err != nil {
		return 0
	}

	return 1
}
