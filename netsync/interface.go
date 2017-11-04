// Copyright (c) 2017 The bchsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netsync

import (
	"github.com/bchsuite/bchd/blockchain"
	"github.com/bchsuite/bchd/chaincfg"
	"github.com/bchsuite/bchd/chaincfg/chainhash"
	"github.com/bchsuite/bchd/mempool"
	"github.com/bchsuite/bchd/peer"
	"github.com/bchsuite/bchd/wire"
	"github.com/bchsuite/bchutil"
)

// PeerNotifier exposes methods to notify peers of status changes to
// transactions, blocks, etc. Currently server (in the main package) implements
// this interface.
type PeerNotifier interface {
	AnnounceNewTransactions(newTxs []*mempool.TxDesc)

	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)

	RelayInventory(invVect *wire.InvVect, data interface{})

	TransactionConfirmed(tx *bchutil.Tx)
}

// Config is a configuration struct used to initialize a new SyncManager.
type Config struct {
	PeerNotifier PeerNotifier
	Chain        *blockchain.BlockChain
	TxMemPool    *mempool.TxPool
	ChainParams  *chaincfg.Params

	DisableCheckpoints bool
	MaxPeers           int
}
