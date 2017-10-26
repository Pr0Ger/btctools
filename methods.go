package btctools

import (
	"encoding/json"

	"github.com/Pr0Ger/btctools/blockchain"
)

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

// GetBlockHeader gets a block header with a particular header hash from the local block database
func (c *Client) GetBlockHeader(hash *blockchain.BlockHash) (*BlockHeader, error) {
	var resp BlockHeader

	rawResp, err := c.sendRequest("getblockheader", hash.String())
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

func (c *Client) ListSinceBlock(hash *blockchain.BlockHash) (*ListSinceBlockResult, error) {
	var resp ListSinceBlockResult

	rawResp, err := c.sendRequest("listsinceblock", hash.String())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(*rawResp, &resp)
	return &resp, err
}
