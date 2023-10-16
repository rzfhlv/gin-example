package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	aMySQL "github.com/rzfhlv/gin-example/adapter/mysql"
	"github.com/rzfhlv/gin-example/pkg/hasher"
)

type Config struct {
	MySQL *sqlx.DB
	Pkg   Pkg
}

type Pkg struct {
	Hasher hasher.HashPassword
}

func Init() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed Load Env %v", err.Error())
	}

	mySql, err := aMySQL.New()
	if err != nil {
		log.Fatalf("Failed to MySQL connection %v", err.Error())
	}

	hasher := hasher.HasherPassword{}

	return &Config{
		MySQL: mySql.GetDB(),
		Pkg: Pkg{
			Hasher: &hasher,
		},
	}
}
