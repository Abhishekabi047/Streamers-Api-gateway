package handlers

import (
	"api/pkg/client/interfaces"
	"api/pkg/middlewares"
	"api/pkg/models"
	"api/pkg/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"

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

	c.JSON(http.StatusOK, gin.H{"Message": "login succesful","token":token,"username":res.Username})
}

func (a *AuthHandler) SearchUser(c *gin.Context) {
	user:=c.PostForm("username")
	limit:=c.Query("limit")
	limitin,_:=strconv.Atoi(limit)
	offset:=c.Query("offset")
	offsetin,_:=strconv.Atoi(offset)
	userde:=&models.SearchRequest{
		Username: user,
		Offset: offsetin,
		Limit: limitin,
	}
	fmt.Println("user",user)
	res,err:=a.Client.SearchUser(context.Background(),*userde) 
	if err != nil {
		errMsg:=utils.ExtractError(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message":errMsg,
			"err": err.Error(),
		})
		return
	}
	fmt.Println("res",res.Userdetails)
	c.JSON(http.StatusOK,res)

}
