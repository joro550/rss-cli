package data

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func ConnectToDatabase() (*gorm.DB, error) {
  Database, err := gorm.Open(sqlite.Open("rss-cli.db"), &gorm.Config{})
  if err != nil {
    return nil, fmt.Errorf("Failed to connect to database: %v", err)
  }

  // Migrate the schema
  Database.AutoMigrate(&Feed{})
  return Database, nil
}

