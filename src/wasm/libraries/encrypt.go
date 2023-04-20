package libraries

import (
	"encoding/base64"
	"encryption/aes"
	"os"
	"syscall/js"

	"github.com/joho/godotenv"
)

var key []byte

// init and read the env using godotenv
func init() {
	_ = godotenv.Load()
	key = []byte(os.Getenv("KEY"))
}

func EncryptFn() {
	js.Global().Set("encrypt", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		_plaintext := []byte(args[0].String())
		plaintext := aes.GCM(string(_plaintext)).PlainText()
		plaintext.Encrypt(key)
		return base64.StdEncoding.EncodeToString([]byte(plaintext.String()))
	}))
}

func DecryptFn() {
	js.Global().Set("decrypt", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		_cipher := []byte(args[0].String())
		_cipher, _ = base64.StdEncoding.DecodeString(string(_cipher))
		cipher := aes.GCM(string(_cipher)).CipherText()
		cipher.Decrypt(key)
		return cipher.String()
	}))
}
