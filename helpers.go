package btctools

import "strings"

// Forks represent daemon type
type Forks int

const (
	// ForkUnknown indicates daemon type is not yet detected
	ForkUnknown Forks = iota

	// ForkOriginal indicates original Satoshi daemon
	ForkOriginal

	// ForkLitecoin indicates Litecoin (SCrypt, lower block time)
	ForkLitecoin

	// ForkDash indicates Dash (Master nodes, PrivateSend, InstantSend)
	ForkDash
)

// GetType will try to guess what exactly
func (c *Client) GetType() (Forks, error) {
	if c.daemonType != ForkUnknown {
		return c.daemonType, nil
	}

	info, err := c.GetNetworkInfo()
	if err != nil {
		return ForkUnknown, err
	}

	if strings.Contains(info.Subversion, "Satoshi") {
		c.daemonType = ForkOriginal
	}

	return c.daemonType, nil
}
