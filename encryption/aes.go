package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type AESEncrypter struct {
	Key string
	IV  string
}

func Aes256(plaintext []byte, key string, iv string, blockSize int) []byte {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding(plaintext, blockSize, len(plaintext))
	block, _ := aes.NewCipher(bKey)
	fmt.Println(block.BlockSize())
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return ciphertext
}
func RemovePKCS5Padding(data []byte) []byte {
	lastByte := int(data[len(data)-1])
	possiblePadding := data[len(data)-lastByte:]
	for idx := range possiblePadding {
		if possiblePadding[idx] != byte(lastByte) {
			return data
		}
	}
	return data[:len(data)-lastByte]
}
func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func (a *AESEncrypter) Encrypt(plaintext []byte) []byte {
	return Aes256(plaintext, a.Key, a.IV, 64)
}
func (a *AESEncrypter) Decrypt(ciphertext []byte) []byte {
	bKey := []byte(a.Key)
	bIV := []byte(a.IV)
	cipherTextDecoded := ciphertext

	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}
	// padding := byte(64 - len(ciphertext)%64)

	mode := cipher.NewCBCDecrypter(block, bIV)
	fmt.Println(len(cipherTextDecoded), block.BlockSize())
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	// paddingByte := cipherTextDecoded[len(cipherTextDecoded)-1]
	fmt.Println("still have padding", string(cipherTextDecoded))
	// fmt.Println("Padding", paddingByte)
	cipherTextDecoded = RemovePKCS5Padding(cipherTextDecoded) //bytes.TrimRight(cipherTextDecoded, string([]byte{paddingByte}))
	return cipherTextDecoded
}
