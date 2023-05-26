package controller

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"code/competence/config"
	"code/competence/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Add_items(c echo.Context) error {
	var item model.Item

	if err_bind := c.Bind(&item); err_bind != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	uuid := uuid.New()
	item.ID = uuid.String()

	if err_insert := config.DB.Save(&item).Error; err_insert != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "success add item",
	})
}

func Get_items(c echo.Context) error {
	var items []model.Item
	name := c.QueryParam("keywords")
	fmt.Println(name)

	if err_select := config.DB.Where("name LIKE ?", "%"+name+"%").Find(&items).Error; err_select != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// var datas []model.Response
	// for _, item := range items {
	// 	data := model.Response{
	// 		ID : item.ID,
	// 		Name:        item.Name,
	// 		Description: item.Description,
	// 		Stock:       item.Stock,
	// 		Price:       item.Price,
	// 		CategoryID:  item.CategoryID,
	// 	}

	// 	datas = append(datas, data)
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "success get items",
		"data":    items,
	})
}

func Get_item_by_id(c echo.Context) error {
	var item model.Item

	if err_select := config.DB.First(&item).Error; err_select != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "success get item by id",
		"data":    item,
	})
}

func Update_item(c echo.Context) error {
	var item model.Item
	id, _ := strconv.Atoi(c.Param("id"))

	if err_select := config.DB.First(&item, id).Error; err_select != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if err_bind := c.Bind(&item); err_bind != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err_update := config.DB.Save(&item).Error; err_update != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "success update item",
	})
}

func Delete_item(c echo.Context) error {
	var item model.Item
	id, _ := strconv.Atoi(c.Param("id"))

	if err_select := config.DB.First(&item, id).Error; err_select != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if err_bind := c.Bind(&item); err_bind != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err_update := config.DB.Delete(&item).Error; err_update != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "success delete item",
	})
}

func Get_items_by_category(c echo.Context) error {
	var items []model.Item
	category_id, _ := strconv.Atoi(c.Param("category_id"))

	if err_select := config.DB.Where("category_id=?", category_id).Find(&items).Error; err_select != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "success get item by id",
		"data":    items,
	})
}

// func Get_items_by_name(c echo.Context) error {
// 	name := c.QueryParam("name")

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"status":  200,
// 		"message": "success get item by id",
// 		"data":    name,
// 	})
// }
