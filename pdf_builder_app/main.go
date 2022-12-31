package main

import (
	"fmt"
	"net/http"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/config"
	"github.com/H0R15H0/html2pdf/pdf_builder_app/infra/gotenberg"
	repositories2 "github.com/H0R15H0/html2pdf/pdf_builder_app/infra/gotenberg/repositories"
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
	// TODO: set webhook urls.
	html2PdfClient := gotenberg.NewGotenbergClient(cnf.Html2PdfServiceOrigin, "", "")
	userRepo := repositories.NewUserRepo(db)
	pdfRepo := repositories.NewPdfRepo(db)
	partialPdfRepo := repositories.NewPartialPdfRepo(db)
	html2PdfRepo := repositories2.NewHtml2PdfRepo(html2PdfClient)
	userUsecase := usecases.NewUserUsecase(userRepo)
	pdfUsecase := usecases.NewPdfUsecase(pdfRepo, userRepo)
	partialPdfUsecase := usecases.NewPartialPdfUsecase(partialPdfRepo, html2PdfRepo)
	userHandler := handlers.NewUserHandler(userUsecase)
	usersPdfHandler := handlers.NewUsersPdfHandler(pdfUsecase)
	partialPdfHandler := handlers.NewPartialPdfHandler(partialPdfUsecase)

	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/users", userHandler.CreateUser)
	e.POST("/users/:id/pdfs", usersPdfHandler.InitializeUsersPdf)
	e.POST("/users/:user_id/pdfs/:pdf_id/partials", partialPdfHandler.Create)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cnf.Port)))
}
