package hash

// refer: post/sec/sec-hash.md
import (
	"crypto/sha1"
	"encoding/hex"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// go test -v -bench=. bcrypt_test.go

func checkSha1(s, hash string) bool {
	return hash == generateSha1(s)
}
func generateSha1(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func Benchmark_sha1(b *testing.B) {
	password := "secret"
	hash := generateSha1(password)
	for i := 0; i < b.N; i++ {
		checkSha1(password, hash)
	}
}

func Benchmark_bcrypt(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	password := "secret"
	hash, _ := bcryptPassword(password) // ignore error for the sake of simplicity

	for i := 0; i < b.N; i++ {
		bcryptPasswordCheck(password, hash)
	}
}

func bcryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func bcryptPasswordCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
