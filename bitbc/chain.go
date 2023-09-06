package bitbc

import (
	"fmt"

	"github.com/boltdb/bolt"
)

const targetBits = 24
const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type BlockChain struct {
	tip []byte
	Db  *bolt.DB
}

func (bc *BlockChain) AddBlock(data string) string {
	var lastHash []byte

	bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	newBlock := newBlock(data, lastHash)

	bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		b.Put(newBlock.Hash, newBlock.Serialize())
		b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})

	return string(newBlock.Hash)
}

func (bc *BlockChain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{currentHash: bc.tip, db: bc.Db}

	return bci
}

func (bc *BlockChain) Print() {
	bci := bc.Iterator()
	block := bci.Next()

	for block != nil {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
		block = bci.Next()
	}

}

func NewBlockchain() *BlockChain {
	var tip []byte
	db, _ := bolt.Open(dbFile, 0600, nil)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := newGenesisBlock()
			b, _ := tx.CreateBucket([]byte(blocksBucket))
			b.Put(genesis.Hash, genesis.Serialize())
			b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := BlockChain{tip, db}

	return &bc
}
