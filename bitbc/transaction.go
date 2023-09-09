package bitbc

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const subsidy = 10 //the amount of reward

type TXInput struct {
	Txid      []byte // stores the ID of such transaction
	Vout      int    //stores an index of an output in the transaction
	ScriptSig string //a script which provides data to be used in an output’s ScriptPubKey
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

type TXOutput struct {
	Value        int //stores the number of satoshis
	ScriptPubKey string
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

/*
Transactions just lock values with a script, which can be unlocked only by the one who locked them.
*/
type Transaction struct {
	ID   []byte
	Vin  []TXInput //Inputs of a new transaction reference outputs of a previous transaction, an input must reference an output
	Vout []TXOutput
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{Txid: []byte{},
		Vout:      -1,
		ScriptSig: data}
	txout := TXOutput{Value: subsidy, ScriptPubKey: to}
	tx := Transaction{ID: nil, Vin: []TXInput{txin}, Vout: []TXOutput{txout}}
	tx.SetID()

	return &tx
}

/*
Outputs are where “coins” are stored. Each output comes with an unlocking script,
which determines the logic of unlocking the output.
Every new transaction must have at least one input and output.
An input references an output from a previous transaction and provides
data (the ScriptSig field) that is used in the output’s unlocking script
to unlock it and use its value to create new outputs.
*/
