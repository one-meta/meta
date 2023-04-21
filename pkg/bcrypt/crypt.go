package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// Encode returns the bcrypt hash of the password at DefaultCost(10)
func Encode(passwd string) (string, error) {
	passByte := []byte(passwd)
	hash, err := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Matches compares a bcrypt hashed password with its possible
// plaintext equivalent. Returns true on success, or false on failure.
func Matches(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	passByte := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, passByte)
	return err == nil
}
