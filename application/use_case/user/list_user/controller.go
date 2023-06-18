package list_user

import (
	"imp-backend/application/misc"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ListUserHandler struct {
	listUserService ListUserService
}

func NewListUserHandler(
	listUserService ListUserService,
) ListUserHandler {
	return ListUserHandler{
		listUserService: listUserService,
	}
}

// List User function handler
func (h *ListUserHandler) ListUser(c echo.Context) error {
	req := ListUserRequest{}
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		misc.LogEf("UserController - ListUser Error while binding request to json : ", err)
		c.JSON(http.StatusInternalServerError, misc.Response(err.Error(), nil))
		return err
	}

	if err := req.Validate(); err != nil {
		misc.LogEf("UserController - ListUser Error validation : ", err.Error())
		c.JSON(http.StatusUnprocessableEntity, misc.NewValidatorError(err))
		return err
	}

	res, err := h.listUserService.ListUser(ctx, &req)
	if err != nil {
		misc.LogEf("UserController - ListUser Error while accessing service : ", err)
		c.JSON(http.StatusBadRequest, misc.Response(err.Error(), nil))
		return err
	}

	return c.JSON(http.StatusOK, misc.PaginationResponse("Success ListUser", res, req.Page, req.Limit))
}
