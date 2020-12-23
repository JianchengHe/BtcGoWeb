package moudles

//该结构体用于解析比特币节点返回数据
type BTCResult struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
	Id     string      `json:"id"`
}

//该结构体用于创建比特币客户端节点连接请求
type BTCJson struct {
	Jsonrpc string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

//比特币节点返回的块信息
type Blcok struct {
	Hash          string   `json:"hash"`
	Confirmations int64    `json:"confirmations"`
	Strippedsize  int64    `json:"strippedsize"`
	Size          int64    `json:"size"`
	Weight        int64    `json:"weight"`
	Height        int64    `json:"height"`
	Version       int64    `json:"version"`
	VersionHex    string   `json:"version_hex"`
	Merkleroot    string   `json:"merkleroot"`
	Tx            []string `json:"merkleroot"`
	Time          int64    `json:"time"`
	Mediantime    int64    `json:"mediantime"`
	Nonce         int64    `json:"nonce"`
	Bits          string   `json:"bits"`
	Difficulty    int64    `json:"difficulty"`
	Chainwork     string   `json:"chainwork"`
	NTx           int64    `json:"n_tx"`
	Nextblockhash string   `json:"nextblockhash"`
}
type AddressInfo struct {
	Address             string   `json:"address"`
	ScriptPubkey        string   `json:"scriptPubKey"`
	Ismine              bool     `json:"ismine"`
	solvable            bool     `json:"solvable"`
	Desc                string   `json:"desc"`
	Iswatchonly         bool     `json:"iswatchonly"`
	Isscript            bool     `json:"isscript"`
	Iswitness           bool     `json:"iswitness"`
	Witness_version     int64    `json:"witness_version"`
	Witness_program     string   `json:"witness_program"`
	Pubkey              string   `json:"pubkey"`
	Ischange            bool     `json:"ischange"`
	Timestamp           int64    `json:"timestamp"`
	Hdkeypath           string   `json:"hdkeypath"`
	Hdseedid            string   `json:"hdseedid"`
	Hdmasterfingerprint string   `json:"hdmasterfingerprint"`
	Lables              []string `json:"lables"`
}
