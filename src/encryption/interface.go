package encryption

type (
	Encryption interface {
		PlainText() PlainText
		CipherText() CipherText
	}

	PlainText interface {
		Encrypt([]byte) error
		AsCipherText() CipherText
		String() string
	}

	CipherText interface {
		Decrypt([]byte) error
		AsPlainText() PlainText
		String() string
	}
)
