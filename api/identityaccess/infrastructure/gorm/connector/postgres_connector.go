package connector

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnector struct {
	db *gorm.DB
}

type PostgresConnectorConfig struct {
	DSN string
}

func NewPostgresConnector(config PostgresConnectorConfig) (*PostgresConnector, error) {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresConnector{db: db}, nil
}

func (p *PostgresConnector) GetDB() *gorm.DB {
	return p.db
}
