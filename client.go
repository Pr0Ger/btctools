package btctools

import "sync/atomic"

type ConnConfig struct {
	// Host is the IP address and port of the RPC server you want to connect to
	Host string

	// User is the username to use to authenticate to the RPC server.
	User string
	// Pass is the password to use to authenticate to the RPC server.
	Pass string
}

type Client struct {
	// config holds the connection configuration associated with this client.
	config *ConnConfig

	// id for next RPC request
	id uint64
}

func (c *Client) nextID() uint64 {
	return atomic.AddUint64(&c.id, 1)
}

func New(config *ConnConfig) (*Client, error) {
	client := Client{
		config: config,
	}
	return &client, nil
}
