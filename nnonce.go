package nnonce

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type payload struct {
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	ID      int      `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
}

type response struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

// Helper
func hex2int(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return uint64(result)
}

// NextNonce will return next available nonce
func NextNonce(address, nodeURL string) (uint64, error) {
	data := payload{
		Method:  "parity_nextNonce",
		Params:  []string{address},
		ID:      1,
		Jsonrpc: "2.0",
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	reqBody := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", nodeURL, reqBody)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var responseObject response
	json.Unmarshal(resBody, &responseObject)

	nonce := hex2int(responseObject.Result)
	return nonce, nil
}
