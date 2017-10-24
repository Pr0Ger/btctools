package btctools

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	client, _ := New(&ConnConfig{
		Host: ts.URL[7:],
	})

	networkInfo, err := client.GetNetworkInfo()
	require.NoError(t, err)

	assert.EqualValues(t, networkInfo.Version, 140100)
}
