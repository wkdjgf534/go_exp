package processor

import (
	"bytes"
	"encoding/gob"
	"os"
)

func Decrypt() ([]byte, error) {
	pkg, err := readEncryptedFile(fileName)
	if err != nil {
		return nil, err
	}

	nonce := pkg.Nonce
	salt := pkg.Salt
	encryptedData := pkg.EncryptedData

	passphrase, err := GetPassphrase()
	if err != nil {
		return nil, err
	}

	key := DeriveKeyFrom(passphrase, salt)

	crypter, err := MakeCrypterFrom(key)
	if err != nil {
		return nil, err
	}

	decipheredBytes, err := crypter.Open(nil, []byte(nonce), encryptedData, nil)
	if err != nil {
		return nil, err
	}

	return decipheredBytes, nil
}

func readEncryptedFile(filename string) (*EncryptedPackage, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)

	var pkg EncryptedPackage

	if err := decoder.Decode(&pkg); err != nil {
		return nil, err
	}

	return &pkg, nil
}
