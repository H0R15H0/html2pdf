package database

import (
	"fmt"
)

type Config struct {
	User string
	Pass string
	Host string
	Port int
	Name string
}

func (c Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.User, c.Pass, c.Host, c.Port, c.Name)
}
