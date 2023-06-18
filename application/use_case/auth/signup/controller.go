package signup

import (
	"imp-backend/application/misc"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SignupHandler struct {
	signupService SignupService
}

func NewSignupHandler(
	signupService SignupService,
) SignupHandler {
	return SignupHandler{
		signupService: signupService,
	}
}

// @Summary      Show successfully signup
// @Description  Signup with username, password and username
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param request body SignupRequest true "Request"
// @Success      200  {object}  domain.SignupResponse
// @Router       /signup [post]
func (h *SignupHandler) Signup(c echo.Context) error {
	req := SignupRequest{}
	ctx := c.Request().Context()
	if err := c.Bind(&req); err != nil {
		misc.LogEf("AuthController - Signup Error while binding request to json : ", err)
		c.JSON(http.StatusInternalServerError, misc.Response(err.Error(), nil))
		return err
	}

	if err := req.Validate(); err != nil {
		misc.LogEf("AuthController - Signup Error validation : ", err.Error())
		c.JSON(http.StatusUnprocessableEntity, misc.NewValidatorError(err))
		return err
	}

	err := h.signupService.Signup(ctx, &req)
	if err != nil {
		misc.LogEf("AuthController - Login Error while accessing service : ", err)
		c.JSON(misc.ResponseErrorCode(err), misc.Response(err.Error(), nil))
		return err
	}

	return c.JSON(http.StatusOK, misc.Response("Success Signup", nil))
}
