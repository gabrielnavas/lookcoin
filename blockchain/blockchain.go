package blockchain

import (
	"fmt"
	"math/rand"
	"time"
)

type Transaction struct{}
type Block struct {
	Height       int
	Transactions []Transaction
	PreviousHash string
	Nonce        string
	Hash         string
	Target       string
	Timestamp    int64
}

type Blockchain struct {
	Chain               []Block
	PendingTransactions []Transaction
	Target              string
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Chain:               []Block{},
		PendingTransactions: []Transaction{},
		Target:              "0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}
}

func (b *Blockchain) LastBlock() *Block {
	if len(b.Chain) > 0 {
		return &b.Chain[len(b.Chain)-1]
	}
	return nil
}

func (b *Blockchain) NewBlock() *Block {

	// Get last block's hash if has
	var lastHashBlock string = ""
	lastBlock := b.LastBlock()
	if lastBlock != nil {
		lastHashBlock = lastBlock.Hash
	}

	// timestamp int format
	var timestamp int64 = time.Now().UTC().UnixNano()

	// seed to random
	rand.Seed(time.Now().UnixNano())

	// the block
	block := &Block{
		Height:       len(b.Chain),
		Transactions: b.PendingTransactions,
		PreviousHash: lastHashBlock,
		Nonce:        fmt.Sprintf("%x", rand.Uint64()),
		Target:       b.Target,
		Timestamp:    timestamp,
	}

	// reset transaction
	b.PendingTransactions = []Transaction{}

	return block
}
