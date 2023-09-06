package bitbc

import (
	"fmt"

	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		if len(encodedBlock) == 0 {
			fmt.Println("End of the chain")
			block = nil
		} else {
			block = DeserializeBlock(encodedBlock)
		}

		return nil
	})

	if block != nil {
		i.currentHash = block.PrevBlockHash
	}

	return block
}
