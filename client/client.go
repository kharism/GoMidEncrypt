package client

import (
	"bufio"
	"fmt"
	"github/kharism/GoMidEncrypt/encryption"
	"net"
)

type Client struct {
	Encrypter      encryption.Encrypter
	Decrypter      encryption.Decrypter
	AddressListen  string
	AddressForward string
}

func handleConn(receiver, sender net.Conn, encrypter encryption.Encrypter, decrypter encryption.Decrypter) {
	//buffRead := make([]byte, 1024)
	// allRead := make([]byte, 1024)
	bufReceiver := bufio.NewReader(receiver)
	message, _ := bufReceiver.ReadBytes('\n')

	// //content := buffRead[:size]
	// allRead = append(allRead, content...)
	// for {

	// }

	fmt.Println("content", string(message))
	// encrypt content
	cipher := encrypter.Encrypt(message[:len(message)-1])
	fmt.Println(cipher)
	cipher = append(cipher, '\n')
	fmt.Println("cipher", cipher)
	_, err := sender.Write(cipher)
	if err != nil {
		fmt.Println("Error sending response", err.Error())
		return
	}
	fmt.Println("Waiting for response")
	//allRead = make([]byte, 1024)
	bufSender := bufio.NewReader(sender)
	allRead, _ := bufSender.ReadBytes('\n')
	fmt.Println("Done reading response")
	// for {
	// 	size, err := sender.Read(buffRead)
	// 	if err != nil && err != io.EOF {
	// 		fmt.Println("Error getting response", err.Error())
	// 		return
	// 	}
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	content := buffRead[:size]
	// 	allRead = append(allRead, content...)
	// }
	plainText := decrypter.Decrypt(allRead[:len(allRead)-1])

	fmt.Println("PlainText", string(plainText))
	receiver.Write(plainText)
	receiver.Write([]byte{'\n'})
	fmt.Println(">>>>>>>>>")
	receiver.Close()
}
func (c *Client) Start() error {
	listener, err := net.Listen("tcp", c.AddressListen)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		fmt.Println("Accepting connection")
		sender, err := net.Dial("tcp", c.AddressForward)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		go handleConn(conn, sender, c.Encrypter, c.Decrypter)
	}

}
