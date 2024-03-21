package middlewares

import (
	"api/pkg/config"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserRetriveCookie(c *gin.Context) {

	valid := ValidateCookie(c)
	if valid == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		c.Abort()
	} else {
		userId, Email, role, err := RetriveJwtToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "cookie retriving failed"})
			c.Abort()
		} else {
			c.Set("userId", userId)
			c.Set("email", Email)
		}
		if role != "user" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "role mismatching"})
			c.Abort()
		}
	}
	c.Next()
}
func AdminRetriveCookie(c *gin.Context) {

	valid := ValidateCookie(c)
	if valid == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
		c.Abort()
	} else {
		userId, Email, role, err := RetriveJwtToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "cookie retriving failed"})
			c.Abort()
		} else {
			c.Set("userID", userId)
			c.Set("Email", Email)
		}
		if role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "role mismatching"})
			c.Abort()
		} else {
			c.Next()
		}
	}

}

func CreateJwtCookie(id int, userEmail string, role string, c *gin.Context) (string){
	config, _ := config.LoadConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": id,
		"email":  userEmail,
		"role":   role,
	})
	tokenString, err := token.SignedString([]byte(config.Jwt))

	if err == nil {
		fmt.Println("token created")
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorise", tokenString, 3600, "", "", false, true)
	return tokenString
}

func ValidateCookie(c *gin.Context) bool {
	cookie, _ := c.Cookie("Authorise")
	if cookie == "" {
		fmt.Println("cookie not found")
		return false
	} else {
		return true
	}

}

func RetriveJwtToken(c *gin.Context) (int, int, string, error) {
	config, _ := config.LoadConfig()
	cookie, _ := c.Cookie("Authorise")
	if cookie == "" {
		return 0, 0, "", errors.New("cookie not found")
	} else {
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Jwt), nil
		})

		if err != nil {
			return 0, 0, "", err
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var userId, email int
			var role string
			if id, exists := claims["userId"]; exists {
				if idFloat, ok := id.(float64); ok {
					userId = int(idFloat)
				}
			}

			if email, exists := claims["email"]; exists {
				if emailstr, ok := email.(string); ok {
					email = emailstr
				}
			}

			if r, exists := claims["role"]; exists {
				if roleString, ok := r.(string); ok {
					role = roleString
				}
			}

			return userId, email, role, nil
		} else {
			return 0, 0, "", fmt.Errorf("invalid token")
		}

	}
}

func DeleteCookie(c *gin.Context) error {
	c.SetCookie("Authorise", "", 0, "", "", true, true)
	fmt.Println("cookie deleted")
	return nil
}
