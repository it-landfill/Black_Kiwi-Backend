package black_kiwi_db_handler

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

//https://app.swaggerhub.com/apis/ITLandfill/Black-Kiwi/1.0.2

func InitConnectionPool() *pgxpool.Pool {
	dbUrl := "postgres://black-kiwi_administrator:6whuUYTEhyA2ShR35@postgis:5432/black-kiwi"
	conn, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Connection to database established")
	}

	return conn
}