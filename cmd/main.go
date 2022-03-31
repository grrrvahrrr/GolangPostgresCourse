package main

import (
	"CourseWork/internal/apichi"
	"CourseWork/internal/apichi/openapichi"
	"CourseWork/internal/config"
	"CourseWork/internal/database/pgxstorage"
	"CourseWork/internal/dbbackend"
	"CourseWork/internal/logging"
	"CourseWork/internal/server"
	"context"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
)

//go:embed config/config.env
var cfg string

func main() {
	//Generate random seed
	rand.Seed(time.Now().UnixNano())

	//Creating Context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	//Logging
	f, err := logging.LogErrors("error.log")
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer f.Close()

	//Load Config
	cfg, err := config.LoadConfig(cfg)
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	//Creating Storage
	const dsn = "postgres://bituser:bit@localhost:5433/bitmedb?sslmode=disable"
	// udf, err := database.NewPgStorage(dsn)
	// if err != nil {
	// 	log.Fatal("Error creating database files: ", err)
	// }

	pgxcfg, err := pgxstorage.NewPgxConfig(dsn, 50, 10, 1, 2)
	if err != nil {
		log.Fatal("Error creating database config: ", err)
	}
	udf, err := pgxstorage.NewPgxStorage(ctx, pgxcfg)
	if err != nil {
		log.Fatal("Error creating database files: ", err)
	}

	dbbe := dbbackend.NewDataStorage(udf)

	//Creating router and server
	hs := apichi.NewHandlers(dbbe)
	rt := openapichi.NewOpenApiRouter(hs)
	srv := server.NewServer(":8000", rt, cfg)

	//Starting
	srv.Start(dbbe)

	fmt.Println("Hello, Bitme!")

	//Shutting down
	<-ctx.Done()

	srv.Stop()
	cancel()
	udf.Close()

	fmt.Print("Server shutdown.")
}
