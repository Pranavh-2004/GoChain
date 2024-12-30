package blockchain

type Blockchain struct {
	Blocks []*Block
}

// Add Block function
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks) -1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain initializes the blockchain with a genesis block
func NewBlockchain() *Blockchain {
	genesisBlock := CreateBlock("Genesis", []byte{})
	return &Blockchain{Blocks: []*Block{genesisBlock}}
} 