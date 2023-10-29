package usecase

import (
	"context"
	"log"
	"math/rand"

	"github.com/voyagesez/auservice/src/internals/repository"
	"github.com/voyagesez/auservice/src/internals/strategies"
)

type AuthUseCase interface {
	ExternalLogin(ctx context.Context, oauthProfile *strategies.OAuthProfile, callback func(ctx context.Context, oauthProfile *strategies.OAuthProfile) (string, error)) (string, error)
	ExternalRegister(ctx context.Context, oauthProfile *strategies.OAuthProfile) (string, error)
	PasswordLogin(email string, password string) (string, error)
	PasswordRegister(email string, password string) (string, error)
}

type AuthUserCaseImpl struct {
	authRepo repository.AuthRepository
}

func NewAuthUseCase(
	authRepo repository.AuthRepository,
) AuthUseCase {
	return &AuthUserCaseImpl{
		authRepo: authRepo,
	}
}

func (useCase *AuthUserCaseImpl) ExternalLogin(ctx context.Context, oauthProfile *strategies.OAuthProfile, callback func(ctx context.Context, oauthProfile *strategies.OAuthProfile) (string, error)) (string, error) {
	requestSuccess := rand.Intn(2) == 1
	if !requestSuccess {
		return callback(ctx, oauthProfile)
	}
	log.Println("login request success")
	return "", nil
}

func (useCase *AuthUserCaseImpl) ExternalRegister(ctx context.Context, oauthProfile *strategies.OAuthProfile) (string, error) {
	log.Println("register request success")
	return "", nil
}

func (useCase *AuthUserCaseImpl) PasswordLogin(email string, password string) (string, error) {
	return "", nil
}

func (useCase *AuthUserCaseImpl) PasswordRegister(email string, password string) (string, error) {
	return "", nil
}
