package main

import(
	pb "github.com/chauhanr/shipcon-user-service/proto/user"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var(
	// secure key string to be used as salt
	key = []byte("mySecureSecretKey")
)

type CustomClaims struct{
	User *pb.User
	jwt.StandardClaims
}

type Authable interface{
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct{
	repo Repository
}

func (svc *TokenService) Decode(tokenString string) (*CustomClaims, error){
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error){
		return key,nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid{
		return claims, nil
	}else{
		return nil, err
	}
}

func (srv *TokenService) Encode(user *pb.User) (string, error){
	expireToken := time.Now().Add(time.Hour*72).Unix()

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer: "go.micro.srv.user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}