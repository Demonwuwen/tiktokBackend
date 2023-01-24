package util

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

// GenPasswordHash 密码加密: pwdHash  同PHP函数 password_hash()
func GenPasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// PasswordVerify verify password is correct: true or false
func PasswordVerify(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))

	return err == nil
}

const letters = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func RandNameString(n int) string {
	var builder strings.Builder
	letterLen := len(letters)
	rand.Seed(time.Now().Unix())
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i := 0; i < n; i++ {
		idx := rand.Intn(letterLen)
		builder.WriteByte(letters[idx])
	}
	return builder.String()
}
