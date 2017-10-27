package blockchain

import (
	"encoding/hex"
	"fmt"
)

// HashSize of array used to store hashes.  See Hash.
const HashSize = 32

// MaxHashStringSize is the maximum length of a Hash hash string.
const MaxHashStringSize = HashSize * 2

// BlockHash is used for storing hash of block
type BlockHash [HashSize]byte

// UnmarshalJSON implements the json.Unmarshaler interface
func (hash *BlockHash) UnmarshalJSON(src []byte) error {
	if len(src) > MaxHashStringSize+2 {
		return fmt.Errorf("max hash string length is %v bytes", MaxHashStringSize)
	}

	src = src[1 : len(src)-1]

	var reversedHash BlockHash
	_, err := hex.Decode(reversedHash[HashSize-hex.DecodedLen(len(src)):], src)
	if err != nil {
		return err
	}

	for i, b := range reversedHash[:HashSize/2] {
		hash[i], hash[HashSize-1-i] = reversedHash[HashSize-1-i], b
	}

	return nil
}

// String returns the Hash as the hexadecimal string of the byte-reversed hash
func (hash BlockHash) String() string {
	for i := 0; i < HashSize/2; i++ {
		hash[i], hash[HashSize-1-i] = hash[HashSize-1-i], hash[i]
	}
	return hex.EncodeToString(hash[:])
}

func NewHashFromStr(str string) (*BlockHash, error) {
	if len(str) > MaxHashStringSize {
		return nil, fmt.Errorf("max hash string length is %v bytes", MaxHashStringSize)
	}

	ret := new(BlockHash)

	var srcBytes []byte
	if len(str)%2 == 0 {
		srcBytes = []byte(str)
	} else {
		srcBytes = make([]byte, 1+len(str))
		srcBytes[0] = '0'
		copy(srcBytes[1:], str)
	}

	var reversedHash BlockHash
	_, err := hex.Decode(reversedHash[HashSize-hex.DecodedLen(len(srcBytes)):], srcBytes)
	if err != nil {
		return nil, err
	}

	for i, b := range reversedHash[:HashSize/2] {
		ret[i], ret[HashSize-1-i] = reversedHash[HashSize-1-i], b
	}

	return ret, nil
}
