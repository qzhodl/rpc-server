package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type MyCustomAPI struct{}

func (api *MyCustomAPI) GetBlob(kvIndex int64, blobHash common.Hash, off, len int64) (hexutil.Bytes, error) {
	if off == 0 {
		return []byte{}, errors.New("there is no blob data")
	}
	return bytes.Repeat([]byte{11}, int(len)), nil // Replace with actual implementation
}


func main() {
	server := rpc.NewServer()

	server.RegisterName("es", new(MyCustomAPI))
	// Define the HTTP handler for the JSON-RPC server
	http.Handle("/", server)

	addr := ":8546"
	fmt.Printf("JSON-RPC server listening on %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
