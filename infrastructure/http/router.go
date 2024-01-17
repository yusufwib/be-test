package infrastructure_http

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/yusufwib/be-test/docs"
	"github.com/yusufwib/be-test/infrastructure"
)

// @title Employee API
// @version 1.0
func (httpServer *HttpServer) PrepareRoute(app *infrastructure.App) {
	dependency := infrastructure.NewDependency(app.Context, app.Logger, app.Validator, app.Cfg, app.Database)

	httpServer.Echo.GET("/employees", dependency.EmployeeHandler.FindAll)
	httpServer.Echo.GET("/employees/:id", dependency.EmployeeHandler.FindByID)
	httpServer.Echo.POST("/employees", dependency.EmployeeHandler.Create)
	httpServer.Echo.PUT("/employees/:id", dependency.EmployeeHandler.Update)
	httpServer.Echo.DELETE("/employees/:id", dependency.EmployeeHandler.Delete)

	httpServer.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
