package postgresql

import (
	"database/sql"

	"github.com/H0R15H0/html2pdf/pdf_builder_app/database"
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

	return sql.Open("postgres", dsn)
}
