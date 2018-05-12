package main

import (
	"encoding/binary"
	"fmt"
)

func fingerprintConverter(data []float) string {

	var stringRepresentationOfFingerPrint string
	for _, dataPoint := range data {

		var buffer bytes.Buffer

		err := binary.Write(&buffer, binary.BigEndian, dataPoint)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}

		snippet := string(buffer.Bytes())
		stringRepresentationOfFingerPrint += snippet
	}

	return stringRepresentationOfFingerPrint
}
