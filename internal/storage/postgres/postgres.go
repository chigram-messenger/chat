package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Params struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
	Sslmode  string
}

func Conn(p Params) (*sql.DB, error) {
	var op = "storage.postgres.Conn"

	db, err := sql.Open("pgx", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", p.Host, p.Port, p.User, p.DBName, p.Password, p.Sslmode))

	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}
	return db, nil
}
