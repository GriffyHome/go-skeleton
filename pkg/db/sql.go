package db

import (
	"fmt"

	"github.com/GriffyHome/go-skeleton/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type sqlConnection struct {
	connection *gorm.DB
}

func NewSQLDB() (*sqlConnection, error) {
	dbConfig := config.DBConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", dbConfig.Username, dbConfig.Password, dbConfig.URL, dbConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // Disable struct cache at the session level
	})
	if err != nil {
		return nil, err
	}
	return &sqlConnection{connection: db}, nil
}

func (d *sqlConnection) GetDatabase() *gorm.DB {
	return d.connection
}

func (d *sqlConnection) Migrate(developmentMode bool) {

}
