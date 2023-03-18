package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encryption"
	"io"

	"github.com/sirupsen/logrus"
)

type (
	gcmString     string
	gcmPlainText  string
	gcmCipherText string
)

func GCM(inp string) encryption.Encryption {
	c := gcmString(inp)
	return &c
}

func (k *gcmString) PlainText() encryption.PlainText {
	c := string(*k)
	p := gcmPlainText(c)
	return &p
}
func (k *gcmString) CipherText() encryption.CipherText {
	c := string(*k)
	p := gcmCipherText(c)
	return &p
}

func (e *gcmPlainText) Encrypt(key []byte) (err error) {
	if len(key) != 32 {
		err = encryption.ErrKeyLength
		logrus.New().Errorf("%s. Expect 32 but got %d", err.Error(), len(key))
		return
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		logrus.New().Error(err)
		return
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	*e = gcmPlainText(gcm.Seal(nonce, nonce, []byte(string(*e)), nil))
	return
}
func (e *gcmPlainText) String() string {
	return string(*e)
}
func (e *gcmPlainText) AsCipherText() encryption.CipherText {
	s := e.String()
	v := gcmCipherText(s)
	return &v
}

func (e *gcmCipherText) Decrypt(key []byte) (err error) {
	if len(key) != 32 {
		err = encryption.ErrKeyLength
		logrus.New().Errorf("%s. Expect 32 but got %d", err.Error(), len(key))
		return
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		logrus.New().Error(err)
		return
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		logrus.New().Error(err)
		return
	}

	ciphertext := []byte(string(*e))
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		logrus.New().Error(err)
		return
	}

	nonce := ciphertext[:nonceSize]
	ciphertext = ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		logrus.New().Error(err)
		return
	}

	*e = gcmCipherText(plaintext)

	return err
}
func (e *gcmCipherText) String() string {
	return string(*e)
}

func (e *gcmCipherText) AsPlainText() encryption.PlainText {
	s := e.String()
	v := gcmPlainText(s)
	return &v
}
