package main

import (
	"fmt"
	"log"

	"github.com/ravikr88/mock-blockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	w := wallet.NewWallet()
	// fmt.Println(w.PrivateKeyStr())
	// fmt.Println(w.PublicKeyStr())

	fmt.Println(w.BlockchainAddress())

}
