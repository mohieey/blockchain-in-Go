package bitbc

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	Headers
	Data []byte
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Fatal(err)
	}

	return result.Bytes()
}

func (b *Block) Print() {
	fmt.Printf("Prev. hash: %x\n", b.PrevBlockHash)
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Hash: %x\n", b.Hash)
	fmt.Println()
}

func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{Headers: Headers{Timestamp: time.Now().Unix(), PrevBlockHash: prevBlockHash, Hash: []byte{}}, Data: []byte(data)}
	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func newGenesisBlock() *Block {
	return newBlock("Genesis Block", []byte{})
}
