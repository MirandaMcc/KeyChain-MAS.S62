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
	//"encoding/json"
	"io/ioutil"
	"math"
	"strconv"
	//"strings"
	"encoding/json"
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
	privateKey := "WHAT'S UP?"
	vault := Lock(privateKey, testData)
	fmt.Println("vault:", vault)
	//coeffs := Unlock(testData, vault)
	//fmt.Println("Decoded Coeffs:", coeffs)
	//ret := Decode(coeffs)
	//fmt.Println("Decoded String:", ret)

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

	var compressedVaultPieces [][]byte
	for pieceIndex := 0; pieceIndex < len(compressedOpReturnDataInByte); pieceIndex += 520 {
		piece := compressedOpReturnDataInByte[pieceIndex:int(math.Min(float64(pieceIndex+520), float64(len(compressedOpReturnDataInByte))))]
		doubleByteOfPiece := [][]byte{piece}
		compressedVaultPieces = append(compressedVaultPieces, doubleByteOfPiece...)
		//fmt.Println("section: ", vaultPieces)
	}

	//fmt.Println("Split vault: ", compressedVaultPieces[1])
	publicAddress, _ := GenerateAddress("KeyChain")

	fmt.Println("address is: ", publicAddress)

	//Call EZTxBuilder to make a transaction
	//2 - TODO - get other transaction details from user input
	txFrom := "1f497ac245eb25cd94157c290f62d042e3bdda1e57920b6d1d2c5cfa362c12da"
	//addressFrom := "mpQQryVrYmGNPxVqNeE5RgoYAv2v66Psao"
	index := uint32(32)
	addressTo := publicAddress
	valueOut := int64(10000)
	var transactionStrings []string
	for _, compressedVaultPiece := range compressedVaultPieces {
		optx := OpReturnTxBuilder(compressedVaultPiece, txFrom, addressTo, valueOut, index, privateKey)
		hexOpt := TxToHex(optx)
		transactionStrings = append(transactionStrings, hexOpt)
	}

	fmt.Println("Transaction 0: ", transactionStrings[0])
	fmt.Println("Transaction 1: ", transactionStrings[1])
	//	fmt.Println("optx is: ", optx)
	//	fmt.Println("hexopt: ", hexOpt)

	retrievedCompressedVaultPieces := append(compressedVaultPieces[0], compressedVaultPieces[1]...)
	decompressedOpReturnData0, _ := base64.StdEncoding.DecodeString(string(retrievedCompressedVaultPieces))
	fmt.Println("decompressedOpReturnData0: ",decompressedOpReturnData0)
	rdata0 := bytes.NewReader(decompressedOpReturnData0)
	fmt.Println("rdata0: ",rdata0)
	r0, _ := gzip.NewReader(rdata0)
	fmt.Println("r0: ",r0)
	decodedVaultPiece, _ := (ioutil.ReadAll(r0))
	fmt.Println("decodedVaultPiece: ",decodedVaultPiece)
	decodedVaultString := string(decodedVaultPiece)
	fmt.Println("decodedVaultString: ",decodedVaultString)

	//decompressedOpReturnData1, _ := base64.StdEncoding.DecodeString(string(compressedVaultPieces[1]))
	//fmt.Println("decompressedOpReturnData1: ",decompressedOpReturnData1)
	//rdata1 := bytes.NewReader(decompressedOpReturnData1)
	//fmt.Println("rdata1: ",rdata1)
	//r1, _ := gzip.NewReader(rdata1)
	//fmt.Println("r1: ",r1)
	//pieceDecodedVaultString, _ := (ioutil.ReadAll(r1))
	//fmt.Println("pieceDecodedVaultString: ",pieceDecodedVaultString)
	//decodedVaultString += string(pieceDecodedVaultString)
	//fmt.Println("Decoded string: ", decodedVaultString)

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
	return
}
