package btctools

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestClient_sendAuthHeader(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		require.Equal(t, auth, `Basic dGVzdDp0ZXN0`)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{}`)
	}))
	defer ts.Close()

	client, _ := New(&ConnConfig{
		Host: ts.URL[7:],
		User: "test",
		Pass: "test",
	})

	client.sendRequest("ping")
}
