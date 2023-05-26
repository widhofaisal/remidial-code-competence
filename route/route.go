package route

import (
	"code/competence/constant"
	"code/competence/controller"
	"code/competence/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.Use(middleware.MiddlewareLogging)
	e.HTTPErrorHandler = middleware.ErrorHandler

	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)

	authJWT := e.Group("")
	authJWT.Use(mid.JWT([]byte(constant.SECRET_JWT)))
	authJWT.POST("/items", controller.Add_items)
	authJWT.GET("/items", controller.Get_items)
	authJWT.GET("/items/:id", controller.Get_item_by_id)
	authJWT.PUT("/items/:id", controller.Update_item)
	authJWT.DELETE("/items/:id", controller.Delete_item)
	authJWT.GET("/items/category/:category_id", controller.Get_items_by_category)

	return e
}
