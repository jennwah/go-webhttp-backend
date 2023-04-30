package mysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	DriverName = "mysql"
)

const (
	dbWaitTimeout    = 14400
	dbMaxConnections = 50
)

type Config struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func (c Config) string() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&allowCleartextPasswords=1&tls=false",
		c.User, c.Pass, c.Host, c.Port, c.Name)
}

// NewClient returns DB connection.
func NewClient(c Config) (*sqlx.DB, error) {
	conn, err := sqlx.Connect(DriverName, c.string())
	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(time.Second * dbWaitTimeout) // 1/2 of a mysql wait_timeout (28800)
	conn.SetMaxOpenConns(dbMaxConnections)               // 1/4 of a mysql max_connections (200)

	return conn, nil
}
