package main

import (
	"fmt"
	//"math/rand"
	"github.com/btcsuite/btcd/chaincfg"
	//"bytes"
	//"encoding/binary"
	//"strconv"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
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
	for index, vaultRow := range vault {
		vaultString += "[" + strconv.FormatFloat(vaultRow[0], 'E', -1, 64) + "," + strconv.FormatFloat(vaultRow[1], 'E', -1, 64) + "]"
		if index < len((vault))-1 {
			vaultString += ","
		}

		//fmt.Println(vaultString)
	}
	vaultString += "]"
	opReturnData := []byte(vaultString)

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(opReturnData); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	compressedOpReturnData := base64.StdEncoding.EncodeToString(b.Bytes())
	fmt.Println("compressed String: ", compressedOpReturnData)
	compressedOpReturnDataInByte := []byte(compressedOpReturnData)

	fmt.Println("Length of compressed vault in bytes: ", len(compressedOpReturnDataInByte))
	//numberOfSplits := math.Ceil(float64(len(compressedOpReturnDataInByte)) / float64(520.0))

	var vaultPieces []byte
	for pieceIndex := 0; pieceIndex < len(compressedOpReturnDataInByte); pieceIndex += 520 {
		vaultPieces = append(vaultPieces, compressedOpReturnDataInByte[pieceIndex:int(math.Min(float64(pieceIndex+520), float64(len(compressedOpReturnDataInByte))))]...)
	}

	fmt.Println("Split vault: ", vaultPieces[0])
	publicAddress, _ := GenerateAddress(privateKey)

	fmt.Println("address is: %s\n", publicAddress)

	//Call EZTxBuilder to make a transaction
	//2 - TODO - get other transaction details from user input
	txFrom := "1f497ac245eb25cd94157c290f62d042e3bdda1e57920b6d1d2c5cfa362c12da"
	//addressFrom := "mpQQryVrYmGNPxVqNeE5RgoYAv2v66Psao"
	index := uint32(30)
	addressTo := "muNaPrVz8D2KcnjdQTZwFreKyw2ef8aDnA"
	valueOut := int64(10000)
	optx := TxToHex(OpReturnTxBuilder([]byte(compressedOpReturnData), txFrom, addressTo, valueOut, index, privateKey))
	fmt.Printf("optx is: %s\n", optx)

	decompressedOpReturnData, _ := base64.StdEncoding.DecodeString(compressedOpReturnData)
	fmt.Println("Decoded string: ", decompressedOpReturnData)
	rdata := bytes.NewReader(decompressedOpReturnData)
	r, _ := gzip.NewReader(rdata)
	decodedVaultString, _ := (ioutil.ReadAll(r))
	fmt.Println("Decoded stuff: ", string(decodedVaultString))

	var vaultArray [][]float64
	dec := json.NewDecoder(strings.NewReader(string(decodedVaultString)))
	err := dec.Decode(&vaultArray)
	fmt.Println(err, vaultArray)
	fmt.Println("vaultArray: ", vaultArray)

	coeffs := Unlock(testData, vaultArray)
	decodedPrivateKey := Decode(coeffs)
	fmt.Println("coefficients: ", coeffs)
	fmt.Println("decoded private key: ", decodedPrivateKey)

	//3 - TODO - push to blockchain
	//You'll get a long hex string which you can test by running the transaction though bitcoin-cli's decoderawtransaction command `./bitcoin-cli decoderawtransaction (tx hex)`

}
