package database

import (
	"fmt"
	"github.com/samber/lo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connect(databaseName string) (*gorm.DB, error) {
	// We'll prefer explicit credentials if they're present, and fallback to crossplane defaults
	host := lo.Ternary(os.Getenv("DB_HOST") != "", os.Getenv("DB_HOST"), os.Getenv("endpoint"))
	port := lo.Ternary(os.Getenv("DB_PORT") != "", os.Getenv("DB_PORT"), os.Getenv("port"))
	username := lo.Ternary(os.Getenv("DB_USERNAME") != "", os.Getenv("DB_USERNAME"), os.Getenv("username"))
	password := lo.Ternary(os.Getenv("DB_PASSWORD") != "", os.Getenv("DB_PASSWORD"), os.Getenv("password"))

	return gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, databaseName)))
}
