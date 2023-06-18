package misc

import (
	"net/http"
	"strings"

	validator "github.com/go-playground/validator/v10"
	base "github.com/refactory-id/go-core-package/response"
)

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BasePagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type Pagination struct {
	Message    string         `json:"message"`
	Data       interface{}    `json:"data"`
	Pagination BasePagination `json:"pagination"`
}

func NewValidatorError(err error) base.BaseResponse {
	var errors string

	for _, v := range err.(validator.ValidationErrors) {
		switch v.Tag() {
		case "min":
			errors = v.Field() + " Minimum " + v.Param()
		case "max":
			errors = v.Field() + " Maximum " + v.Param()
		case "email":
			errors = v.Field() + " must be a valid email "
		case "numeric":
			errors = v.Field() + " must be number "
		case "required":
			errors = v.Field() + " cannot be empty "
		}
	}

	return base.BaseResponse{Success: false, Message: errors}
}

func Response(msg string, data interface{}) BaseResponse {
	return BaseResponse{
		Message: msg,
		Data:    data,
	}
}

func PaginationResponse(msg string, data interface{}, page int, limit int) Pagination {
	return Pagination{
		Message: msg,
		Data:    data,
		Pagination: BasePagination{
			Page:  page,
			Limit: limit,
		},
	}
}

func ResponseErrorCode(err error) int {
	if strings.Contains(err.Error(), "already exist") {
		return http.StatusConflict
	}
	return http.StatusBadRequest
}
