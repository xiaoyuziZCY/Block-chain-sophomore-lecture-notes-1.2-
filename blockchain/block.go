package blockchain

import (
	"DataCertProject/blockchain"
	"bytes"
    "encoding/gob"
	"time"
)

type Block struct {
	Height int64
	TimeStamp int64
	Hash []byte
	Data []byte
	PrevHash []byte
	Version string //版本号
	Nonce int64
}

func CreateGenesisBlock() Block{
block:=Newblock(0,[]byte{},[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
return block
}

func Newblock(height int64, data []byte, prevHash []byte)(Block){
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Version:   "0x01",
	}
pow:= blockchain.NewPoW(blockchain.Block(block))
blockHash,nonce:=pow.Run()
block.Nonce=nonce
block.Hash = blockHash
return block
}
func (bk Block) Serialize() ([]byte, error) {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(bk)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
func DeSerialize(data []byte) (*Block, error) {
	var block Block
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}


