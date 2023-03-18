package libraries

import (
	"encoding/base64"
	"encryption/aes"
	"syscall/js"
)

var key = []byte("AES256Key-32Characters1234567890")

func EncryptFn() {
	js.Global().Set("encrypt", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		_plaintext := []byte(args[0].String())
		plaintext := aes.GCM(string(_plaintext)).PlainText()
		plaintext.Encrypt(key)
		return base64.StdEncoding.EncodeToString([]byte(plaintext.String()))
	}))
}
