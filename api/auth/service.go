package auth

import (
	"app/api/auth/dto"
	"app/api/auth/model"
	"app/arch/network"
	"app/arch/postgres"
	"app/config"
	"app/utils"

	logger "github.com/sirupsen/logrus"
)

type AuthService interface {
	RegisterUser(dto.RegistrationDTO) (user *model.User, err error)
	Login(dto.LoginDTO) (res dto.LoginRessponseDTO, err error)
}

type authService struct {
	network.BaseService
	db  postgres.Database
	env *config.Env
	// rsaPrivateKey *rsa.PrivateKey
	// rsaPublicKey  *rsa.PublicKey
}

func CreateAuthService(db postgres.Database, env *config.Env) AuthService {
	// privatePem, err := utils.LoadPEMFileInto(env.RSAPrivateKeyPath)
	// if err != nil {
	// 	panic(err)
	// }
	// rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePem)
	// if err != nil {
	// 	panic(err)
	// }

	// publicPem, err := utils.LoadPEMFileInto(env.RSAPublicKeyPath)
	// if err != nil {
	// 	panic(err)
	// }

	// rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPem)
	// if err != nil {
	// 	panic(err)
	// }

	logger.WithFields(logger.Fields{
		"rsaPublicKey":  env.RSAPublicKey,
		"rsaPrivateKey": env.RSAPublicKey,
	}).Info("CONFIGURATION")

	if env.RSAPublicKey != "" {
		logger.Info("GET RSAPublicKey success")
	}
	return &authService{
		BaseService: network.NewBaseService(),
		db:          db.GetInstance(),
		env:         env,
		// rsaPrivateKey: rsaPrivateKey,
		// rsaPublicKey:  rsaPublicKey,
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
	// c.db.GetInstance().Where()
	// tokens, err := c.GeneratePairToken(&model.User{})
	return dto.LoginRessponseDTO{
		AccessToken:  "",
		RefreshToken: "",
	}, nil
}

func (c *authService) GeneratePairToken(input model.User) (res dto.LoginRessponseDTO, err error) {
	return dto.LoginRessponseDTO{}, nil
}
