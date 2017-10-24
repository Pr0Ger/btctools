package btctools

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
	Softforks            []struct {
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
		Proxy                     string `json:"proxy"`
		ProxyRandomizeCredentials bool   `json:"proxy_randomize_credentials"`
	} `json:"networks"`
	RelayFee       float64 `json:"relayfee"`
	LocalAddresses []struct {
		Address string `json:"address"`
		Port    uint16 `json:"port"`
		Score   uint32 `json:"score"`
	} `json:"localaddresses"`
	Warnings string `json:"warnings"`
}
