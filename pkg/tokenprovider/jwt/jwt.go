package jwt

import (
	"github.com/cesc1802/auth-service/pkg/tokenprovider"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtProvider struct {
	secret string
	expiry uint
}

func NewTokenJWTProvider(secret string, expiry uint) *jwtProvider {
	return &jwtProvider{secret: secret, expiry: expiry}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload) (*tokenprovider.Token, error) {
	// generate the JWT
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(j.expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  j.expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (*tokenprovider.TokenPayload, error) {

	// parse the public key
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	// return the token
	return &claims.Payload, nil
}

func (j *jwtProvider) String() string {
	return "JWT implement Provider"
}
