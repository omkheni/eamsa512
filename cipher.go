package cipher



// BlockSize is the fixed block size in bytes.

const BlockSize = 64 // 512 bits



// Cipher holds expanded key material and state.

type Cipher struct {

// internal fields (hidden)

}



// NewCipher constructs a cipher instance from 11 × 128-bit subkeys.

func NewCipher(subkeys [][]byte) (*Cipher, error)



// Encrypt encrypts plaintext and returns ciphertext (may be padded or chunked).

func (c *Cipher) Encrypt(plaintext []byte) ([]byte, error)



// Decrypt decrypts ciphertext and returns the original plaintext.

func (c *Cipher) Decrypt(ciphertext []byte) ([]byte, error)
