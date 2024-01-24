package dbConfig

import (
	"fmt"
	"github.com/hudayberdipolat/golang-auth-jwtToken/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	dbConfig config.Config
}

func NewDBConnection(dbConfig config.Config) DBConfig {
	return DBConfig{
		dbConfig: dbConfig,
	}
}

func (dbConfig DBConfig) DBConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.dbConfig.DbConfig.DbHost,
		dbConfig.dbConfig.DbConfig.DbUser,
		dbConfig.dbConfig.DbConfig.DbPassword,
		dbConfig.dbConfig.DbConfig.DbName,
		dbConfig.dbConfig.DbConfig.DbPort,
		dbConfig.dbConfig.DbConfig.DbSllMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
