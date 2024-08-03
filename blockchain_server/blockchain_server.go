package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ravikr88/mock-blockchain/block"
	"github.com/ravikr88/mock-blockchain/wallet"
)

var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

// BlockchainServer represents a simple blockchain server.
type BlockchainServer struct {
	port uint16
}

// NewBlockchainServer creates a new BlockchainServer instance with the specified port.
func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port: port}
}

// Port returns the port on which the BlockchainServer is running.
func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minersWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc
		log.Printf("private_key %v\n", minersWallet.PrivateKeyStr())
		log.Printf("public_key %v\n", minersWallet.PublicKey())
		log.Printf("blockchain_address %v\n", minersWallet.BlockchainAddress())
	}
	return bc
}

func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("Error: Invalid HTTP Method")
	}
}

// Run starts the HTTP server on the specified port.
func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	addr := "0.0.0.0:" + strconv.Itoa(int(bcs.Port()))
	log.Printf("Listening on http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
