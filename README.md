## Golang Package for Parity Next Nonce

A simple Go package for [Parity Next Nonce](https://wiki.parity.io/JSONRPC-parity-module#parity_nextnonce)

### Installation

```
$ go get github.com/zulhfreelancer/go_parity_next_nonce
```

### Usage

```
func yourFunction() {
	// insert an address with some past transactions
	address := "0x..."

	// insert the Parity node RPC URL and port number
	nodeURL := "http://xx.xxx.xxx.xxx:8545"

	nonce, err := NextNonce(address, nodeURL)
	if err != nil {
		// handle error
	}

	fmt.Printf("Nonce : %v\n", nonce)
	fmt.Printf("Type  : %T\n", nonce)
}
```
