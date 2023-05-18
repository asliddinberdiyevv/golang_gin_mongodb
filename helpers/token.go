package helpers

import (
	"bookhub/database"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email string
	FirstName string
	LastName string
	Uid string
	UserType string
	jwt.StandartClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SecretKey string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedAccessToken string, signedRefreshToken string, err error) {
	accessClaims := &SignedDetails{
		Email: email,
		FirstName: firstName,
		LastName: lastName,
		Uid: uid,
		UserType: userType,
		StandartClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandartClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(SecretKey))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SecretKey))
	if err != nil {
		log.Panic(err)
		return
	}

	return accessToken, refreshToken, err
}
