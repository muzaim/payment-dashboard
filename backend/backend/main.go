package main

import (
	"database/sql"
	"log"
	"time"

	internaldb "github.com/durianpay/fullstack-boilerplate/internal/db"

	"github.com/durianpay/fullstack-boilerplate/internal/api"
	"github.com/durianpay/fullstack-boilerplate/internal/config"
	ah "github.com/durianpay/fullstack-boilerplate/internal/module/auth/handler"
	ar "github.com/durianpay/fullstack-boilerplate/internal/module/auth/repository"
	au "github.com/durianpay/fullstack-boilerplate/internal/module/auth/usecase"
	srv "github.com/durianpay/fullstack-boilerplate/internal/service/http"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_ = godotenv.Load()

	db, err := sql.Open("sqlite3", "dashboard.db?_foreign_keys=1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// migrate tables
	if err := internaldb.Migrate(db); err != nil {
		log.Fatal(err)
	}

	// seed data
	if err := internaldb.Seed(db); err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 5)

	JwtExpiredDuration, err := time.ParseDuration(config.JwtExpired)
	if err != nil {
		panic(err)
	}

	userRepo := ar.NewUserRepo(db)
	authUC := au.NewAuthUsecase(userRepo, config.JwtSecret, JwtExpiredDuration)
	authH := ah.NewAuthHandler(authUC)

	apiHandler := &api.APIHandler{
		Auth: authH,
		DB:   db,
	}

	server := srv.NewServer(apiHandler, config.OpenapiYamlLocation)

	addr := config.HttpAddress
	log.Printf("starting server on %s", addr)
	server.Start(addr)
}
