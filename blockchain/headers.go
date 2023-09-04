package blockchain

type Headers struct {
	Timestamp     int64
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}
