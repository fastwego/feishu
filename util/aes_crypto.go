// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

func AESDecrypt(base64CipherText string, key []byte) (unpadDecrypted []byte, err error) {
	sum := sha256.Sum256(key)

	key = []byte(``)
	for _, b := range sum {
		key = append(key, b)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(base64CipherText)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	if len(ciphertext) < aes.BlockSize {
		err = errors.New("len(crypt) < aes.BlockSize")
		return
	}
	cbc := cipher.NewCBCDecrypter(block, ciphertext[:aes.BlockSize])
	ciphertext = ciphertext[aes.BlockSize:]
	decrypted := make([]byte, len(ciphertext))
	cbc.CryptBlocks(decrypted, ciphertext)

	unpadDecrypted = PKCS5Trimming(decrypted)
	return
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
