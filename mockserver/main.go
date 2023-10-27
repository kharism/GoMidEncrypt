package main

import (
	"bufio"
	"fmt"
	"github/kharism/GoMidEncrypt/encryption"
	"net"
)

func DummyServerStart(port string, realAddress string, encDec encryption.EncDec) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err.Error(), port)
	}
	//aes := encryption.AESEncrypter{Key: "jreosoirppjreosoirppjreosoirppjj",IV:"928304kkrjtiqqdf"}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go func(conn net.Conn) {
			fmt.Println("Accepting connection")
			bufReader := bufio.NewReader(conn)
			// buff := make([]byte, 1024)
			fmt.Println("Start Reading")
			// size, err := conn.Read(buff)
			// if err != nil {
			// 	fmt.Println(err.Error())
			// }
			// content := buff[:size]
			content, _ := bufReader.ReadBytes('\n')
			// fmt.Println("Decrypting stuff")
			// fmt.Println(content)
			fmt.Println("cipher", content)
			plaintext := encDec.Decrypt(content[:len(content)-1])
			// plaintext = bytes.TrimRight(plaintext, " ")
			fmt.Println(string(plaintext))
			fmt.Println("Encrypting stuff")
			//sending to actual server

			connReal, _ := net.Dial("tcp", realAddress)
			bufWriter := bufio.NewWriter(connReal)
			bufWriter.Write(plaintext)
			bufReaderReal := bufio.NewReader(connReal)
			response, _ := bufReaderReal.ReadString('\n')
			//
			cipher := encDec.Encrypt([]byte(response))
			fmt.Println("writing response")
			cipher = append(cipher, '\n')
			conn.Write(cipher)
			fmt.Println("Done writting response")
			fmt.Println(">>>>>>>>")
			conn.Close()
		}(conn)

	}
}
func main() {
	aes := encryption.AESEncrypter{Key: "jreosoirppjreosoirppjreosoirppjj", IV: "928304kkrjtiqqdf"}
	DummyServerStart(":8088", ":8888", &aes)
}
