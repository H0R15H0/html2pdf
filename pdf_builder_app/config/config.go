package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type config struct {
	Port                  int                   `env:"PORT" envDefault:"1323"`
	Html2PdfServiceConfig Html2PdfServiceConfig `envPrefix:"HTML2PDF_SERVICE_"`
	DBConfig              DBConfig              `envPrefix:"DB_"`
	AWSConfig             AWSConfig             `envPrefix:"AWS_"`
}

type Html2PdfServiceConfig struct {
	Origin             string `env:"ORIGIN,required"`
	WebhookMethod      string `env:"WEBHOOK_METHOD,required"`
	WebhookErrorUrl    string `env:"WEBHOOK_ERROR_URL,required"`
	WebhookErrorMethod string `env:"WEBHOOK_ERROR_METHOD,required"`
}

type DBConfig struct {
	Host string `env:"HOST,required"`
	Port int    `env:"PORT,required"`
	User string `env:"USER,required"`
	Pass string `env:"PASS,required" secret:"true"`
	Name string `env:"NAME,required"`
}

type AWSConfig struct {
	ID       string   `env:"ID,required"`
	Secret   string   `env:"SECRET,required"`
	Token    string   `env:"TOKEN,required"`
	Region   string   `env:"REGION,required"`
	S3Config S3Config `envPrefix:"S3_"`
}

type S3Config struct {
	Origin               string `env:"ORIGIN,required"`
	ForcePathStyle       bool   `env:"FORCE_PATH_STYPE" envDefault:"false"`
	PartialPdfBucketName string `env:"PARTIAL_PDF_BUCKET_NAME,required"`
}

func Get() *config {
	_ = godotenv.Load(".env")

	cnf := &config{}
	if err := env.Parse(cnf); err != nil {
		log.Fatal(err)
	}

	return cnf
}
