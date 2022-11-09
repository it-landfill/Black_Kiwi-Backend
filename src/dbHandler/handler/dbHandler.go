package black_kiwi_db_handler

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

//https://app.swaggerhub.com/apis/ITLandfill/Black-Kiwi/1.0.2

func InitConnectionPool() *pgxpool.Pool {

	var dbUrl string

	if (os.Getenv("Black_Kiwi_ENV") == "dev") {
		dbUrl = "postgres://black-kiwi_administrator:6whuUYTEhyA2ShR35@127.0.0.1:5432/black-kiwi"
	} else {
		dbUrl = "postgres://black-kiwi_administrator:6whuUYTEhyA2ShR35@postgis:5432/black-kiwi"
	}

	log.WithFields(log.Fields{"dbUrl": dbUrl}).Info("Generating DB connection pool")

	conn, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to generate connection pool: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Connection pool generated")
	}

	return conn
}