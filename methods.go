package btctools

import "encoding/json"

func (c *Client) GetBlockChainInfo() (*BlockChainInfo, error) {
	var resp BlockChainInfo

	rawResp, err := c.sendRequest("getblockchaininfo")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*rawResp, &resp)
	return &resp, err
}

func (c *Client) GetNetworkInfo() (*ClientNetworkInfo, error) {
	var resp ClientNetworkInfo

	rawResp, err := c.sendRequest("getnetworkinfo")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*rawResp, &resp)
	return &resp, err
}
