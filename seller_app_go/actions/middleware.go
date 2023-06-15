package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

const (
	DummyToken = "mydummyToken"
)

// MiddleWareForValidateUserToken ...
func MiddleWareForValidateUserToken(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) (err error) {
		if c.Request().Header.Get("authorization") != DummyToken {
			return c.Render(http.StatusUnauthorized, r.JSON(map[string]string{"message": "Invalid Token"}))
		}
		return next(c)
	}
}
