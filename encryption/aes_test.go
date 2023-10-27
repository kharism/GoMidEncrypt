package encryption_test

import (
	"github/kharism/GoMidEncrypt/encryption"
	"testing"
)

func TestAes(t *testing.T) {
	key := "12345678901234567890123456789012"
	iv := "1234567890123456"
	plaintext := "abcdefghijklmnopqrstuvwxyzABCDEF"
	aes := encryption.AESEncrypter{Key: key, IV: iv}
	ciphertext := aes.Encrypt([]byte(plaintext))
	decrypted := aes.Decrypt(ciphertext)
	// removePadding := bytes.TrimRight(decrypted, " ")
	if string(decrypted) != plaintext {
		t.Log(decrypted)
		t.Log([]byte(plaintext))
		t.Log(string(decrypted) != plaintext)
		t.Fail()
	}

	plaintext = "abcde"
	ciphertext = aes.Encrypt([]byte(plaintext))
	decrypted = aes.Decrypt(ciphertext)
	// removePadding := bytes.TrimRight(decrypted, " ")
	if string(decrypted) != plaintext {
		t.Log(decrypted)
		t.Log([]byte(plaintext))
		t.Log(string(decrypted) != plaintext)
		t.Fail()
	}
}

func TestRemovePadding(t *testing.T) {
	input := [][]byte{
		[]byte{0x01, 0x03, 0x03, 0x03},
		[]byte{0x01, 0x03, 0x02, 0x03},
		[]byte{0x01, 0x02, 0x03, 0x03},
	}
	output := [][]byte{
		[]byte{0x1},
		[]byte{0x01, 0x03, 0x02, 0x03},
		[]byte{0x01, 0x02, 0x03, 0x03},
	}
	for idx := range input {
		tt := encryption.RemovePKCS5Padding(input[idx])
		if string(output[idx]) != string(tt) {
			print(output[idx])
			print(tt)
			t.FailNow()
		}
	}

}
