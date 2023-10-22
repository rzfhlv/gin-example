package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	aMySQL "github.com/rzfhlv/gin-example/adapter/mysql"
	aRedis "github.com/rzfhlv/gin-example/adapter/redis"
	"github.com/rzfhlv/gin-example/pkg/hasher"
	pJwt "github.com/rzfhlv/gin-example/pkg/jwt"
)

type Config struct {
	MySQL *sqlx.DB
	Redis *redis.Client
	Pkg   Pkg
}

type Pkg struct {
	Hasher  hasher.HashPassword
	JWTImpl pJwt.JWTInterface
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

	redis, err := aRedis.New()
	if err != nil {
		log.Fatalf("Failed to MySQL connection %v", err.Error())
	}

	hasher := hasher.HasherPassword{}
	jwtImpl := pJwt.JWTImpl{}

	return &Config{
		MySQL: mySql.GetDB(),
		Redis: redis.GetClient(),
		Pkg: Pkg{
			Hasher:  &hasher,
			JWTImpl: &jwtImpl,
		},
	}
}
