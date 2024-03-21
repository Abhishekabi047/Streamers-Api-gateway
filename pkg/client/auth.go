package auth

import (
	"api/pkg/client/interfaces"
	"api/pkg/config"
	"api/pkg/models"
	"api/pkg/pb/auth"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	Server auth.AuthServiceClient
}

func InitClient(c *config.Config) (auth.AuthServiceClient, error) {
	cc, err := grpc.Dial(c.AuthService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return auth.NewAuthServiceClient(cc), nil
}

func NewAuthServiceClient(server auth.AuthServiceClient) interfaces.AuthClient {
	return &AuthClient{
		Server: server,
	}
}

func (c *AuthClient) Signup(ctx context.Context, request models.SignupBody) (*auth.SignUpResponse,error) {
	res,err:=c.Server.Signup(ctx,&auth.SignUpRequest{
		Email: request.Email,
		Dob: request.DOB,
		Username: request.UserName,
		Password: request.Password,
		Cpassword: request.CPassword,
		Phone: request.Phone,
	})
	if err != nil{
		return nil,err
	}
	return res,nil
}

func(c *AuthClient) Otp(ctx context.Context,req models.Otpbody) (*auth.OtpResponse,error) {
	res,err:=c.Server.Otp(ctx,&auth.OtpRequest{
		Otp: req.Otp,
		Key: req.Key,
	})
	if err != nil{
		return nil,err
	}
	return res,nil
}

func(c *AuthClient) Login(ctx context.Context,req models.Login) (*auth.LoginResponse,error) {
	res,err:=c.Server.Login(ctx,&auth.LoginRequest{
		Email: req.Email,
		Password: req.Password,
	})
	if err != nil{
		return nil,err
	}
	return res,nil
}