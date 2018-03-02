package blockchain_prework

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// Blockchain is our global blockchain.
var Blockchain []Block

// Block is our basic data structure!
type Block struct {
	Data      string
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
}

// InitBlockchain creates our first Genesis node.
func InitBlockchain() {
	genesisBlock := Block{"Generis Block", time.Now().Unix(), []byte{}, []byte{}}
	genesisBlock.Hash = genesisBlock.calculateHash()
	Blockchain = []Block{genesisBlock}
}

// NewBlock creates a new Blockchain Block.
func NewBlock(oldBlock Block, data string) Block {
	block := Block{
		data,
		time.Now().Unix(),
		oldBlock.Hash,
		[]byte{},
	}
	block.Hash = block.calculateHash()
	return block
}

// AddBlock adds a new block to the Blockchain.
func AddBlock(b Block) error {
	spew.Dump(Blockchain)

	lastBlock := Blockchain[len(Blockchain)-1]
	if !bytes.Equal(b.PrevHash, lastBlock.Hash) {
		return errors.New("invalid block prev hash")
	}

	Blockchain = append(Blockchain, b)

	return nil
}

func (b *Block) calculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := []byte(b.Data)
	headers := bytes.Join([][]byte{b.PrevHash, data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}
