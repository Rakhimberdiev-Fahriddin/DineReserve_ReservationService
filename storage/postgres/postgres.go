package postgres

import (
	"database/sql"
	"fmt"
	"reservation-service/config"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	cfg := config.Load()
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USERNAME, cfg.DB_DATABASE, cfg.DB_PASSWORD, )

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	
	if err = db.Ping(); err != nil {
		return nil, err
	}
	
	return db, err
}
