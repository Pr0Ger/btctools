package btctools

import "github.com/Pr0Ger/btctools/blockchain"

// BlockChainInfo is a response for `GetBlockChainInfo` RPC call
type BlockChainInfo struct {
	Chain                string  `json:"chain"`
	Blocks               uint64  `json:"blocks"`
	Headers              uint64  `json:"headers"`
	BestBlockHash        string  `json:"bestblockhash"`
	Difficulty           float64 `json:"difficulty"`
	MedianTime           int64   `json:"mediantime"`
	VerificationProgress float64 `json:"verificationprogress"`
	ChainWork            string  `json:"chainwork"`
	Pruned               bool    `json:"pruned"`
	PruneHeight          int64   `json:"pruneheight,omitempty"`
	SoftForks            []struct {
		ID      string `json:"id"`
		Version int    `json:"version"`
		Reject  struct {
			Status   bool   `json:"status,omitempty"`
			Found    uint64 `json:"found,omitempty"`
			Required uint64 `json:"required"`
			Window   uint64 `json:"window"`
		} `json:"reject"`
	} `json:"softforks"`
	// bip9_softforks is broken with Dash
}

// BlockHeader is a response for `GetBlockHeader` RPC call
type BlockHeader struct {
	Hash              string               `json:"hash"`
	Confirmations     int64                `json:"confirmations"`
	Height            int64                `json:"height"`
	Version           int64                `json:"version"`
	VersionHex        string               `json:"versionHex"`
	MerkleRoot        string               `json:"merkleroot"`
	Time              uint64               `json:"time"`
	MedianTime        int64                `json:"mediantime"`
	Nonce             int64                `json:"nonce"`
	Bits              string               `json:"bits"`
	Difficulty        float64              `json:"difficulty"`
	ChainWork         string               `json:"chainwork"`
	PreviousBlockHash blockchain.BlockHash `json:"previousblockhash"`
	NextBlockHash     blockchain.BlockHash `json:"nextblockhash"`
}

// ClientNetworkInfo is a response for `GetNetworkInfo` RPC call
type ClientNetworkInfo struct {
	Version         int64  `json:"version"`
	Subversion      string `json:"subversion"`
	ProtocolVersion int64  `json:"protocolversion"`
	LocalServices   string `json:"localservices"`
	LocalRelay      bool   `json:"localrelay"`
	TimeOffset      int64  `json:"timeoffset"`
	Connections     uint64 `json:"connections"`
	Networks        []struct {
		Name                      string `json:"name"`
		Limited                   bool   `json:"limited"`
		Reachable                 bool   `json:"reachable"`
		ProxyRandomizeCredentials bool   `json:"proxy_randomize_credentials"`
		Proxy                     string `json:"proxy"`
	} `json:"networks"`
	RelayFee       float64 `json:"relayfee"`
	LocalAddresses []struct {
		Address string `json:"address"`
		Port    uint16 `json:"port"`
		Score   uint32 `json:"score"`
	} `json:"localaddresses"`
	Warnings string `json:"warnings"`
}

// TxCategory represents type of transaction
type TxCategory string

const (
	TxCategorySend     = "send"
	TxCategoryReceive  = "receive"
	TxCategoryGenerate = "generate"
	TxCategoryImmature = "immature"
	TxCategoryOrphan   = "orphan"
)

// ListSinceBlockResult is a response for `ListSinceBlock` RPC call
type ListSinceBlockResult struct {
	Transactions []struct {
		Address blockchain.AddressPubKeyHash `json:"address,omitempty"`
		Amount  float64                      `json:"amount"`
		//bip125-replaceable": "no",
		Blockhash     blockchain.BlockHash `json:"blockhash,omitempty"`
		BlockIndex    int64                `json:"blockindex,omitempty"`
		BlockTime     uint64               `json:"blocktime,omitempty"`
		Category      TxCategory           `json:"category"`
		Confirmations uint64               `json:"confirmations"`
		Label         string               `json:"label"`
		Time          uint64               `json:"time"`
		TimeReceived  uint64               `json:"timereceived"`
		TxID          string               `json:"txid"`
		Vout          uint32               `json:"vout"`
		//walletconflicts": []
	} `json:"transactions"`
	LastBlock blockchain.BlockHash `json:"lastblock"`
}
