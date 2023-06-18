package signup

import "github.com/go-playground/validator/v10"

type SignupRequest struct {
	Username string `json:"username" validate:"required,min=2"`
	Password string `json:"password" validate:"required,min=5"`
	Fullname string `json:"fullname"`
}

type SignupResponse struct{}

func (c *SignupRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
