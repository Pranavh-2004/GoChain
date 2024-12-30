package main

import (
	"fmt"

	"github.com/Pranavh-2004/GoChain/internal/blockchain"
	"github.com/Pranavh-2004/GoChain/internal/storage"
)

func main() {
	// Initialize or Load Blockchain
	bc, err := storage.LoadBlockchain()
	if err != nil {
		fmt.Println("No existing blockchain found. Creating a new one...")
		bc = blockchain.NewBlockchain()
	} else {
		fmt.Println("Blockchain loaded from storage.")
	}

	// Add Blocks
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Third Block")

	// Save Blockchain
	err = storage.SaveBlockchain(bc)
	if err != nil {
		fmt.Printf("Error saving blockchain: %v\n", err)
	}

	// Print Blockchain
	for _, block := range bc.Blocks {
		fmt.Printf("Data: %s\nHash: %x\nPrevHash: %x\nNonce: %d\n\n",
			block.Data, block.Hash, block.PrevHash, block.Nonce)
	}
}