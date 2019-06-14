## Golang Package for Parity Next Nonce

A simple Go package for [Parity Next Nonce](https://wiki.parity.io/JSONRPC-parity-module#parity_nextnonce)

### Installation

```
// If you are not using Go 1.11 Modules
$ go get github.com/zulhfreelancer/go_parity_next_nonce

// If you are using Go 1.11 Modules, run this before executing the code below
$ go mod tidy
```

### Usage

```
import (
	// Short version (implicit)
	"github.com/zulhfreelancer/go_parity_next_nonce"

	// If you prefer to make it explicit
	nnonce "github.com/zulhfreelancer/go_parity_next_nonce"
)

func yourFunction() {
	// insert an address with some past transactions
	address := "0x..."

	// insert the Parity node RPC URL and port number
	nodeURL := "http://xx.xxx.xxx.xxx:8545"

	nonce, err := nnonce.NextNonce(address, nodeURL)
	if err != nil {
		// handle error
	}

	fmt.Printf("Nonce : %v\n", nonce) // 193
	fmt.Printf("Type  : %T\n", nonce) // uint64
}
```
