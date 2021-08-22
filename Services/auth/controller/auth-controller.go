package controller

import (
	"fmt"
	"goprueba/Services/auth/dto"
	"goprueba/Services/auth/repository"
	"goprueba/Services/auth/service"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

func Login(c *gin.Context) {
	var u dto.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	fmt.Printf("u: %v\n", u)

	foundUser, err := repository.GetUserByEmail(u.Email)
	fmt.Printf("err3333: %v\n", err)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	fmt.Printf("foundUser: %v\n", foundUser)

	pwdOk := service.CheckPasswordHash(u.Password, foundUser.Password)
	//compare the user from the request, with the one we defined:
	if foundUser.Email != u.Email || !pwdOk {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	ts, err := service.CreateToken(foundUser.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	fmt.Printf("tssssssss: %v\n", ts)
	saveErr := service.CreateAuth(foundUser.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func Logout(c *gin.Context) {
	au, err := service.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	deleted, delErr := service.DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

func Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Error occurred")
			return
		}
		//Delete the previous Refresh Token
		deleted, delErr := service.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := service.CreateToken(userId)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}
		//save the tokens metadata to redis
		saveErr := service.CreateAuth(userId, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}

func CreateUser(c *gin.Context) {
	var u *dto.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	hashPwd, err := service.HashPassword(u.Password)
	if err != nil {
		c.JSON(http.StatusNoContent, "fail pwd")
		return
	}

	u.Password = hashPwd

	fmt.Printf("hashPwd: %v\n", hashPwd)
	fmt.Printf("u.Email: %v\n", u.Email)

	//userId, err := c.Get("userId")

	userCreated, err := repository.CreateUser(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Printf("userCreated: %v\n", userCreated)
	c.JSON(http.StatusCreated, userCreated)
}
