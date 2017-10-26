package btctools

import (
	"encoding/json"

	"github.com/Pr0Ger/btctools/blockchain"
)

func (c *Client) getRPCResponse(method string, response interface{}, params ...interface{}) error {
	rawResp, err := c.sendRequest(method, params...)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*rawResp, &response)
	return err
}

// GetBlockChainInfo provides information about the current state of the block chain.
func (c *Client) GetBlockChainInfo() (*BlockChainInfo, error) {
	var resp BlockChainInfo

	err := c.getRPCResponse("getblockchaininfo", &resp)

	return &resp, err
}

// GetBlockHeader gets a block header with a particular header hash from the local block database
func (c *Client) GetBlockHeader(hash *blockchain.BlockHash) (*BlockHeader, error) {
	var resp BlockHeader

	err := c.getRPCResponse("getblockheader", &resp, hash.String())

	return &resp, err
}

// GetNetworkInfo returns information about the node’s connection to the network.
func (c *Client) GetNetworkInfo() (*ClientNetworkInfo, error) {
	var resp ClientNetworkInfo

	err := c.getRPCResponse("getnetworkinfo", &resp)

	return &resp, err
}

// GetNewAddress returns a new address for receiving payments
func (c *Client) GetNewAddress(account string) (blockchain.Address, error) {
	var resp string

	err := c.getRPCResponse("getnewaddress", &resp, account)
	if err != nil {
		return nil, err
	}
	return blockchain.DecodeAddress(resp)
}

// ListSinceBlock gets all transactions affecting the wallet which have occurred since a particular block,
// plus the header hash of a block at a particular depth.
func (c *Client) ListSinceBlock(hash *blockchain.BlockHash) (*ListSinceBlockResult, error) {
	var resp ListSinceBlockResult

	err := c.getRPCResponse("listsinceblock", &resp)

	return &resp, err
}
