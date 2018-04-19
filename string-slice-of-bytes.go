package main

import "fmt"

func main() {
	// hexadecimal
	byteSliceHexa := []byte{0x43, 0x61, 0x66, 0xc3, 0xa9}
	strHexa := string(byteSliceHexa)
	fmt.Println(strHexa)

	// decimal
	byteSliceDeci := []byte{67, 97, 102, 195, 169}
	strDeci := string(byteSliceDeci)
	fmt.Println(strDeci)
}
