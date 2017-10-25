package blockchain

type Chain struct {
	// Address encoding magics
	PubKeyHashAddrID byte // First byte of a P2PKH address
	ScriptHashAddrID byte // First byte of a P2SH address
}

var BitcoinMainChain = Chain{
	PubKeyHashAddrID: 0x00,
}

var BitcoinTestChain = Chain{
	PubKeyHashAddrID: 0x6f,
}
