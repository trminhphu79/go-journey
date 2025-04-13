package auth

import (
	"app/api/auth/dto"
	"app/api/auth/model"
	"app/arch/network"
	"app/arch/postgres"
	"app/config"
	"app/utils"

	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	logger "github.com/sirupsen/logrus"
)

type AuthService interface {
	RegisterUser(dto.RegistrationDTO) (user *model.User, err error)
	Login(dto.LoginDTO) (res dto.LoginRessponseDTO, err error)
	Authenticate(accessToken string) (user *model.User, err error)
	GeneratePairToken(input model.User) (res dto.LoginRessponseDTO, err error)
	ValidateAccessToken(tokenString string) (claims jwt.MapClaims, err error)
}

type authService struct {
	network.BaseService
	db  postgres.Database
	env *config.Env
}

func CreateAuthService(db postgres.Database, env *config.Env) AuthService {
	if env.RSAPublicKey != "" {
		logger.Info("GET RSAPublicKey success")
	}
	return &authService{
		BaseService: network.NewBaseService(),
		db:          db.GetInstance(),
		env:         env,
	}
}

func (c *authService) RegisterUser(input dto.RegistrationDTO) (user *model.User, err error) {
	value, _ := utils.HashPassword(input.Password)
	logger.Info("Hash password success: ", value)
	model := &model.User{
		Username: input.Username,
		Password: value,
		FullName: input.FullName,
	}
	result := c.db.GetInstance().Create(&model)
	if result.Error != nil {
		logger.WithError(result.Error).Error("Failed to create user")
		return nil, result.Error
	}

	model.Password = ""
	logger.WithField("rowsAffected", result.RowsAffected).Info("User created successfully")
	return model, nil
}

func (c *authService) Login(input dto.LoginDTO) (res dto.LoginRessponseDTO, err error) {
	var user model.User

	result := c.db.GetInstance().Where("username = ?", input.Username).First(&user)
	if result.Error != nil {
		logger.WithError(result.Error).Error("User not found")
		return dto.LoginRessponseDTO{}, result.Error
	}

	valid := utils.ComparePassword(user.Password, input.Password)
	if !valid {
		logger.Error("Invalid credentials")
		return dto.LoginRessponseDTO{}, network.NewUnauthorizedErr("Invalid credentials!", fmt.Errorf("Invalid credentials!"))
	}

	tokens, err := c.GeneratePairToken(user)
	if err != nil {
		logger.Error("Generate tokens failed!")
		return dto.LoginRessponseDTO{}, err
	}

	return tokens, nil
}

func (c *authService) Authenticate(accessToken string) (user *model.User, err error) {
	logger.Info("User authenticate: ", accessToken)
	return nil, fmt.Errorf("")
}

func (c *authService) GeneratePairToken(input model.User) (res dto.LoginRessponseDTO, err error) {
	accessTokenExpiry := time.Now().Add(time.Second * time.Duration(c.env.AccessTokenValiditySec))
	accessClaims := jwt.MapClaims{
		"sub":  input.ID.String(),
		"user": input.Username,
		"name": input.FullName,
		"exp":  accessTokenExpiry.Unix(),
		"iat":  time.Now().Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(c.env.RSAPrivateKey))
	if err != nil {
		logger.WithError(err).Error("Failed to generate access token")
		return dto.LoginRessponseDTO{}, err
	}

	refreshTokenExpiry := time.Now().Add(time.Second * time.Duration(c.env.RefreshTokenValiditySec))
	refreshClaims := jwt.MapClaims{
		"sub": input.ID.String(),
		"exp": refreshTokenExpiry.Unix(),
		"iat": time.Now().Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(c.env.RSAPrivateKey))
	if err != nil {
		logger.WithError(err).Error("Failed to generate refresh token")
		return dto.LoginRessponseDTO{}, err
	}

	return dto.LoginRessponseDTO{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (c *authService) ValidateAccessToken(tokenString string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.WithError(err).Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, network.NewUnauthorizedErr("Invalid credentials!", err)
		}

		return []byte(c.env.RSAPrivateKey), nil
	})

	if err != nil {
		logger.WithError(err).Error("Failed to parse token")
		return nil, network.NewUnauthorizedErr("Invalid credentials!", err)
	}

	if !token.Valid {
		return nil, network.NewUnauthorizedErr("Invalid credentials!", fmt.Errorf("Invalid credentials"))
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	logger.Error("failed to extract claims from token")
	return nil, network.NewUnauthorizedErr("Invalid credentials!", fmt.Errorf("Invalid credentials"))
}
