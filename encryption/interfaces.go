package encryption

type Encrypter interface {
	Encrypt(plaintext []byte) []byte
}

type Decrypter interface {
	Decrypt(ciphertext []byte) []byte
}
type EncDec interface {
	Encrypter
	Decrypter
}
