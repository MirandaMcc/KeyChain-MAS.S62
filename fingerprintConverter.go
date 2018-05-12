package main

import (
	"encoding/binary"
	"fmt"
	"bytes"
)

func fingerprintConverter(data []float64) string {

	var stringRepresentationOfFingerPrint string
	for _, dataPoint := range data {

		buffer := new(bytes.Buffer)

		err := binary.Write(buffer, binary.BigEndian, dataPoint)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}

		snippet := string(buffer.Bytes())
		stringRepresentationOfFingerPrint += snippet
	}

	return stringRepresentationOfFingerPrint
}
