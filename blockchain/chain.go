package blockchain

type BlockChain struct {
	blocksArray []*Block
	blocksMap   map[string]*Block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocksArray[len(bc.blocksArray)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocksArray = append(bc.blocksArray, newBlock)
	bc.blocksMap[string(newBlock.Hash)] = newBlock
}

func NewBlockchain() *BlockChain {
	genesisBlock := NewGenesisBlock()
	return &BlockChain{blocksArray: []*Block{genesisBlock}, blocksMap: map[string]*Block{string(genesisBlock.Hash): genesisBlock}}
}
