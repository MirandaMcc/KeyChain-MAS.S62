package main

import (
	"encoding/binary"
	"fmt"
	"bytes"
	"reflect"
	"unsafe"
)

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func fingerprintConverter(data []float64) string {

	var stringRepresentationOfFingerPrint string
	for _, dataPoint := range data {

		buffer := new(bytes.Buffer)

		err := binary.Write(buffer, binary.BigEndian, dataPoint)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}

		snippet := BytesToString(buffer.Bytes())
		stringRepresentationOfFingerPrint += snippet
	}

	return stringRepresentationOfFingerPrint
}
