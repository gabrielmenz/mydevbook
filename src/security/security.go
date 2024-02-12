package security

import "golang.org/x/crypto/bcrypt"

// Receives a string and adds a hash to it
func Hash(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
}

// Compares a pw to a hsh and returns if they're equivalent
func VerifyPw(pwWithHash, pwString string) error {
	return bcrypt.CompareHashAndPassword([]byte(pwWithHash), []byte(pwString))
}
