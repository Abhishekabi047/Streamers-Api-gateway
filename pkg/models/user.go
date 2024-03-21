package models

type SignupBody struct{
	Email string `json:"email"`
	Password string `json:"password"`
	CPassword string `json:"cpassword"`
	DOB string `json:"dob"`
	UserName string `json:"username"`
	Phone string `json:"phone"`
}

type Otpbody struct{
	Otp string `json:"otp"`
	Key string `json:"key"`
}

type Login struct{
	Email string `json:"email"`
	Password string `json:"password"`
}