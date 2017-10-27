package blockchain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var mainNetGenesisHash = BlockHash{
	0x6f, 0xe2, 0x8c, 0x0a, 0xb6, 0xf1, 0xb3, 0x72,
	0xc1, 0xa6, 0xa2, 0x46, 0xae, 0x63, 0xf7, 0x4f,
	0x93, 0x1e, 0x83, 0x65, 0xe1, 0x5a, 0x08, 0x9c,
	0x68, 0xd6, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00,
}

var litecoinGenesisHash = BlockHash{
	0xe2, 0xbf, 0x04, 0x7e, 0x7e, 0x5a, 0x19, 0x1a,
	0xa4, 0xef, 0x34, 0xd3, 0x14, 0x97, 0x9d, 0xc9,
	0x98, 0x6e, 0x0f, 0x19, 0x25, 0x1e, 0xda, 0xba,
	0x59, 0x40, 0xfd, 0x1f, 0xe3, 0x65, 0xa7, 0x12,
}

func TestBlockHash_UnmarshalJSON(t *testing.T) {
	in := `"000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"`

	hash := new(BlockHash)
	hash.UnmarshalJSON([]byte(in))

	require.Equal(t, mainNetGenesisHash, *hash)
}

func TestBlockHash_String(t *testing.T) {
	res := mainNetGenesisHash.String()

	require.Equal(t, `000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f`, res)
}

func TestNewHashFromStr(t *testing.T) {
	tests := []struct {
		in   string
		want BlockHash
	}{
		// Bitcoin genesis block
		{
			`000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f`,
			mainNetGenesisHash,
		},
		// Bitcoin genesis block with stripped leading zeroes
		{
			`19d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f`,
			mainNetGenesisHash,
		},
		// Litecoin genesis block
		{
			`12a765e31ffd4059bada1e25190f6e98c99d9714d334efa41a195a7e7e04bfe2`,
			litecoinGenesisHash,
		},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			hash, err := NewHashFromStr(test.in)

			require.NoError(t, err)
			require.Equal(t, test.want, *hash)
		})
	}
}
