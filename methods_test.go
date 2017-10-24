package btctools

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTestRPCServer(response string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, response)
	}))
	return ts
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

	ts := createTestRPCServer(response)
	defer ts.Close()

	client, _ := New(&ConnConfig{
		Host: ts.URL[7:],
	})

	blockChainInfo, err := client.GetBlockChainInfo()
	require.NoError(t, err)

	assert.EqualValues(t, blockChainInfo.BestBlockHash, "000000000505975c1a91cb553dd896e15f6ae8e110366fd1024efac9fa3bfa30")
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

	ts := createTestRPCServer(response)
	defer ts.Close()

	client, _ := New(&ConnConfig{
		Host: ts.URL[7:],
	})

	networkInfo, err := client.GetNetworkInfo()
	require.NoError(t, err)

	assert.EqualValues(t, networkInfo.Version, 140100)
}
