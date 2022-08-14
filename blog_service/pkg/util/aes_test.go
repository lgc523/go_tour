package util

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	SKey := "spiderspider5023" //DiRh2uiXRFW6SvB4S0kCHw
	encodingString := Encrypt("fuck", SKey)
	//decodingString := Decrypt(encodingString, SKey)
	fmt.Printf("AES-128-CBC\n加密：%s\n解密：%s\n", encodingString, "decodingString")

	//Zthe2zoUt7F2euPGQJoq5Q
}
