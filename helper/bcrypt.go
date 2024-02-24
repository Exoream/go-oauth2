package helper

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword menghasilkan hash dari password menggunakan bcrypt.
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPasswordHash memeriksa apakah password cocok dengan hash yang diberikan.
func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateRandomState(length int) (string, error) {
	// Jumlah byte yang dibutuhkan untuk string acak
	byteLength := (length * 3) / 4

	// Membuat slice untuk menampung byte acak
	randomBytes := make([]byte, byteLength)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Mengonversi byte menjadi string base64 yang aman untuk URL
	randomString := base64.URLEncoding.EncodeToString(randomBytes)

	// Mengambil sebagian string untuk mencapai panjang yang diinginkan
	return randomString[:length], nil
}
