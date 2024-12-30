package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const Difficulty = 18

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// Creating a ProofOfWork instance with the target calculated based on difficulty
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) 
	return &ProofOfWork{Block: block, Target: target}
}

// Run executes proof of work algorithm
func (pow *ProofOfWork) Run() ([]byte, int) {
	var intHash big.Int
	var hash [32]byte
	nonce := 0

	for {
		// Generate data by combining block properties with nonce
		data := pow.Block.InitData(nonce)
		hash = sha256.Sum256(data) 
		intHash.SetBytes(hash[:])  

		// Compare the hash with the target
		if intHash.Cmp(pow.Target) == -1 {
			fmt.Printf("Valid hash found: %x\n", hash)
			break
		}
		nonce++
	}
	return hash[:], nonce
}

// InitData is used to prepare block's data for hashing
func (b *Block) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			b.PrevHash,
			b.Data,
			[]byte(fmt.Sprintf("%d", nonce)),
		},
		[]byte{},
	)
	return data
}

// Validate verifies if the block's hash meets the proof-of-work criteria
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	// Generate data for validation using the block's stored nonce
	data := pow.Block.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data) 
	intHash.SetBytes(hash[:])   

	// Check if the calculated hash is below the target
	return intHash.Cmp(pow.Target) == -1
}