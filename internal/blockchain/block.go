package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Generate hash
func (b *Block) CalculateHash() []byte {
	// Convert nonce to a byte slice 
	nonceBytes := []byte(fmt.Sprintf("%d", b.Nonce))

	data := bytes.Join([][]byte{
		b.PrevHash,
		b.Data,
		nonceBytes,
	}, []byte{})

	hash := sha256.Sum256(data)
	return hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{Data: []byte(data), PrevHash: prevHash, Nonce: 0}
	pow := NewProofOfWork(block) // Create the ProofOfWork instance
	hash, nonce := pow.Run()     // Run the proof of work
	block.Hash = hash            
	block.Nonce = nonce          
	return block
}