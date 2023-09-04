package blockchain

import "fmt"

type BlockChain struct {
	blocksArray []*Block
	blocksMap   map[string]*Block
}

func (bc *BlockChain) AddBlock(data string) string {
	prevBlock := bc.blocksArray[len(bc.blocksArray)-1]
	newBlock := newBlock(data, prevBlock.Hash)
	bc.blocksArray = append(bc.blocksArray, newBlock)
	bc.blocksMap[string(newBlock.Hash)] = newBlock
	return string(newBlock.Hash)
}

func (bc *BlockChain) Print() {
	for _, block := range bc.blocksArray {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}

func (bc *BlockChain) GetBlock(hash string) *Block {
	return bc.blocksMap[hash]
}

func NewBlockchain() *BlockChain {
	genesisBlock := newGenesisBlock()
	return &BlockChain{blocksArray: []*Block{genesisBlock}, blocksMap: map[string]*Block{string(genesisBlock.Hash): genesisBlock}}
}
