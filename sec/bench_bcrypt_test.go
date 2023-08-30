package gotest
import "testing"
import "fmt"
import "golang.org/x/crypto/bcrypt"
// go test -v -bench=. benchmark_bcrypt_test.go

import (
    "crypto/sha1"
    "encoding/hex"
)

func checkSha1(s, hash string) bool {
    return hash == generateSha1(s)
}
func generateSha1(password string) string{
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
    b.StopTimer() //调用该函数停止压力测试的时间计数
    password := "secret"
    hash, _ := HashPassword(password) // ignore error for the sake of simplicity
    b.StartTimer() //重新开始时间

    fmt.Println("Password:", password)
    fmt.Println("Hash:    ", hash)

    for i := 0; i < b.N; i++ {
        CheckPasswordHash(password, hash)
    }
}





func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

