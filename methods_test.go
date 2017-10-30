package btctools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Pr0Ger/btctools/blockchain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type parsedJSON map[string]interface{}

func emptyValidator(*testing.T, *parsedJSON) {}

func testMethod(t *testing.T, method string, validator func(*testing.T, *parsedJSON), tester func(*testing.T, *Client, *parsedJSON)) {
	t.Helper()

	requests, err := loadResource(method, "request")
	require.NoError(t, err)

	responses, err := loadResource(method, "response")
	require.NoError(t, err)

	params, err := loadResource(method, "params")
	if err != nil {
		params = &parsedJSON{
			"all": []parsedJSON{},
		}
	}

	currencies := make([]string, 0, len(*responses))
	for key := range *responses {
		if key != "all" {
			currencies = append(currencies, key)
		}
	}

	getDataForCurrency := func(resource *parsedJSON, currency string) []parsedJSON {
		var data []interface{}
		switch required := (*resource)[currency].(type) {
		case []interface{}:
			data = required
		default:
			switch fallback := (*resource)["all"].(type) {
			case []interface{}:
				data = fallback
			default:
				return []parsedJSON{{}}
			}
		}
		res := make([]parsedJSON, len(data))
		for i, val := range data {
			res[i] = val.(map[string]interface{})
		}
		return res
	}

	if len(currencies) == 0 {
		testMethodForCurrency(t, getDataForCurrency(requests, "all"), getDataForCurrency(responses, "all"),
			getDataForCurrency(params, "all"), validator, tester)
	} else {
		for _, currency := range currencies {
			t.Run(currency, func(t *testing.T) {
				testMethodForCurrency(t, getDataForCurrency(requests, currency), getDataForCurrency(responses, currency),
					getDataForCurrency(params, currency), validator, tester)
			})
		}
	}
}

func testMethodForCurrency(t *testing.T, requests []parsedJSON, responses []parsedJSON, params []parsedJSON,
	validator func(*testing.T, *parsedJSON), tester func(*testing.T, *Client, *parsedJSON)) {
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		request := new(parsedJSON)
		require.NoError(t, json.Unmarshal(body, &request))

		expectedRequest := requests[0]

		for key, value := range expectedRequest {
			require.Equal(t, value, (*request)[key])
		}

		validator(t, request)

		w.Header().Set("Content-Type", "application/json")

		data, err := json.Marshal(responses[0])
		require.NoError(t, err)
		w.Write(data)
	}))
	defer ts.Close()

	client, _ := New(&ConnConfig{
		Host: ts.URL[7:],
	})

	tester(t, client, &params[0])
}

func loadResource(method string, resType string) (*parsedJSON, error) {
	path := fmt.Sprintf("./rpc_test_mocks/%v/%v.json", method, resType)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read response for method: %v", method)
	}

	parsedFile := new(parsedJSON)
	err = json.Unmarshal(data, parsedFile)
	return parsedFile, nil
}

func TestClient_GetBlockChainInfo(t *testing.T) {
	testMethod(t, "GetBlockChainInfo", emptyValidator, func(t *testing.T, client *Client, params *parsedJSON) {
		blockChainInfo, err := client.GetBlockChainInfo()
		require.NoError(t, err)

		assert.EqualValues(t, (*params)["bestblockhash"], blockChainInfo.BestBlockHash)
	})
}

func TestClient_GetBlockHeader(t *testing.T) {
	testMethod(t, "GetBlockHeader", emptyValidator, func(t *testing.T, client *Client, params *parsedJSON) {
		hash, _ := blockchain.NewHashFromStr((*params)["block"].(string))
		blockHeader, err := client.GetBlockHeader(hash)

		require.NoError(t, err)
		assert.EqualValues(t, (*params)["confirmations"], blockHeader.Confirmations)
	})
}

func TestClient_GetNetworkInfo(t *testing.T) {
	testMethod(t, "GetNetworkInfo", emptyValidator, func(t *testing.T, client *Client, params *parsedJSON) {
		networkInfo, err := client.GetNetworkInfo()
		require.NoError(t, err)

		require.NotNil(t, params)
		assert.EqualValues(t, networkInfo.Version, (*params)["networkVersion"])
	})
}

func TestClient_GetNewAddress(t *testing.T) {
	testMethod(t, "GetNewAddress", emptyValidator, func(t *testing.T, client *Client, params *parsedJSON) {
		resp, err := client.GetNewAddress("")
		require.NoError(t, err)

		require.NotNil(t, params)
		assert.Equal(t, (*params)["address"], resp.String())
	})
}

func TestClient_ListSinceBlock(t *testing.T) {
	testMethod(t, "ListSinceBlock", emptyValidator, func(t *testing.T, client *Client, params *parsedJSON) {
		hash, _ := blockchain.NewHashFromStr((*params)["block"].(string))
		resp, err := client.ListSinceBlock(hash)

		require.NoError(t, err)
		assert.Len(t, resp.Transactions, 1)
	})
}

func TestClient_SendToAddress(t *testing.T) {
	testMethod(t, "SendToAddress", emptyValidator, func(t *testing.T, client *Client, params *parsedJSON) {
		addr, _ := blockchain.DecodeAddress((*params)["addr"].(string))

		txID, err := client.SendToAddress(addr, 1, "", "", true)

		require.NoError(t, err)
		assert.Equal(t, (*params)["tx"].(string), txID)
	})
}
