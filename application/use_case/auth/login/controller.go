package login

import (
	"imp-backend/application/misc"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	loginService LoginService
}

func NewLoginHandler(
	loginService LoginService,
) LoginHandler {
	return LoginHandler{
		loginService: loginService,
	}
}

// @Summary      Show successfully login
// @Description  Login with username and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param request body LoginRequest true "Request"
// @Success      200  {object}  domain.LoginResponse
// @Router       /login [post]
func (h *LoginHandler) Login(c echo.Context) error {
	req := LoginRequest{}
	ctx := c.Request().Context()
	if err := c.Bind(&req); err != nil {
		misc.LogEf("AuthController - Login Error while binding request to json : ", err)
		c.JSON(http.StatusInternalServerError, misc.Response(err.Error(), nil))
		return err
	}

	if err := req.Validate(); err != nil {
		misc.LogEf("AuthController - Login Error validation : ", err.Error())
		c.JSON(http.StatusUnprocessableEntity, misc.NewValidatorError(err))
		return err
	}

	res, err := h.loginService.Login(ctx, &req)
	if err != nil {
		misc.LogEf("AuthController - Login Error while accessing service : ", err)
		c.JSON(http.StatusBadRequest, misc.Response(err.Error(), nil))
		return err
	}
	return c.JSON(http.StatusCreated, misc.Response("Success Login", res))
}
