package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

//testo.go - used to quickly decrypt / encrypt packet json data without running the entire bot app


var (
	dec         = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	encryptKey  = "E5QRecZA"
	decryptKey  = "DvNw3mJT"
	ndec        = "\x44\x76\x4e\x77\x33\x6d\x4a\x54"
	bEncryptKey = []byte{0x45, 0x35, 0x51, 0x52, 0x65, 0x63, 0x5a, 0x41}
	bDecryptKey = []byte{0x44, 0x76, 0x4e, 0x77, 0x33, 0x6d, 0x4a, 0x54}
)

func DesEncryption(key, iv, plainText []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData := PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}

func DesDecryption(skey, siv, scipherText string) ([]byte, error) {
	key := []byte(skey)
	iv := []byte(siv)
	cipherText := []byte(scipherText)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = PKCS5UnPadding(origData)
	//origData = pkcs7Unpad(origData, block.BlockSize())
	return origData, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func main() {
	originalText := "1rAS2W+sTlxAg5zePp/nyCYDvNGs7Zne6YZYednVCA02H7ik0YmpAiO1hy7TurudP6VAN1NBDD7KXElblJgqti9iHhb7kZSEgHAs3cgVs1Pcp7X0KGIFn344c473Lcbzv+1GV26nbZeS46uPgn6VGQwNSxqeDCnLqzy1xd/SISjG0r9TV/PB4NCerGvl3UwOVVXf3WwXdLlSCziHiPYLhtS9/uBYHME8gAKKFdkwkNd46yhl+ZkDYH2MfH7JCHq7vzdfKhUv6RBhU1wemT7N7UlzzWCbZ2C6UhsDt7TNGgUiCfKawiI/"
	x := make([]byte, len(originalText))
	x, _ = base64.StdEncoding.DecodeString(originalText)
	key := encryptKey
	iv := encryptKey
	cryptoText, _ := DesDecryption(key, iv, string(x))
	fmt.Println(string(cryptoText))
	fmt.Println(cryptoText)
}

// func main() {
// 	originalText := `{"Email":"<REDACTED>@gmail.com","Password":"<REDACTED>","FirstName":"<REDACTED>","LastName":"<REDACTED>","Device":{"DeviceToken":"<REDACTED>","OperatingSystem":0,"DeviceVersion":"iPod Touch 6G"}}`
// 	//x := make([]byte, len(originalText))
// 	xl := []byte(originalText)
// 	//_, _ = base64.StdEncoding.Decode(x, []byte(originalText))
// 	fmt.Println(originalText)
// 	key := bEncryptKey
// 	iv := bEncryptKey
// 	cipherText, _ := DesEncryption(key, iv, xl)
// 	x := make([]byte, len(cipherText)*2)
// 	base64.StdEncoding.Encode(x, cipherText)
// 	fmt.Println(string(x))
// }
