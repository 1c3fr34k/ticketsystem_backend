package database

import (
	"github.com/1c3fr34k/ticketsystem_backend/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, pass, database, port, sslmode)
	// dsn := "host=localhost user=postgres password=Start123! dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	Migrate(db)

	return db, err
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Ticket{})
}
