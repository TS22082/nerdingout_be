package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TS22082/nerdingout_be/types"
	"github.com/TS22082/nerdingout_be/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
	"os"
	"time"
)

func GhLogin(ctx *fiber.Ctx) error {

	code := ctx.Query("code")
	if code == "" || code == "null" {
		return fiber.ErrBadRequest
	}

	url := "https://github.com/login/oauth/access_token"
	ghAuthPayload := map[string]string{
		"client_id":     os.Getenv("GH_ID"),
		"client_secret": os.Getenv("GH_SECRET"),
		"code":          code,
	}

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	ghAuthParams := types.HTTPRequestParams{
		URL:     url,
		Method:  "POST",
		Headers: headers,
		Body:    ghAuthPayload,
	}

	var bodyBytes []byte
	bodyBytes, marshalErr := json.Marshal(ghAuthParams.Body)

	if marshalErr != nil {
		return fiber.ErrInternalServerError
	}

	var req, err = http.NewRequest(ghAuthParams.Method, ghAuthParams.URL, bytes.NewReader(bodyBytes))

	if err != nil {
		return fiber.ErrInternalServerError
	}

	for key, value := range ghAuthParams.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	defer resp.Body.Close()

	bodyBytes, err = io.ReadAll(resp.Body)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	var ghAuthResult map[string]interface{}

	err = json.Unmarshal(bodyBytes, &ghAuthResult)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	if ghAuthResult["error"] != nil {
		return fiber.ErrBadRequest
	}

	var accessToken = ghAuthResult["access_token"].(string)

	if at, ok := ghAuthResult["access_token"].(string); ok {
		accessToken = at
	} else {
		return fiber.ErrInternalServerError
	}

	userEmailParams := map[string]interface{}{
		"URL":    "https://api.github.com/user/emails",
		"Method": "GET",
		"Headers": map[string]string{
			"Accept":        "application/json",
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("token %v", accessToken),
		},
	}

	req, err = http.NewRequest(userEmailParams["Method"].(string), userEmailParams["URL"].(string), nil)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	for key, value := range userEmailParams["Headers"].(map[string]string) {
		req.Header.Set(key, value)
	}

	client = &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	var emailData []map[string]interface{}
	err = json.Unmarshal(bodyBytes, &emailData)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	primaryEmail := ""

	for _, email := range emailData {
		if email["primary"].(bool) {
			primaryEmail = email["email"].(string)
			break
		}
	}

	if err := utils.ValidateEmail(primaryEmail); err != nil {
		return fiber.ErrInternalServerError
	}

	mongoDB := ctx.Locals("mongoDB").(*mongo.Database)
	userCollection := mongoDB.Collection("Users")
	userFound := userCollection.FindOne(context.Background(), bson.D{{Key: "email", Value: primaryEmail}})

	var user types.User

	err = userFound.Decode(&user)

	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		newUser := types.User{
			Email:     primaryEmail,
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
		}
		insertResult, err := userCollection.InsertOne(context.Background(), newUser)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		newUser.ID = insertResult.InsertedID.(primitive.ObjectID)
		user = newUser
	}

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return fiber.ErrInternalServerError
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID.Hex(),
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return ctx.JSON(fiber.Map{
		"token": tokenString,
		"user":  user,
	})
}
