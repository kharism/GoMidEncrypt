package main_test

import (
	"bufio"
	"fmt"
	"github/kharism/GoMidEncrypt/encryption"
	"net"
	"testing"
)

func DummyServerStart(port string, encDec encryption.EncDec) {
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
		buff := make([]byte, 1024)
		size, err := conn.Read(buff)
		if err != nil {
			fmt.Println(err.Error())
		}
		content := buff[:size]
		plaintext := encDec.Decrypt(content)
		fmt.Println(plaintext)
		response := "hello " + string(plaintext)
		cipher := encDec.Encrypt([]byte(response))
		conn.Write(cipher)
	}
}
func TestClient(t *testing.T) {
	// client := client.Client{}
	// client.AddressListen = ":8090"
	// client.AddressForward = ":8088"
	// aes := encryption.AESEncrypter{Key: "jreosoirppjreosoirppjreosoirppjj", IV: "928304kkrjtiqqdf"}
	// client.Encrypter = &aes
	// client.Decrypter = &aes
	// go DummyServerStart(":8088", &aes)
	// go client.Start()
	conn, err := net.Dial("tcp", ":8090")
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	payload := []byte("udin\n")
	_, err = conn.Write(payload)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	t.Log("Done writting")
	// conn.Close()
	t.Log("Start Reading")
	buff := bufio.NewReader(conn)
	response, err := buff.ReadBytes('\n')
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	t.Log(response)
	t.Log([]byte("hello udin"))
	if string(response[:len(response)-1]) != "hello udin" {
		t.Log("Tidak sama", response, "hello udin")
		t.FailNow()
	}
	conn.Close()
}
