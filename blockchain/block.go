package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Headers
	Data []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func (b *Block) Print() {
	fmt.Printf("Prev. hash: %x\n", b.PrevBlockHash)
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Hash: %x\n", b.Hash)
	fmt.Println()
}

func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{Headers: Headers{Timestamp: time.Now().Unix(), PrevBlockHash: prevBlockHash, Hash: []byte{}}, Data: []byte(data)}
	block.SetHash()
	return block
}

func newGenesisBlock() *Block {
	return newBlock("Genesis Block", []byte{})
}
