package bitbc

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	Headers
	Transactions []*Transaction
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

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func (b *Block) Print() {
	fmt.Printf("Prev. hash: %x\n", b.PrevBlockHash)
	fmt.Printf("Hash: %x\n", b.Hash)
	fmt.Println()
}

func newBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{Headers: Headers{Timestamp: time.Now().Unix(), PrevBlockHash: prevBlockHash, Hash: []byte{}}, Transactions: transactions}
	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func newGenesisBlock(coinbase *Transaction) *Block {
	return newBlock([]*Transaction{coinbase}, []byte{})
}
