// kdf.go
package main

import (
    "crypto/sha3"
    "encoding/binary"
)

// Derive chaos system initial conditions from master key + nonce
func deriveChaosParams(masterKey []byte, nonce []byte) (seed int64) {
    hash := sha3.New512()
    hash.Write(masterKey)
    hash.Write(nonce)
    digest := hash.Sum(nil)
    // Use first 8 bytes as seed
    seed = int64(binary.LittleEndian.Uint64(digest[:8]))
    return
}
