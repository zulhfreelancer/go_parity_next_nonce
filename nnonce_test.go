package nnonce

import (
	"fmt"
	"testing"
)

func TestNextNonce(t *testing.T) {
	// insert an address with some past transactions
	address := "0x..."

	// insert the node RPC URL and port number
	nodeURL := "http://xx.xxx.xxx.xxx:8545"

	nonce, err := NextNonce(address, nodeURL)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Nonce : %v\n", nonce)
	fmt.Printf("Type  : %T\n", nonce)
}

/* ---------------------------------------------------------------
+++++++++++++++++++++++ Example output: ++++++++++++++++++++++++++

$ go test
Nonce : 193
Type  : uint64
PASS
ok  	github.com/zulhfreelancer/go_parity_next_nonce	0.036s
--------------------------------------------------------------- */
