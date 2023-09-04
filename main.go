package main

import (
	"blockchain/blockchain"
	"fmt"
	"strconv"
)

func main() {
	bc := blockchain.NewBlockchain()

	b1Hash := bc.AddBlock("hello")
	bc.AddBlock("welcome")
	bc.AddBlock("holaaaa")

	bc.Print()

	b1 := bc.GetBlock(b1Hash)
	b1.Print()
	pow := blockchain.NewProofOfWork(b1)
	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))

}
