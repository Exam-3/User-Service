package auth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
	pb "user_service/genproto/authentication"
)

const (
	signingkey = "visca barsa"
)

func GeneratedAccessJWTToken(req *pb.LoginResponse) error {
	token := *jwt.New(jwt.SigningMethodHS256)

	//payload
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = req.Access.Id
	claims["username"] = req.Access.Username
	claims["email"] = req.Access.Email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()

	newToken, err := token.SignedString([]byte(signingkey))

	if err != nil {
		log.Println(err)
		return err
	}
	req.Access.Accesstoken = newToken

	return nil
}

func ValidateAccesToken(tokenStr string) (bool, error) {
	_, err := ExtractAccesClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractAccesClaim(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingkey), nil
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
