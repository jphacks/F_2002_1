package config

import (
	"fmt"
	"os"
)

// DSN はデータベースに接続するためのData Source Nameを返却します。
func DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	) + "?parseTime=true&collation=utf8mb4_bin"
}