package btctools

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Pr0Ger/btctools/blockchain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRPCCall(t *testing.T, response string, tester func(client *Client)) {
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	client, _ := New(&ConnConfig{
		Host: ts.URL[7:],
	})

	tester(client)
}

func TestClient_GetBlockChainInfo(t *testing.T) {
	response := `{
    "error": null,
    "id": 1,
    "result": {
        "bestblockhash": "000000000505975c1a91cb553dd896e15f6ae8e110366fd1024efac9fa3bfa30",
        "bip9_softforks": {
            "csv": {
                "since": 770112,
                "startTime": 1456790400,
                "status": "active",
                "timeout": 1493596800
            },
            "segwit": {
                "since": 834624,
                "startTime": 1462060800,
                "status": "active",
                "timeout": 1493596800
            }
        },
        "blocks": 1210950,
        "chain": "test",
        "chainwork": "0000000000000000000000000000000000000000000000318e744f1a22f5d028",
        "difficulty": 1,
        "headers": 1210950,
        "mediantime": 1508846329,
        "pruned": false,
        "softforks": [
            {
                "id": "bip34",
                "reject": {
                    "status": true
                },
                "version": 2
            },
            {
                "id": "bip66",
                "reject": {
                    "status": true
                },
                "version": 3
            },
            {
                "id": "bip65",
                "reject": {
                    "status": true
                },
                "version": 4
            }
        ],
        "verificationprogress": 0.9999955420699167
    }
}`

	testRPCCall(t, response, func(client *Client) {
		blockChainInfo, err := client.GetBlockChainInfo()
		require.NoError(t, err)

		assert.EqualValues(t, blockChainInfo.BestBlockHash, "000000000505975c1a91cb553dd896e15f6ae8e110366fd1024efac9fa3bfa30")
	})
}

func TestClient_GetBlockHeader(t *testing.T) {
	response := `{
    "error": null,
    "id": 1,
    "result": {
        "bits": "1a236480",
        "chainwork": "00000000000000000000000000000000000000000000002a5d722cf794fa35ac",
        "confirmations": 8314,
        "difficulty": 474024.8065780034,
        "hash": "00000000000021420990192c4e6143f51f024a6ae9b0312bb11119462fcbdebf",
        "height": 1202774,
        "mediantime": 1506956494,
        "merkleroot": "e35f0aa03bb3a187a73ecd166d54c2b21965505d328da1cddd26d4bf4964aabb",
        "nextblockhash": "000000000000148b91151d83e3b3db9f6d8ce28985ff6ad34ec53e08390a75a9",
        "nonce": 3807933500,
        "previousblockhash": "00000000000012c753de0f61e2d6d5af569a0f6ddb0cda5e36edb1e5129a1d0b",
        "time": 1506960823,
        "version": 536870912,
        "versionHex": "20000000"
    }
}`

	testRPCCall(t, response, func(client *Client) {
		hash, _ := blockchain.NewHashFromStr("00000000000021420990192c4e6143f51f024a6ae9b0312bb11119462fcbdebf")
		blockHeader, err := client.GetBlockHeader(hash)

		require.NoError(t, err)
		require.EqualValues(t, 8314, blockHeader.Confirmations)
	})
}

func TestClient_GetNetworkInfo(t *testing.T) {
	response := `{
"error": null,
"id": 1,
"result": {
	"connections": 8,
	"incrementalfee": 1e-05,
	"localaddresses": [],
	"localrelay": true,
	"localservices": "000000000000000d",
	"networkactive": true,
	"networks": [
		{
			"limited": false,
			"name": "ipv4",
			"proxy": "",
			"proxy_randomize_credentials": false,
			"reachable": true
		},
		{
			"limited": false,
			"name": "ipv6",
			"proxy": "",
			"proxy_randomize_credentials": false,
			"reachable": true
		},
		{
			"limited": true,
			"name": "onion",
			"proxy": "",
			"proxy_randomize_credentials": false,
			"reachable": false
		}
	],
	"protocolversion": 70015,
	"relayfee": 1e-05,
	"subversion": "/Satoshi:0.14.1/",
	"timeoffset": -1,
	"version": 140100,
	"warnings": "Warning: unknown new rules activated (versionbit 28)"
}
}`

	testRPCCall(t, response, func(client *Client) {
		networkInfo, err := client.GetNetworkInfo()
		require.NoError(t, err)

		assert.EqualValues(t, networkInfo.Version, 140100)
	})
}

func TestClient_GetNewAddress(t *testing.T) {
	response := `{
    "error": null,
    "id": 1,
    "result": "mwkpPfgSFj4fq2Xm96tUUGVPSAwhzWrXva"
}`
	testRPCCall(t, response, func(client *Client) {
		resp, err := client.GetNewAddress("")

		require.NoError(t, err)
		require.Equal(t, "mwkpPfgSFj4fq2Xm96tUUGVPSAwhzWrXva", resp.String())
	})
}

func TestClient_ListSinceBlock(t *testing.T) {
	response := `{
    "error": null,
    "id": 1,
    "result": {
        "lastblock": "0000000041aa2d2ceabb358401e072dbc8b8cc463e463a65cfeda6475cd60db4",
        "transactions": [
            {
                "account": "",
                "address": "mtHeXNNCuSotNyTqYCGvwtBmRp3MY2SyHT",
                "amount": 3e-05,
                "bip125-replaceable": "no",
                "blockhash": "00000000000021420990192c4e6143f51f024a6ae9b0312bb11119462fcbdebf",
                "blockindex": 29,
                "blocktime": 1506960823,
                "category": "receive",
                "confirmations": 8320,
                "label": "",
                "time": 1506960589,
                "timereceived": 1506960589,
                "txid": "8319d287855594e6a4e7fa17b9053922aa0b77d4176476ea238d9ef59ca1653c",
                "vout": 0,
                "walletconflicts": []
            }
        ]
    }
}`

	testRPCCall(t, response, func(client *Client) {
		hash, _ := blockchain.NewHashFromStr("00000000000021420990192c4e6143f51f024a6ae9b0312bb11119462fcbdebf")
		resp, err := client.ListSinceBlock(hash)

		require.NoError(t, err)
		require.Len(t, resp.Transactions, 1)
	})
}
