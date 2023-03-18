# EfsCrypt
This repository is supposed to be all crypto wrapper in golang for eFishery usage (currently in enterprise only).<br/>
## Current implementations:<br/>
- AES/GCM -> AES algorithm with symetrical secreet key.

## How it works:
```
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
```

When you invoke any func that name as the algorithm:
```
// EfsCrypt/Aes
func GCM(x string) Encryption { ... }

// yourCode.go
txt := "Hello World"
crypt := EfsCrypt.GCM(txt)
```

The crypt variable is not specific text mode. To begin we should specify the text is either plain text or cipher text.
```
plaintext := crypt.PlainText()
ciphertext := crypt.CipherText()
```

To begin Encrypt
```
key := []byte("mySuperSecreet32BitlongOfStrings")
err := plaintext.Encrypt(key)
...
```

Convert encrypted plain text to cypher text and decrypted back.
```
ciphertext := plaintext.AsCipherText()
ciphertext.Decrypt(key)

ciphertext.String() == txt // true
plaintext.String() == txt // false because it's already encrypted
```