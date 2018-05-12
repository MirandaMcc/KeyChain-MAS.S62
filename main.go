package main

import (
	"fmt"
	//"math/rand"

	"github.com/btcsuite/btcd/chaincfg"
)

var (
	// we're running on testnet3
	testnet3Parameters = &chaincfg.TestNet3Params
)

//App Entrance
func main() {
	fmt.Printf("KeyChain\n")

	//1 - TODO - key fingerprint data from file or whatever
	testData := []float{12.3434, 15.9090, 10.43434, 0.0345}
	fingerprint := fingerprintConverter(testData)
	fmt.Println(fingerprint)
	// Call AddressFrom PrivateKey() to make a keypair
	addressTo, _ := GenerateAddress(fingerprint)

	//fmt.Printf("address is: %s\n", result)

	//Call EZTxBuilder to make a transaction
	//2 - TODO - get other transaction details from user input
	txFrom := "txid"
	addressFrom := "address"
	index := 0
	addressTo := "address"
	valueOut := 0
	optx := TxToHex(EZTxBuilder(txFrom, addressFrom, uint32(index), addressTo, fingerprint, int64(valueOut)))

	fmt.Printf("optx is: %s\n", optx)

	//3 - TODO - push to blockchain
	//You'll get a long hex string which you can test by running the transaction though bitcoin-cli's decoderawtransaction command `./bitcoin-cli decoderawtransaction (tx hex)`

	return
}
