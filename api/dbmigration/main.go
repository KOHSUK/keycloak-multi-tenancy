package main

import (
	"api/identityaccess/infrastructure/gorm/connector"
	"api/identityaccess/infrastructure/gorm/repository/tenantrepo"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func main() {
	config := connector.PostgresConnectorConfig{
		DSN: "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable",
	}
	connector, err := connector.NewPostgresConnector(config)
	if err != nil {
		log.Fatal(err)
	}

	m := gormigrate.New(connector.GetDB(), gormigrate.DefaultOptions, []*gormigrate.Migration{{
		ID: "20240501120000",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().CreateTable(&tenantrepo.Tenant{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&tenantrepo.Tenant{})
		},
	},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
