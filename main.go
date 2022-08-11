package main

import (
	"blockchain/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.NewBlockchain()
	block := bc.NewBlock()
	fmt.Println(block)
}
