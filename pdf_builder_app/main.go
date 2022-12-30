package main

import (
	"fmt"
	"net/http"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/config"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/infra/postgresql"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/infra/postgresql/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/interfaces/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	cnf := config.Get()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	db, err := postgresql.NewPostgresqlDB(cnf.DBConfig.User, cnf.DBConfig.Pass, cnf.DBConfig.Host, cnf.DBConfig.Port, cnf.DBConfig.Name)
	if err != nil {
		panic(err)
	}
	handlers.NewPdfHandler(repositories.NewPdfRepo(db))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cnf.Port)))
}
