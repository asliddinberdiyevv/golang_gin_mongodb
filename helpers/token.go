package helpers

import (
	"bookhub/database"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email string
	FirstName string
	LastName string
	Uid string
	UserType string
	jwt.StandardClaims
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
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
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

func ValidateToken(accessToken string)(claims *SignedDetails, msg string){
	token, err := jwt.ParseWithClaims(
		accessToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token os invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}

func UpdateAllTokens(signedAccessToken string, signedRefreshToken string, userID string){
	var c, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{"access_token", signedAccessToken})
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", updatedAt})

	upsert := true
	filter := bson.M{"user_id": userID}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(
		c, 
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	)
	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}
}