package common

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"
	"github.com/davecgh/go-spew/spew"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	mu "seyes-core/internal/model/user"
)

// UserInfo define user map with claims jwt
type UserInfo struct {
	ID        int64  `json:"id"`
	Active    bool   `json:"active"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Tel       string `json:"tel"`
	Password  string `json:"password"` //FIXME
	Email     string `json:"email"`
}

// User define for map request login
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Authenticator domain interface
type Authenticator struct {
	User        interface{}
	privateKey  *rsa.PrivateKey
	publicKey   *rsa.PublicKey
}


// AuthUserParams define params for user
type AuthUserParams struct {
	ID        int64  `json:"id"`
	Active    bool   `json:"active"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Tel       string `json:"tel"`
	Password  string `json:"password"` //FIXME
	Email     string `json:"email"`
}

// NewAuthenticator return new Authenticator instance
func NewAuthenticator() (*Authenticator, error) {
	priv := os.Getenv("AUTH_PRIVATE_KEY")
	pub := os.Getenv("AUTH_PUBLIC_KEY")

	privbytes, _ := base64.StdEncoding.DecodeString(priv)
	pubbytes, _ := base64.StdEncoding.DecodeString(pub)

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privbytes))

	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubbytes))

	if err != nil {
		return nil, err
	}


	return &Authenticator{
		privateKey:  privateKey,
		publicKey:   publicKey,
	
	}, nil
}


// GetUserByEmail get a user by email
func GetUserByEmail(db *gorm.DB, email string) (*mu.User, error) {
	var u mu.User

	if err := db.
		Where("users.email = ?", email).
		First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}


// GetUserFromCtx get a user by user context
func GetUserFromCtx(db *gorm.DB, uID int64) (*AuthUserParams, error) {
	var user AuthUserParams

	if err := db.Model(&mu.User{}).
		Select("users.id AS id",
			"users.active AS active",
			"users.first_name AS first_name",
			"users.last_name AS last_name",
			"users.tel AS tel",
			"users.email AS email").
		Where("users.id = ?", uID).
		Scan(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}


// SignJWTToken generate signed JWT token for given user
func (a *Authenticator) SignJWTToken( ps map[string]interface{}, limit *time.Duration) (string, jwt.MapClaims, error) {
	t := time.Now()
	c := make(jwt.MapClaims)
	c["sub"] = ps["id"].(uint)
	c["iat"] = time.Now().Unix()

	if limit != nil {
		c["exp"] = t.Add(*limit).Unix()
	} else {
		c["exp"] = t.Add(time.Hour * 24 * 7).Unix()
	}

	if ps != nil {
		for k, v := range ps {
			c[k] = v
		}
	}

	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims = c
	tk, err := token.SignedString(a.privateKey)

	if err != nil {
		return "",  nil, err
	}
	
	return tk, c, nil
}

// VerifyJWTToken verify jwt token string
func (a *Authenticator) VerifyJWTToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid Token")
}

// hashPassword hashes a password with bcrypt
func hashPassword(password string) []byte {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hashedPassword
}

const passwordCost = 14

// HashPassword hashes plain text password
func (a *Authenticator) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	return string(bytes), err
}

// CheckPassword compare hashed and plain text password
func (a *Authenticator) CheckPassword(hashed string, password string) bool {


spew.Dump(hashed)

spew.Dump(password)

res := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
	
spew.Dump(res)


return res
}


// ParseClaimsContext parse claims to struct data
func (a *Authenticator) ParseClaimsContext(db *gorm.DB, claims map[string]interface{}) (*UserInfo, error) {
	var us mu.User
	userID := claims["id"].(float64)

	if err := db.Where("ID=?", userID).First(&us).Error; err != nil {
		return nil, err
	}

	ctx := UserInfo{
		ID:        int64(us.ID),
		Active:      us.Active,
		FirstName:    us.FirstName,
		LastName:  us.LastName,
		Tel:     us.Tel,
		Password:      us.Password,
		Email: us.Email,
	}

	return &ctx, nil
}


