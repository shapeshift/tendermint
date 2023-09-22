package psql

import (
	"github.com/shapeshift/tendermint/state/indexer"
	"github.com/shapeshift/tendermint/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
