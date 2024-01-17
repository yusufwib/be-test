package infrastructure_http

import (
	"net/http"

	"github.com/yusufwib/be-test/infrastructure"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (httpServer *HttpServer) PrepareMiddleware(app *infrastructure.App) {
	httpServer.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	httpServer.Echo.Use(middleware.Recover())

	httpServer.Echo.HTTPErrorHandler = func(err error, ctx echo.Context) {
		if ctx.Response().Committed {
			return
		}
		code := http.StatusInternalServerError
		he, ok := err.(*echo.HTTPError)
		if ok {
			code = he.Code
		}
		errMSg := map[string]interface{}{
			"data":    nil,
			"error":   err,
			"message": err.Error(),
		}
		ctx.JSON(code, errMSg)
	}
}
