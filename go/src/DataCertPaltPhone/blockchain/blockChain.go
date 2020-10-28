package blockchain

import "github.com/boltdb/bolt-master"

const BLOCKCHAIN  = "blockchain.db"
const BUCKET_NAME  = "blocks"
const LAST_HASH  = "lasthash"
/*
区块链结构体的定义，代表的是一条链

bolt.Db 功能
① 将新的区块数据与已有区块连接
②查询某个区块的数据和信息
③遍历区块信息
*/
type BlockChain struct {
	LastHash []byte//表示区块链中最新的区块的哈希，用于查找最新的区块内容
	BoltDb   *bolt.DB//区块链中操作区块数据文件的数据库操作对象
}
//创建一条区块链
func NewBlockChain()  BlockChain{
	var bc BlockChain
	//先打开文件
	db,err := bolt.Open(BLOCKCHAIN,0600,nil)
	//查看chain.db文件
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {//说明没有桶，创建新桶
			bucket,err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		LashHash := bucket.Get([]byte(LAST_HASH))
		if len(LashHash) == 0 {//桶中没有lasthash记录，需要创建
			//创世区块
			genesis := CreateGenesisBlock()
			//序列化以后的数据
			genesisBytes := genesis.Serialize()
			//创世区块保存到boltDb中
			bucket.Put(genesis.Hash,genesisBytes)
			//更新指向最新区块的lasthash
			bucket.Put([]byte(LAST_HASH),genesis.Hash)
			bc = BlockChain{
				LastHash: genesis.Hash,
				BoltDb: db,
			}
		}else {
			//桶中已有lasthash的记录，不再需要创世区块，只需要读取即可
			lastHash := bucket.Get([]byte(LAST_HASH))
			bc = BlockChain{
				LastHash: lastHash,
				BoltDb:   db,
			}
		}
		return nil
	})
	return bc
}
/*
用户存数据的时候调用，保存到数据库当中，先 生成一个新区块，然后将新区块添加到区块链中
 */
func (bc BlockChain)SaveData(data []byte)  {
	//从文件中读取到最新的区块。
	db := bc.BoltDb
	var lastBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(LAST_HASH))
		if bucket == nil {
			panic("读取区块链数据失败")
		}
		//lashHash := bucket.Get([]byte(LAST_HASH))
		lastBlockBytes := bucket.Get(bc.LastHash)
		//反序列化
		lastBlock,_ = DeSerialize(lastBlockBytes)
		return nil
	})
	//新建一个区块
	newBlock := NewBlock(lastBlock.Height+1,lastBlock.Hash,data)
	//把新区块存到文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("BUCKET_NAME"))
		//把新创建的区块存入到boltdb数据库中
		bucket.Put(newBlock.Hash,newBlock.Serialize())
		//更新LASTHASH对应的值，更新为最新存储的区块的hash值
		bucket.Put([]byte(LAST_HASH),newBlock.Hash)
		//将区块链实例的LASTHASH值更新为最新区块的哈希
		bc.LastHash = newBlock.Hash
		return nil
	})
}