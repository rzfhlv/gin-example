package config

import (
	"log"

	"github.com/joho/godotenv"
	aMySQL "github.com/rzfhlv/gin-example/adapter/mysql"
	aRedis "github.com/rzfhlv/gin-example/adapter/redis"
	"github.com/rzfhlv/gin-example/pkg/hasher"
	pJwt "github.com/rzfhlv/gin-example/pkg/jwt"
)

type Config struct {
	MySQL *aMySQL.MySQL
	Redis *aRedis.Redis
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
		MySQL: mySql,
		Redis: redis,
		Pkg: Pkg{
			Hasher:  &hasher,
			JWTImpl: &jwtImpl,
		},
	}
}
