package interfaces

import (
	"api/pkg/models"
	"api/pkg/pb/auth"
	"context"
)

type AuthClient interface {
	Login(context.Context, models.Login) (*auth.LoginResponse, error)
	Otp(context.Context, models.Otpbody) (*auth.OtpResponse, error)
	Signup(context.Context, models.SignupBody) (*auth.SignUpResponse, error)
	SearchUser( context.Context, models.SearchRequest) (*auth.SearchUserResponse,error)
}
