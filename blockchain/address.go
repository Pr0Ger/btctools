package blockchain

import (
	"errors"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

var (
	ErrChecksumMismatch = errors.New("checksum mismatch")
)

// encodeAddress returns a human-readable payment address given a ripemd160 hash
// and netID which encodes the bitcoin network and address type.  It is used
// in both pay-to-pubkey-hash (P2PKH) and pay-to-script-hash (P2SH) address
// encoding.
func encodeAddress(hash160 []byte, netID byte) string {
	// Format is 1 byte for a network and address class (i.e. P2PKH vs
	// P2SH), 20 bytes for a RIPEMD160 hash, and 4 bytes of checksum.
	return base58.CheckEncode(hash160[:ripemd160.Size], netID)
}

type Address interface {
	EncodeAddress() string
	String() string
	UnmarshalJSON([]byte) error
}

type AddressPubKeyHash struct {
	hash  [ripemd160.Size]byte
	netID byte
}

func (a *AddressPubKeyHash) EncodeAddress() string {
	return encodeAddress(a.hash[:], a.netID)
}

// String returns human-readable string for using with fmt.Stringer
func (a *AddressPubKeyHash) String() string {
	return a.EncodeAddress()
}

func (a *AddressPubKeyHash) UnmarshalJSON(src []byte) error {
	src = src[1 : len(src)-1]

	decoded, netID, err := base58.CheckDecode(string(src))
	if err != nil {
		if err == base58.ErrChecksum {
			return ErrChecksumMismatch
		}
		return errors.New("decoded address is of unknown format")
	}
	if len(decoded) != ripemd160.Size {
		return errors.New("decoded address is of unknown size")
	}
	a.netID = netID
	copy(a.hash[:], decoded)

	return nil
}

// DecodeAddress will attempt to recognize and parse address from string
func DecodeAddress(addr string) (Address, error) {
	decoded, netID, err := base58.CheckDecode(addr)
	if err != nil {
		if err == base58.ErrChecksum {
			return nil, ErrChecksumMismatch
		}
		return nil, errors.New("decoded address is of unknown format")
	}
	switch len(decoded) {
	case ripemd160.Size: // P2PKH or P2SH
		addr := AddressPubKeyHash{netID: netID}
		copy(addr.hash[:], decoded)
		return &addr, nil
	default:
		return nil, errors.New("decoded address is of unknown size")
	}
}
