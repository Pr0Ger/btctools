package btctools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientNewID(t *testing.T) {
	c := Client{}

	id1 := c.nextID()
	id2 := c.nextID()

	assert.Equal(t, id1+1, id2)
}
