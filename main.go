package main

import (
	"fmt"

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
	string fingerprint = "abcdefghijklmnop"
	// Call AddressFrom PrivateKey() to make a keypair
	addressTo, _ := GenerateAddress(fingerprint)

	fmt.Printf("address is: %s\n", result)

	//Call EZTxBuilder to make a transaction
	//2 - TODO - get other transaction details from user input
	string txFrom := "txid"
	string addressFrom := "address"
	int index := 0
	string addressTo := "address" 
	int valueOut := 0
	optx := TxToHex(EZTxBuilder(txFrom, addressFrom, index, addressTo, fingerprint, valueOut))

	fmt.Printf("optx is: %s\n", optx)

	//3 - TODO - push to blockchain 
	//You'll get a long hex string which you can test by running the transaction though bitcoin-cli's decoderawtransaction command `./bitcoin-cli decoderawtransaction (tx hex)`


	return
}
