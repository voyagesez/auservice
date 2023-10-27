package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func successResponse(w http.ResponseWriter, r *http.Request, code int, data SuccessResponse) {
	render.Status(r, code)
	render.JSON(w, r, data)
}

func errorResponse(w http.ResponseWriter, r *http.Request, code int, data ErrorResponse) {
	render.Status(r, code)
	render.JSON(w, r, data)
}
