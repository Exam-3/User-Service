package auth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
	pb "user_service/genproto/authentication"
)

const (
	newsigningkey = "visca barsa visca kataluniya"
)

func GeneratedRefreshJWTToken(req *pb.LoginResponse) error {
	token := *jwt.New(jwt.SigningMethodHS256)

	//payload
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = req.Refresh.Userid
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	newToken, err := token.SignedString([]byte(newsigningkey))
	if err != nil {
		log.Println(err)
		return err
	}
	
	req.Refresh.Refreshtoken = newToken
	return nil
}

func ValidateRefreshToken(tokenStr string) (bool, error) {
	_, err := ExtractRefreshClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractRefreshClaim(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(newsigningkey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return &claims, nil
}

func GetUserIdFromRefreshToken(refreshTokenString string) (string, error) {
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) { return []byte(newsigningkey), nil })
	if err != nil || !refreshToken.Valid {
		return "", err
	}
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}
	userID := claims["user_id"].(string)

	return userID, nil
}
