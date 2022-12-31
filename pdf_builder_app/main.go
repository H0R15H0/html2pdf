package main

import (
	"fmt"
	"net/http"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/config"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/infra/postgresql"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/infra/postgresql/repositories"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/interfaces/handlers"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cnf := config.Get()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	db, err := postgresql.NewPostgresqlDB(cnf.DBConfig.User, cnf.DBConfig.Pass, cnf.DBConfig.Host, cnf.DBConfig.Port, cnf.DBConfig.Name)
	if err != nil {
		panic(err)
	}
	userRepo := repositories.NewUserRepo(db)
	pdfRepo := repositories.NewPdfRepo(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	pdfUsecase := usecases.NewPdfUsecase(pdfRepo, userRepo)
	usersPdfHandlers := handlers.NewUsersPdfHandler(pdfUsecase)
	userHandler := handlers.NewUserHandler(userUsecase)

	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/users", userHandler.CreateUser)
	e.POST("/users/:id/pdfs", usersPdfHandlers.InitializeUsersPdf)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cnf.Port)))
}
