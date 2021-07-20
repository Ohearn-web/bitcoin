package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

const genesisInfo = "创世块"

type Block struct {

	Version uint64 //版本号

	PrevBlockHash []byte //前区块哈希

	MerKleRoot []byte //默克尔根

	TimeStamp uint64 //时间戳

	Difficulty uint64 //挖矿的难度值

	Nonce uint64 //随机数，挖矿找的就是它

	Data []byte //数据

	Hash []byte //当前区块哈希
}

//创建区块，对Block的每一个字段填充数据即可
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version: 00,

		PrevBlockHash: prevBlockHash,

		MerKleRoot: []byte{},

		TimeStamp: uint64(time.Now().Unix()),

		Difficulty: 10, //先定义为10

		Nonce: 10, //同Difficulty

		Data: []byte(data),

		Hash: []byte{},
	}

	block.SetHash()

	return &block
}

//为了生成区块哈希，实现一个简单的函数，来计算哈希值，没有随机值，没有难度值
func (block *Block) SetHash() {

	tmp := [][]byte{
		uintToByte(block.Version),
		block.PrevBlockHash,
		block.MerKleRoot,
		uintToByte(block.TimeStamp),
		uintToByte(block.Difficulty),
		block.Data,
		uintToByte(block.Nonce),
	}

	data := bytes.Join(tmp, []byte{})

	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}
