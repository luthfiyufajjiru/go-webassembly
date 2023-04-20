package libraries

import (
	"encoding/base64"
	"encryption/aes"
	"syscall/js"
)

func EncryptFn(key []byte) {
	js.Global().Set("encrypt", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		_plaintext := []byte(args[0].String())
		plaintext := aes.GCM(string(_plaintext)).PlainText()
		plaintext.Encrypt(key)
		return base64.StdEncoding.EncodeToString([]byte(plaintext.String()))
	}))
}

func DecryptFn(key []byte) {
	js.Global().Set("decrypt", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		_cipher := []byte(args[0].String())
		_cipher, _ = base64.StdEncoding.DecodeString(string(_cipher))
		cipher := aes.GCM(string(_cipher)).CipherText()
		cipher.Decrypt(key)
		return cipher.String()
	}))
}
