package btctools

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupClientTest() *Client {
	config := ConnConfig{
		Host: "127.0.0.1:18332",
		User: "docker",
		Pass: "docker",
	}
	client, _ := New(&config)

	return client
}

func TestClient_newID(t *testing.T) {
	client := Client{}

	id1 := client.nextID()
	id2 := client.nextID()

	assert.EqualValues(t, id1+1, id2)
}

func TestClient_generateRequest(t *testing.T) {
	tests := []struct {
		method string
		params interface{}
		out    []byte
	}{
		{
			method: "noParams",
			params: nil,
			out:    []byte("{\"method\":\"noParams\",\"params\":[],\"id\":1}"),
		},
		{
			method: "arrayParam",
			params: []string{"a", "b"},
			out:    []byte("{\"method\":\"arrayParam\",\"params\":[\"a\",\"b\"],\"id\":1}"),
		},
	}
	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			client := Client{}

			resp, err := client.generateRequest(test.method, test.params)

			require.NoError(t, err)
			require.EqualValues(t, resp, test.out)
		})
	}
}

//func (suite *ClientTestSuite) TestGetNetworkInfo() {
//	res, err := suite.client.GetNetworkInfo()
//
//	suite.NoError(err)
//	suite.EqualValues(res.Version, 140100)
//}
