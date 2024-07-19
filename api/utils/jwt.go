package utils

import (
	"UBC/api/library/ctxdata"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GetJwtToken(email, userId, accessSecret string, accessExpire int64) (accessToken string, expire, aRefreshAfter int64, err error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = accessExpire + now
	claims["iat"] = accessExpire
	claims[string(ctxdata.CtxKeyJwtUserEmail)] = email
	claims[string(ctxdata.CtxKeyUserId)] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	accessToken, err = token.SignedString([]byte(accessSecret))
	if err != nil {
		return
	}
	expire = now + accessExpire
	aRefreshAfter = now + accessExpire/2
	return
}
