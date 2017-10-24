package btctools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sync/atomic"
)

// ConnConfig describes the connection configuration parameters for the client
type ConnConfig struct {
	// Host is the IP address and port of the RPC server you want to connect to
	Host string

	// User is the username to use to authenticate to the RPC server
	User string
	// Pass is the password to use to authenticate to the RPC server
	Pass string
}

// Client represents a Bitcoin RPC client which allows easy access to the
// various RPC methods available on a Bitcoin RPC server.  Each of the wrapper
// functions handle the details of converting the passed and return types to and
// from the underlying JSON types which are required for the JSON-RPC
// invocations
type Client struct {
	// config holds the connection configuration associated with this client
	config *ConnConfig

	// id for next RPC request
	id uint64

	httpClient *http.Client

	daemonType Forks
}

func (c *Client) nextID() uint64 {
	return atomic.AddUint64(&c.id, 1)
}

type clientRequest struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	ID     uint64        `json:"id"`
}

func (c *Client) generateRequest(method string, params interface{}) ([]byte, error) {
	req := clientRequest{
		Method: method,
		ID:     c.nextID(),
	}

	rt := reflect.ValueOf(params)
	switch rt.Kind() {
	case reflect.Slice:
		req.Params = make([]interface{}, rt.Len())
		for i := 0; i < rt.Len(); i++ {
			req.Params[i] = rt.Index(i).Interface()
		}
	default:
		if params != nil {
			req.Params = []interface{}{params}
		} else {
			req.Params = []interface{}{}
		}
	}

	return json.Marshal(req)
}

type clientResponse struct {
	ID     uint64           `json:"id"`
	Result *json.RawMessage `json:"result"`
	Error  interface{}      `json:"error"`
}

func (c *Client) sendRequest(method string) (*json.RawMessage, error) {
	fullURL := fmt.Sprintf("http://%v/", c.config.Host)

	body, err := c.generateRequest(method, nil)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fullURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(c.config.User, c.config.Pass)
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	var parsedResponse clientResponse
	err = dec.Decode(&parsedResponse)
	if err != nil {
		return nil, err
	}

	return parsedResponse.Result, nil
}

// New creates a new RPC client based on the provided connection configuration details
func New(config *ConnConfig) (*Client, error) {
	httpClient := &http.Client{}

	client := Client{
		config:     config,
		httpClient: httpClient,
	}
	return &client, nil
}
