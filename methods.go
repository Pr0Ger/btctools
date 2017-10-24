package btctools

import "encoding/json"

// GetBlockChainInfo provides information about the current state of the block chain.
func (c *Client) GetBlockChainInfo() (*BlockChainInfo, error) {
	var resp BlockChainInfo

	rawResp, err := c.sendRequest("getblockchaininfo")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*rawResp, &resp)
	return &resp, err
}

// GetNetworkInfo returns information about the node’s connection to the network.
func (c *Client) GetNetworkInfo() (*ClientNetworkInfo, error) {
	var resp ClientNetworkInfo

	rawResp, err := c.sendRequest("getnetworkinfo")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*rawResp, &resp)
	return &resp, err
}
