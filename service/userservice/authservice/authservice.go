package authservice

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatemehmirarab/gameapp/entity"
	"github.com/golang-jwt/jwt/v5"
)

// location of the files used for signing and verification
const (
	privKeyPath = "test/sample_key"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "test/sample_key.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

// read the key files before starting http handlers
func init() {
	signBytes, err := os.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)
	verifyBytes, err := os.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	ExpirationTime        time.Duration
	RefreshExpirationTime time.Duration
	AccessSubject         string
	RefreshSubject        string
}

type Service struct {
	Config Config
}

func New(config Config) Service {

	return Service{Config: config}
}

type Claims struct {
	jwt.RegisteredClaims
	UserId uint
}

func (c Claims) Valid() error {
	return c.Valid()
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.CreateToken(user.Id, s.Config.AccessSubject, s.Config.ExpirationTime)
}

func (s Service) RefreshToken(user entity.User) (string, error) {
	return s.CreateToken(user.Id, s.Config.RefreshSubject, s.Config.RefreshExpirationTime)
}

func (s Service) CreateToken(userId uint, subject string, expirationTime time.Duration) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// set the expire time
			// see https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			Subject:   subject,
		},
		UserId: userId,
	}

	// Creat token string
	return t.SignedString(signKey)
}

func (s Service) ParseJWT(tokenString string) (*Claims, error) {

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*Claims)
	fmt.Println(claims.UserId, claims.RegisteredClaims.ExpiresAt)

	return claims, nil
	// Output: test
}
