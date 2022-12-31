package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type config struct {
	Port                  int      `env:"PORT" envDefault:"1323"`
	Html2PdfServiceOrigin string   `env:"HTML2PDF_SERVICE_ORIGIN,required"`
	DBConfig              DBConfig `envPrefix:"DB_"`
}

type DBConfig struct {
	Host string `env:"HOST,required"`
	Port int    `env:"PORT,required"`
	User string `env:"USER,required"`
	Pass string `env:"PASS,required" secret:"true"`
	Name string `env:"NAME,required"`
}

func Get() *config {
	_ = godotenv.Load(".env")

	cnf := &config{}
	if err := env.Parse(cnf); err != nil {
		log.Fatal(err)
	}

	return cnf
}
