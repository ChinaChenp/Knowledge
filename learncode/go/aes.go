package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
)

func AesEncrypt(text, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("invalid decrypt key")
	}

	text = pKCS5Padding(text, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)

	re := make([]byte, len(text))
	blockMode.CryptBlocks(re, text)

	return re, nil
}

func AesDecrypt(text, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("invalid decrypt key")
	}

	blockSize := block.BlockSize()

	if len(text) < blockSize {
		return nil, errors.New("ciphertext too short")
	}

	if len(text)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	blockModel := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(text))
	blockModel.CryptBlocks(plaintext, text)
	plaintext = pKCS5UnPadding(plaintext)

	return plaintext, nil
}

func pKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func pKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func main() {
	key := os.Args[1]
	decode, _ := base64.StdEncoding.DecodeString(key)
	re, err := AesDecrypt([]byte(decode), []byte("didi_i_checkcard"), []byte("didi_i_checkcard"))
	if err != nil {
		fmt.Println("decrypt error", err)
		return
	}
	fmt.Println("decrypt ok: ", string(re))

	re, err = AesEncrypt(re, []byte("didi_i_checkcard"), []byte("didi_i_checkcard"))
	if err != nil {
		fmt.Println("encrypt error", err)
		return
	}
	s := base64.StdEncoding.EncodeToString(re)
	fmt.Println("encrypt ok: ", s)
}
