package blockchain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddressPubKeyHash_UnmarshalJSON(t *testing.T) {
	in := `"17VZNX1SN5NtKa8UQFxwQbFeFc3iqRYhem"`

	addr := new(AddressPubKeyHash)
	err := addr.UnmarshalJSON([]byte(in))

	require.NoError(t, err)
	require.Equal(t, "17VZNX1SN5NtKa8UQFxwQbFeFc3iqRYhem", addr.String())
}

func TestDecodeAddress(t *testing.T) {
	tests := []struct {
		address string
	}{
		// Main net (Bitcoin)
		{`17VZNX1SN5NtKa8UQFxwQbFeFc3iqRYhem`},
		{`1AfPimL3ZYHnDVt7X8dHhmMWjCQ7PNjSpQ`},
		// Test net (Bitcoin)
		{`mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn`},
		{`mtHeXNNCuSotNyTqYCGvwtBmRp3MY2SyHT`},
		// Main net (Litecoin)
		{`LLXywUAU8c7N65GARYR2M3AzFkK3Jf1vt6`},
	}
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			addr, err := DecodeAddress(test.address)

			require.NoError(t, err)
			require.Equal(t, test.address, addr.String())
		})
	}
}
