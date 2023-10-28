package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/voyagesez/auservice/src/utils"
)

type UserHandler interface {
	GetMyProfile(w http.ResponseWriter, r *http.Request)
	GetAnotherProfile(w http.ResponseWriter, r *http.Request)
}

type UserHandlerImpl struct {
}

func NewUserHandlers() UserHandler {
	return &UserHandlerImpl{}
}

func (u *UserHandlerImpl) GetMyProfile(w http.ResponseWriter, r *http.Request) {
	state := utils.RandomString(32)
	successResponse(w, r, http.StatusOK, SuccessResponse{
		Message: "your profile",
		Data:    "it is your profile: " + state,
	})
}

func (u *UserHandlerImpl) GetAnotherProfile(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	successResponse(w, r, http.StatusOK, SuccessResponse{
		Message: "your profile",
		Data:    "it is another profile " + slug,
	})
}
