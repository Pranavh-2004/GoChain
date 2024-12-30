package storage

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Pranavh-2004/GoChain/internal/blockchain"
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"      // Path to database
const blockBucket = "blocks"        // Bucket for storing blocks in BoltDB (Buckets are like containers of key-value pairs in BoltDB)

// SaveBlockchain stores the blockchain in the database
func SaveBlockchain(bc *blockchain.Blockchain) error { 
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return fmt.Errorf("could not open db: %v", err)
	}
	defer db.Close()

	// Update database
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(blockBucket))
		if err != nil {
			return fmt.Errorf("could not create bucket: %v", err)
		}

		// Iterate through the blocks in the blockchain
		for i, block := range bc.Blocks {
			blockData, err := json.Marshal(block)
			if err != nil {
				return fmt.Errorf("could not marshal block: %v", err)
			}

			// Use the block index as the key
			err = b.Put([]byte(fmt.Sprintf("%d", i)), blockData)
			if err != nil {
				return fmt.Errorf("could not save block to database: %v", err)
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("Failed to save blockchain: %v", err)
	}
	return err
}

// LoadBlockchain loads the blockchain from the BoltDB file
func LoadBlockchain() (*blockchain.Blockchain, error) {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db: %v", err)
	}
	defer db.Close()

	var bc blockchain.Blockchain

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		if b == nil {
			return fmt.Errorf("bucket %s not found", blockBucket)
		}

		// Iterate over all blocks in the bucket
		return b.ForEach(func(k, v []byte) error {
			var block blockchain.Block
			err := json.Unmarshal(v, &block)
			if err != nil {
				return fmt.Errorf("could not unmarshal block: %v", err)
			}

			// Validate the block
			pow := blockchain.NewProofOfWork(&block)
			if !pow.Validate() {
				return fmt.Errorf("block validation failed for key %s", k)
			}

			bc.Blocks = append(bc.Blocks, &block)
			return nil
		})
	})

	if err != nil {
		return nil, fmt.Errorf("could not load blockchain: %v", err)
	}

	return &bc, nil
}