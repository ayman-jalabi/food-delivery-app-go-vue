package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"main/models"
	"strings"
	"time"
)

type AccessTokenClaims struct {
	//ID uint `json:"id"`
	//below is an embedded structure, if you hover over it you'll see that'll be embedded
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	//ID uint `json:"id"`
	jwt.RegisteredClaims
}

type TokenService struct {
	accessConfig  *models.AccessConfig
	refreshConfig *models.RefreshConfig
}

func NewTokenService(accessConfig *models.AccessConfig, refreshConfig *models.RefreshConfig) *TokenService {
	return &TokenService{
		accessConfig:  accessConfig,
		refreshConfig: refreshConfig,
	}
}

//func NewRefreshService(claims jwt.RegisteredClaims) *RefreshTokenClaims {
//	return &RefreshTokenClaims{}
//}

func (s *TokenService) GenerateAccessToken(userId string) (token string, err error) {
	claims := &AccessTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID: userId,
			//The field below determines how long our token will be valid for.
			//The code below will make the token valid for 30 minutes since its creation:
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(s.accessConfig.AccessTokenLifetime))),
		},
	}
	//this will have the token structure:
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenStruct.SignedString([]byte(s.accessConfig.AccessTokenSecret))
}

func (s *TokenService) ValidateAccessToken(tokenString string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString,
		&AccessTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.accessConfig.AccessTokenSecret), nil
		})
	// Handle parsing errors
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token expired")
		}
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("malformed token")
		}
		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, errors.New("token not valid yet")
		}
		return nil, errors.New("invalid token")
	}

	// Extract the claims and ensure they are valid
	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse claims or invalid token")
	}

	// Additional checks on claims can be done here, e.g., checking user ID
	if claims.ID == "" {
		return nil, errors.New("invalid token: missing user ID")
	}

	return claims, nil
}

// GenerateRefreshToken creates a refresh token with a longer 7 day expiration time.
func (s *TokenService) GenerateRefreshToken(userId string) (string, error) {
	claims := &RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID: userId,
			//The field below determines how long our token will be valid for.
			//The code below will make the refresh-token valid for 7 days since its creation:
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(
				time.Hour * time.Duration(s.refreshConfig.RefreshTokenLifetime))),
		},
	}

	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenStruct.SignedString([]byte(s.refreshConfig.RefreshTokenSecret))
}

func (s *TokenService) ValidateRefreshToken(tokenString string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString,
		&RefreshTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.refreshConfig.RefreshTokenSecret), nil
		})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse claims token")
	}

	return claims, nil
}

func GetTokenFromBearerString(bearerString string) string {
	if bearerString == "" {
		return ""
	}

	//this will split the string Bearer from the token after it within the authorization part of the http response header
	parts := strings.Split(bearerString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	token := strings.TrimSpace(parts[1])
	if len(token) == 0 {
		return ""
	}
	return token
}
