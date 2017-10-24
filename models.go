package btctools

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
