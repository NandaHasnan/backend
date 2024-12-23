// package lib

// import (
// 	"github.com/pilinux/argon2"
// )

// var SECRET_key string = "FAZZTRACK"

// func VerifyHash(hash string, password string) bool {
// 	result, _ := argon2.ComparePasswordAndHash(password, SECRET_key, hash)
// 	return result
// }

// func GenerateHash(password string) string {
// 	result, _ := argon2.CreateHash(password, SECRET_key, argon2.DefaultParams)
// 	return result
// }

package lib

import (
	"github.com/pilinux/argon2"
)

var SECRET_key = "FAZZTRACK"

func VerifyHash(hash string, password string) bool {
	result, _ := argon2.ComparePasswordAndHash(password, SECRET_key, hash)
	return result
}

func GenerateHash(password string) string {
	hash, _ := argon2.CreateHash(password, SECRET_key, argon2.DefaultParams)
	return hash
}
