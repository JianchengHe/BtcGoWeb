package blockchain

import "time"

/*
定义区块结构体，用于表示区块
*/
type Block struct {
	Height       int64  //表示区块的高度，第几个区块
	TimeStamp    int64  //区块产生的时间戳
	PerviousHash []byte //前一个区块的哈希值
	Data         []byte //数据字段
	Hash         []byte //当前区块的哈希值
	Version      string //版本号
}

/*
创建一个新区块
*/
func NewBlock(height int64, perviousHash []byte, data []byte) Block {
	block := Block{
		Height:       height + 1,
		TimeStamp:    time.Now().Unix(),
		PerviousHash: perviousHash,
		Data:         data,
		Version:      "0X01",
	}
	//block.Hash =
	return block
}
/*
创建创世区块
 */
func CreateGenesisBlock()  Block{
	genesisBlock := NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},[]byte{0})
	return genesisBlock
}