package cmd

import (
	"api-env-example/infra"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

func StartHttp(ctx context.Context, env infra.Config) {
	e := echo.New()

	e.GET("/health_check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Available")
	})

	e.GET("/env_check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, env)
	})

	err := e.Start(":8080")
	if err != nil {
		panic(err)
	}
}
