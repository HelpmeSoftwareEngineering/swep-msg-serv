package infrastructure

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRESQL_CONNECTION")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//if err := db.AutoMigrate(&entity.Message{}, &entity.Notification{}); err != nil {
	//	log.Fatalf("AutoMigrate failed: %v", err)
	//	return nil, err
	//} else {
	//	log.Println("Migrated database successfully")
	//}
	return db, nil
}
