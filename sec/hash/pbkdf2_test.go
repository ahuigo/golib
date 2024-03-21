package hash

// refer: post/sec/sec-hash-kdf.md
import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"hash"
	"testing"

	"golang.org/x/crypto/pbkdf2"
)

// http://www.ietf.org/rfc/rfc2898.txt
func pbkdf2hash(passphrase string, salt []byte, h func() hash.Hash) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 16)
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 8192, 32, h), salt
}

func benchmark_pbkdf2(b *testing.B, h func() hash.Hash) {
	b.StopTimer()
	b.StartTimer()
	password := "secret"
	hash, salt := pbkdf2hash(password, nil, h) // ignore error for the sake of simplicity

	for i := 0; i < b.N; i++ {
		hash2, _ := pbkdf2hash(password, salt, h)
		if string(hash) != string(hash2) {
			b.Errorf("hash not matched")
		}
	}
}

func Benchmark_pbkdf2sha256(b *testing.B) {
	benchmark_pbkdf2(b, sha256.New)
}

func Benchmark_pbkdf2md5(b *testing.B) {
	benchmark_pbkdf2(b, md5.New)
}
