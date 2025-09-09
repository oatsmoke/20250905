package env

import (
	"fmt"
	"log"
	"os"

	"github.com/oatsmoke/20250905/internal/lib/logger"
)

const (
	HttpPort    = "HTTP_PORT"
	PostgresDsn = "POSTGRES_DSN"
)

func GetHttpPort() string {
	return get(HttpPort)
}

func GetPostgresDsn() string {
	return get(PostgresDsn)
}

func get(key string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	} else {
		switch key {
		case HttpPort:
			message(HttpPort)
			return ":8080"
		case PostgresDsn:
			message(PostgresDsn)
			return "postgres://root:password@localhost:5432/postgres?sslmode=disable"
		default:
			log.Printf("%s not found\n", key)
			return ""
		}
	}
}

func message(key string) {
	logger.Info(fmt.Sprintf("%s not set, set default value", key))
}
