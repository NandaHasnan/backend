package lib

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

func Getmd5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

var SECRET []byte = []byte(Getmd5("NandaHasnan"))

func GeneratePass(payload any) string {

	sig, _ := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: SECRET}, (&jose.SignerOptions{}).WithType("JWT"))
	baseInfo := jwt.Claims{
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}
	token, _ := jwt.Signed(sig).Claims(baseInfo).Claims(payload).Serialize()
	return token
}
