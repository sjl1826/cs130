package models

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

func generateTokenPair(u *User) (map[string]string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["id"] = u.ID
	claims["time"] = time.Now()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["email"] = u.Email
	rtClaims["time"] = time.Now()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("secret")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

// Token token model
type Token struct {
	UserID           int           `json:"id"`
	AccessToken      string        `json:"access_token"`
	AccessCreateAt   time.Time     `json:"AccessCreateAt"`
	AccessExpiresIn  time.Duration `json:"AccessExpiresIn"`
	RefreshToken     string        `json:"refresh_token"`
	RefreshCreateAt  time.Time     `json:"RefreshCreateAt"`
	RefreshExpiresIn time.Duration `json:"RefreshExpiresIn"`
}

// New create to token model instance
func (t *Token) New(db *gorm.DB, u *User) error {
	db.Exec("DELETE FROM tokens WHERE user_id=" + strconv.Itoa(u.ID))
	tokens, err := generateTokenPair(u)
	t.AccessToken = tokens["access_token"]
	t.RefreshToken = tokens["refresh_token"]

	now := time.Now()
	t.AccessCreateAt = now
	t.RefreshCreateAt = now

	aExp, err1 := time.ParseDuration("24h")
	rExp, err2 := time.ParseDuration("48h")
	if err1 != nil {
		return err1
	} else if err2 != nil {
		return err2
	}

	t.AccessExpiresIn = aExp
	t.RefreshExpiresIn = rExp
	retVal := db.Create(t).Table("tokens").Scan(&t)

	if retVal.Error != nil {
		return retVal.Error
	}
	return err
}

// GetToken retrieves the token paired to the user
func (t *Token) GetToken(db *gorm.DB) error {
	retVal := db.Raw("SELECT * FROM tokens WHERE user_id=" + strconv.Itoa(t.UserID)).Scan(&t)
	return retVal.Error
}

// GetTokenByAccess retrieves the token paired to the user using the access_token
func (t *Token) GetTokenByAccess(db *gorm.DB) error {
	retVal := db.Raw("SELECT * FROM tokens WHERE access_token='" + t.AccessToken + "'").Scan(&t)
	return retVal.Error
}

// GetTokenByRefresh retrieves the token paired to the user using the refresh_token
func (t *Token) GetTokenByRefresh(db *gorm.DB) error {
	retVal := db.Raw("SELECT * FROM tokens WHERE refresh_token='" + t.RefreshToken + "'").Scan(&t)
	return retVal.Error
}

// RevokeToken removes the token from the database
func (t *Token) RevokeToken(db *gorm.DB) error {
	retVal := db.Exec("DELETE FROM tokens WHERE user_id=" + strconv.Itoa(t.UserID))
	return retVal.Error
}
