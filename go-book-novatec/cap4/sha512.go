package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	var shType int
	sh256 := sha256.Sum256([]byte("x"))
	fmt.Printf("sha256: %x\n", sh256)

	fmt.Println("1 - Para SHA384")
	fmt.Println("2 - Para SHA512")
	fmt.Scanf("%d", &shType)
	switch shType {
	case 1:
		sh384 := sha512.Sum384([]byte(""))
		fmt.Printf("sha512: %x\n", sh384)
	case 2:
		sh512 := sha512.Sum512([]byte(""))
		fmt.Printf("sha512: %x\n", sh512)
	}

}
