package main

import (
	"github/kharism/GoMidEncrypt/client"
	"github/kharism/GoMidEncrypt/encryption"
)

func main() {
	client := client.Client{}
	client.AddressListen = ":8090"
	client.AddressForward = ":8088"
	aes := encryption.AESEncrypter{Key: "jreosoirppjreosoirppjreosoirppjj", IV: "928304kkrjtiqqdf"}
	client.Encrypter = &aes
	client.Decrypter = &aes
	client.Start()
}
