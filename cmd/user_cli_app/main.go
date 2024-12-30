package main

import (
	"bufio"
	"fmt"
	"os"

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

	// Set up reader for user input
	reader := bufio.NewReader(os.Stdin)

	// Take user input for adding blocks
	for {
		fmt.Print("Enter data for the block (or type 'exit' to stop): ")
		data, _ := reader.ReadString('\n')
		data = data[:len(data)-1] // Remove the newline character

		if data == "exit" {
			break
		}

		// Add the block with the user-provided data
		bc.AddBlock(data)
	}

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