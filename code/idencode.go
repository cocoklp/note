package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

// EncryptDecrypt runs a XOR encryption on the input string,
// encrypting it if it hasn't already been,
// and decrypting it if it has, using the key provided.

var defaultKey string = "test-uuid"
var idPrefix string = ""

// https://haobook.readthedocs.io/zh_CN/latest/periodical/201608/ligang.html
// https://github.com/KyleBanks/XOREncryption/blob/master/Go/xor.go
func EncryptDecrypt(input, key string) (output string) {
	for i, _ := range input {
		fmt.Printf("[%s]\n", string(input[i]^key[i%len(key)]))
		output += string(input[i] ^ key[i%len(key)])
	}
	return output
}

func Transform(input string) string {
	return EncryptDecrypt(input, defaultKey)
}

// use base64 encode
func Encrypt(input int64) string {
	enc := Transform(idPrefix + strconv.FormatInt(input, 10))
	fmt.Println("~~~~~~", enc)
	return base64.StdEncoding.EncodeToString([]byte(enc))
}

func Decrypt(encodedId string) (int64, error) {
	enc, err := base64.StdEncoding.DecodeString(encodedId)
	if err != nil {
		return -1, err
	}
	fmt.Println("~~~", string(enc))
	idStrWithPrefix := Transform(string(enc))
	if !strings.HasPrefix(idStrWithPrefix, idPrefix) {
		return -1, fmt.Errorf("fake id")
	}
	idStr := strings.TrimPrefix(idStrWithPrefix, idPrefix)
	return strconv.ParseInt(idStr, 10, 64)
}

func main() {
	// ===========Test Encrypting  decrypting id==============
	fmt.Println("// 3 ===========Test Encrypting  decrypting id")
	var id int64
	id = 278
	enc := Encrypt(id)
	fmt.Printf("[%d] --Encrypt to--> [%s]\n", id, enc)

	decId, err := Decrypt(enc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[%s] --Decrypt to--> [%d]\n", enc, decId)
	if decId != id {
		fmt.Println("decId != id")
	}

	// =========== try base64 id str attack
	idbase64 := base64.StdEncoding.EncodeToString([]byte("9123456789"))
	decId, err = Decrypt(idbase64)
	fmt.Printf("[%d]  [%v]\n", decId, err)
}
