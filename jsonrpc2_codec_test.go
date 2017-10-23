package btctools

import (
	"io"
	"net/rpc"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testReadWriteCloser struct {
	mock.Mock
	io.ReadWriteCloser
}

func (m *testReadWriteCloser) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *testReadWriteCloser) Write(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *testReadWriteCloser) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestClientCodec_WriteRequestWithoutParams(t *testing.T) {
	conn := new(testReadWriteCloser)
	codec := newClientCodec(conn)

	req := rpc.Request{
		ServiceMethod: "methodName",
		Seq:           1,
	}

	result := "{\"method\":\"methodName\",\"params\":[null],\"id\":1}\n"
	conn.On("Write", []byte(result)).Return(47, nil)

	codec.WriteRequest(&req, nil)

	conn.AssertExpectations(t)
}

func TestClientCodec_WriteRequestWithParamsArray(t *testing.T) {
	conn := new(testReadWriteCloser)
	codec := newClientCodec(conn)
	req := rpc.Request{
		ServiceMethod: "methodName",
		Seq:           1,
	}

	result := "{\"method\":\"methodName\",\"params\":[\"a\",\"b\"],\"id\":1}\n"
	conn.On("Write", []byte(result)).Return(50, nil)

	codec.WriteRequest(&req, []string{"a", "b"})

	conn.AssertExpectations(t)
}

func TestClientCodec_ReadResponse(t *testing.T) {
	conn := new(testReadWriteCloser)
	codec := newClientCodec(conn)

	network_response := "{\"error\": null, \"id\": 1, \"result\": \"asd\"}\n"
	conn.On("Read", mock.AnythingOfType("[]uint8")).
		Return(42, nil).
		Run(func(arguments mock.Arguments) {
			arg := arguments.Get(0).([]byte)
			copy(arg, []byte(network_response))
		})

	var resp rpc.Response
	codec.ReadResponseHeader(&resp)

	assert.Equal(t, resp.Error, "")
	assert.Equal(t, resp.Seq, uint64(1))

	var response string
	codec.ReadResponseBody(&response)

	assert.Equal(t, response, "asd")

	conn.AssertExpectations(t)
}

func TestClientCodec_Close(t *testing.T) {
	conn := new(testReadWriteCloser)
	codec := newClientCodec(conn)

	conn.On("Close").Return(nil)

	codec.Close()

	conn.AssertExpectations(t)
}
