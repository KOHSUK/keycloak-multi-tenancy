package connector

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnector struct {
	db *gorm.DB
}

func NewPostgresConnector(dsn string) (*PostgresConnector, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresConnector{db: db}, nil
}

func (p *PostgresConnector) GetDB() *gorm.DB {
	return p.db
}
