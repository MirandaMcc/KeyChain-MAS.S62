package main

import (
	"fmt"
	//"math/rand"
	"github.com/btcsuite/btcd/chaincfg"
	//"bytes"
	//"encoding/binary"
	//"strconv"
	"strconv"
)

var (
	// we're running on testnet3
	testnet3Parameters = &chaincfg.TestNet3Params
)

//App Entrance
func main() {
	fmt.Printf("KeyChain\n")

	//1 - TODO - key fingerprint data from file or whatever
	testData := []float64{12.3434, 15.9090, 10.43434, 0.0345, 0.004, 0.132, 0.454, 34.343}
	//fingerprint := fingerprintConverter(testData)
	//fmt.Println(fingerprint)
	// Call AddressFrom PrivateKey() to make a keypair
	privateKey := "mas.s62"
	vault := Lock(privateKey, testData)

	vaultString := "["
	for _, vaultRow := range vault {

		vaultString += "[" + strconv.FormatFloat(vaultRow[0], 'E', -1, 64) + "," + strconv.FormatFloat(vaultRow[1], 'E', -1, 64) + "]"
		//fmt.Println(vaultString)
	}
	vaultString += "]"
	opReturnData := []byte(vaultString)

	publicAddress, _ := GenerateAddress(privateKey)

	fmt.Println("address is: %s\n", publicAddress)

	//Call EZTxBuilder to make a transaction
	//2 - TODO - get other transaction details from user input
	txFrom := "1f497ac245eb25cd94157c290f62d042e3bdda1e57920b6d1d2c5cfa362c12da"
	//addressFrom := "mpQQryVrYmGNPxVqNeE5RgoYAv2v66Psao"
	index := uint32(30)
	addressTo := "muNaPrVz8D2KcnjdQTZwFreKyw2ef8aDnA"
	valueOut := int64(10000)
	optx := TxToHex(OpReturnTxBuilder(opReturnData, txFrom, addressTo, valueOut, index, privateKey))

	fmt.Printf("optx is: %s\n", optx)

	//3 - TODO - push to blockchain
	//You'll get a long hex string which you can test by running the transaction though bitcoin-cli's decoderawtransaction command `./bitcoin-cli decoderawtransaction (tx hex)`

}
