package helper

import (
	"errors"
	"finalproject4/repository"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

type MyClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claims := MyClaims{
		userID,
		jwt.RegisteredClaims{
			Issuer:    "hacktiv8-final3",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *jwtService) AuthMiddleware(authService Service, userRepository repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}

		userID := int(claim["user_id"].(float64))

		getUser, err := userRepository.GetUserByID(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}

		getRole := getUser.Role

		c.Set("currentUser", getUser)
		c.Set("currentUserRole", getRole)
	}
}

func GetUserID(ctx *gin.Context) (int, bool) {
	var userID int
	id, ok := ctx.Get("userID")
	if !ok {
		return userID, false
	}
	userID = id.(int)

	return userID, true
}
