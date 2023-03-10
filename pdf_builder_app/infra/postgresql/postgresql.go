package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/database"
	_ "github.com/lib/pq"
)

func NewPostgresqlDB(user string, pass string, host string, port int, name string) (*sql.DB, error) {
	dbConfig := &database.Config{
		User: user,
		Pass: pass,
		Host: host,
		Port: port,
		Name: name,
	}

	dsn := dbConfig.DSN()

	return sql.Open("postgres", fmt.Sprintf("postgres://%s?sslmode=disable", dsn))
}
