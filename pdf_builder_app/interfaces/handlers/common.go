package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIError struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Code    string      `json:"code"`
}

func JsonError(c echo.Context, err error, a *APIError) error {
	// 通信自体は完了しているので200
	c.JSON(http.StatusOK, a)
	return err
}
