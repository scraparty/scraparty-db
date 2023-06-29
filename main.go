package scrapartydb

import (
	"errors"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (gorm.DB, error) {
	connection_string := os.Getenv("CONNECTION_STRING")

	if connection_string == "" {
    return gorm.DB{}, errors.New("the connection_string environment variable is not present")
	}

	connection_string = strings.ReplaceAll(
		connection_string,
		"\n", 
		"",
	)

	db, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return gorm.DB{}, err
	}

	return *db, nil
}
