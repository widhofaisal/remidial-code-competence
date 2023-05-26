package controller

import (
	"code/competence/config"
	"code/competence/middleware"
	"code/competence/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)



func Register(c echo.Context) error {
	var admin model.Admin

	if err_bind := c.Bind(&admin); err_bind != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	uuid := uuid.New()
	admin.ID = uuid.String()

	if err_insert := config.DB.Save(&admin).Error; err_insert != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "success register",
	})
}

func Login(c echo.Context) error {
	var admin model.Admin

	if err_bind := c.Bind(&admin); err_bind != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err_select := config.DB.Where("email=? AND password=?", admin.Email, admin.Password).First(&admin).Error; err_select != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	token, err_token := middleware.CreateToken(admin.ID, admin.Email, admin.Password)
	if err_token != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	
	id, err_claimsId := middleware.GetIDFromToken(token)
	if err_claimsId != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "success register",
		"token":   token,
		"id":      id,
	})

}
