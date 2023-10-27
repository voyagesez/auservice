package handlers

import "net/http"

type UserHandler interface {
	GetMyProfile(w http.ResponseWriter, r *http.Request)
}

type UserHandlerImpl struct {
}

func NewUserHandlers() UserHandler {
	return &UserHandlerImpl{}
}

func (u *UserHandlerImpl) GetMyProfile(w http.ResponseWriter, r *http.Request) {
	successResponse(w, r, http.StatusOK, SuccessResponse{
		Message: "ok",
		Data: map[string]string{
			"test": "ok",
		},
	})
}
