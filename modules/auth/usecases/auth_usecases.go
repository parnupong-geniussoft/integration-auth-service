package usecases

import (
	"errors"
	"integration-auth-service/modules/auth/entities"
	"integration-auth-service/modules/auth/repositories"
	"integration-auth-service/pkg/utils"
	"time"

	"integration-auth-service/configs"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase interface {
	GetToken(p *entities.TokenRequest) (*entities.TokenResponse, error)
}

type authUsecase struct {
	Cfg      *configs.Configs
	AuthRepo repositories.AuthRepository
}

// Constructor
func NewAuthUsecase(cfg *configs.Configs, authRepo repositories.AuthRepository) AuthUsecase {
	return &authUsecase{
		Cfg:      cfg,
		AuthRepo: authRepo,
	}
}

func (u *authUsecase) GetToken(p *entities.TokenRequest) (*entities.TokenResponse, error) {

	key := p.ClientID + p.GrantType

	clientSecret := u.AuthRepo.GetClientSecretByClientId(key)
	if key == clientSecret {
		return nil, errors.New("unauthorized")
	}

	inputHash := utils.StringHash(p.ClientSecret)
	if clientSecret != inputHash {
		return nil, errors.New("unauthorized")
	}

	systemSource := u.AuthRepo.GetClientSystemSourceByClientId(key)
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"systemSource": systemSource,
			"exp":          time.Now().Add(time.Hour).Unix(),
		})

	t, err := token.SignedString([]byte(u.Cfg.Auth.OauthJwtSecret))
	if err != nil {
		return nil, err
	}

	return &entities.TokenResponse{
		AccessToken: t,
		ExpiresIn:   uint32(time.Hour.Seconds()),
		TokenType:   "Bearer",
	}, nil
}
