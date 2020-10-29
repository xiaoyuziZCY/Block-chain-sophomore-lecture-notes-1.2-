package blockchain

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
)

var BUCKET_NAME ="blocks"

var LAST_KEY ="lasthash"

var CHAINDB = "chain.db"

type BlockChain struct {
	LastHash []byte
	BoltDb    *bolt.DB
}

func NewBlockChain() BlockChain  {
	db,err :=bolt.Open(CHAINDB,0600,nil)
	if err !=nil{
		panic(err.Error())
	}
	var bl BlockChain
	db.Update(func(tx *bolt.Tx) error {
		bucket :=tx.Bucket([]byte(BUCKET_NAME))
		if bucket ==nil {
			bucket,err =tx.CreateBucket([]byte(BUCKET_NAME))
			if err !=nil {
				panic(err.Error())
			}
		}
			lastHash:=bucket.Get([]byte(LAST_KEY))
			if len(lastHash) ==0 {
				genesis :=CreateGenesisBlock()
				fmt.Printf("genesis的hash值：%x\n",genesis.Hash)
				bl =BlockChain{
					LastHash: genesis.Hash,
					BoltDb:    db,
				}
				genesisBytes,_ :=genesis.Serialize()
				bucket.Put(genesis.Hash,genesisBytes)
				bucket.Put([]byte(LAST_KEY),genesis.Hash)
		}else {
			lastHash :=bucket.Get([]byte(LAST_KEY))
			lastBlockBytes :=bucket.Get(lastHash)
			lastBlock,err :=DeSerialize(lastBlockBytes)
				if err !=nil {
					panic("读取区块链数据失败")
				}
				bl=BlockChain{
					LastHash: lastBlock.Hash,
					BoltDb:    db,
				}
			}
			return nil

	})
	return bl
}
func (bc BlockChain) SaveData(data []byte) (Block,error) {
	db := bc.BoltDb
	var e error
	var lastBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			e = errors.New("boltdb未创建，请重试!")
			return e
		}
		lastBlockBytes := bucket.Get(bc.LastHash)
		lastBlock, _ = DeSerialize(lastBlockBytes)
		return nil
	})
	newBlock := Newblock(lastBlock.Height+1, data, lastBlock.Hash)
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		newBlockBytes, _ := newBlock.Serialize()
		bucket.Put(newBlock.Hash, newBlockBytes)
		bucket.Put([]byte(LAST_KEY), newBlock.Hash)
		bc.LastHash = newBlock.Hash
		return nil
	})
	return newBlock, e
}