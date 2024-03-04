package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// Blockchain represents the blockchain
type Blockchain struct {
	Blocks []*Block
}

// NewBlock creates a new block
func (bc *Blockchain) NewBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := &Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}

// DisplayAllBlocks displays all blocks in the blockchain
func (bc *Blockchain) DisplayAllBlocks() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}

// ModifyBlock modifies a block with new data
func (bc *Blockchain) ModifyBlock(index int, newData string) {
	if index < 0 || index >= len(bc.Blocks) {
		fmt.Println("Invalid block index")
		return
	}
	block := bc.Blocks[index]
	block.Data = newData
	block.Hash = calculateHash(block)
}

// calculateHash calculates the hash of a block
func calculateHash(block *Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
