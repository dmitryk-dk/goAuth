package token

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

var mySigningKey = "./mykey.priv"

const (
	TOKEN_JWT_ID  = "123-id"
	TOKEN_ISSUER  = "http://supersite.com/"
	TOKEN_SUBJECT = "username@supersite.com"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func GetTokenHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		bytes, _ := ioutil.ReadFile(mySigningKey)
		// create claims
		claims := jws.Claims{}
		// set params for token
		claims.SetSubject(TOKEN_ISSUER)
		claims.SetIssuer(TOKEN_SUBJECT)
		claims.SetExpiration(time.Now().Add(time.Hour * 24))
		claims.SetJWTID(TOKEN_JWT_ID)
		// signed token of secure string
		rsaPrivate, _ := crypto.ParseRSAPrivateKeyFromPEM(bytes)
		token := jws.NewJWT(claims, crypto.SigningMethodRS256)
		serializedToken, _ := token.Serialize(rsaPrivate)

		// response
		w.Write(serializedToken)
	})
}

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	var token TokenResponse
	token.Token = string(tokenGen()[:])
	data, _ := json.Marshal(token)
	w.Write(data)
}

func tokenGen() []byte {
	bytes, _ := ioutil.ReadFile(mySigningKey)
	// create claims
	claims := jws.Claims{}
	// set params for token
	claims.SetSubject(TOKEN_ISSUER)
	claims.SetIssuer(TOKEN_SUBJECT)
	claims.SetExpiration(time.Now().Add(time.Hour * 24))
	claims.SetJWTID(TOKEN_JWT_ID)
	// signed token of secure string
	rsaPrivate, _ := crypto.ParseRSAPrivateKeyFromPEM(bytes)
	token := jws.NewJWT(claims, crypto.SigningMethodRS256)
	serializedToken, _ := token.Serialize(rsaPrivate)
	return serializedToken
}
