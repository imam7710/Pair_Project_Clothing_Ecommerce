package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // <--- Tanda _ ini WAJIB ada
)

func ConnectDB() (*sql.DB, error) {
	// Pastikan DSN-nya benar
	dsn := "root:Im@m981999@tcp(127.0.0.1:3306)/clothing_ecommerce_db?parseTime=true&tls=skip-verify"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal terhubung ke MySQL: %v", err)
	}

	return db, nil
}
