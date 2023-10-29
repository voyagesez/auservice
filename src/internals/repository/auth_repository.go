package repository

import (
	"github.com/voyagesez/auservice/src/internals/db"
	"github.com/voyagesez/auservice/src/internals/strategies"
)

type AuthRepository interface {
	GetUserViaExternal(oauthProfile strategies.OAuthProfile) (string, error)
	CreateNewAccountViaExternal(oauthProfile strategies.OAuthProfile) (string, error)
}

type AuthRepositoryImpl struct {
	dbInstance *db.DatabaseInstance
}

func NewAuthRepository(dbInstance *db.DatabaseInstance) AuthRepository {
	return &AuthRepositoryImpl{
		dbInstance: dbInstance,
	}
}

func (repo *AuthRepositoryImpl) GetUserViaExternal(oauthProfile strategies.OAuthProfile) (string, error) {
	return "", nil
}

func (repo *AuthRepositoryImpl) CreateNewAccountViaExternal(oauthProfile strategies.OAuthProfile) (string, error) {
	return "", nil
}
