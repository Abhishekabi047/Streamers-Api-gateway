package handlers

import (
	"api/pkg/client/interfaces"
	"api/pkg/middlewares"
	"api/pkg/models"
	"api/pkg/utils"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Client interfaces.AuthClient
}

func NewAuthHandler(client interfaces.AuthClient) AuthHandler {
	return AuthHandler{
		Client: client,
	}
}

func (cc *AuthHandler) Signup(c *gin.Context) {
	phone := c.PostForm("phone")
	dob := c.PostForm("dob")
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	cpassword := c.PostForm("cpassword")
	fmt.Println("check 1", username)
	body := models.SignupBody{
		UserName:  username,
		DOB:       dob,
		Phone:     phone,
		Email:     email,
		Password:  password,
		CPassword: cpassword,
	}

	res, err := cc.Client.Signup(context.Background(), body)
	if err != nil {
		errmsg:=utils.ExtractError(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message":errmsg,
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, &res)
}

func (cc *AuthHandler) Otp(c *gin.Context) {
	otp := c.PostForm("otp")
	key := c.PostForm("key")
	body := models.Otpbody{
		Otp: otp,
		Key: key,
	}
	res, err := cc.Client.Otp(context.Background(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (cc *AuthHandler) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	fmt.Println("check 1", password)
	body := models.Login{
		Email:    email,
		Password: password,
	}
	
	res, err := cc.Client.Login(context.Background(), body)
	if err != nil {
		errMsg:=utils.ExtractError(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message":errMsg,
			"err": err.Error(),
		})
		return
	}
	fmt.Println("check 2",res.Id)
	token:= middlewares.CreateJwtCookie(int(res.Id), email, "user", c)

	c.JSON(http.StatusOK, gin.H{"Message": "login succesful","token":token})
}
