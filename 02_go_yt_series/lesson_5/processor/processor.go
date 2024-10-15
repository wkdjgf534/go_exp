package processor

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"

	"golang.org/x/crypto/argon2"
	"golang.org/x/term"
)

const fileName = "encrypted_data.bin"

type EncryptedPackage struct {
	Nonce         []byte
	Salt          []byte
	EncryptedData []byte
}

func MakeCrypterFrom(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return cipher.NewGCM(block) // Galois Counter Mode
}

func DeriveKeyFrom(passphrase, salt []byte) []byte {
	return argon2.IDKey([]byte(passphrase), salt, 1, 64*1024, 4, 32)
}

func GetPassphrase() ([]byte, error) {
	println("Enter passphrase:")
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("failed to grab passphrase: %w", err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	passphrase, err := term.ReadPassword(int(os.Stderr.Fd()))
	if err != nil {
		return nil, fmt.Errorf("failed to grab passphrase: %w", err)
	}

	return passphrase, nil
}
