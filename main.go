package main

import (
	"blockchain"
	"fmt"
	"time"
)

func main() {
	genesisBlock := &blockchain.Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	genesisBlock.Hash = blockchain.calculateHash(genesisBlock)

	blockchain := &blockchain.Blockchain{Blocks: []*blockchain.Block{genesisBlock}}

	blockchain.NewBlock("Transaction 1")
	blockchain.NewBlock("Transaction 2")

	fmt.Println("All Blocks:")
	blockchain.DisplayAllBlocks()

	fmt.Println("Modifying Block 1...")
	blockchain.ModifyBlock(1, "Modified Data")

	fmt.Println("\nAll Blocks after modification:")
	blockchain.DisplayAllBlocks()
}
