package database

import (
	"fmt"

	"github.com/andrietri/guest-registration/database/migration"
	"github.com/erajayatech/go-helper"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// newPGConnection postgre db connection
func NewDBConnection() *gorm.DB {
	db := pgConnect()

	if helper.MustGetEnv("MIGRATE_DATABASE") == "true" {
		migration.MigrationRubenv(db)
	}

	return db
}

func Migration(db gorm.DB) {
	migration.MigrateExec(&db)
}

// pgConnect Connection to database postgre
func pgConnect() *gorm.DB {
	var (
		dbUser  = helper.MustGetEnv("DB_USER")
		dbPass  = helper.MustGetEnv("DB_PASSWORD")
		dbHost  = helper.MustGetEnv("DB_HOST")
		dbName  = helper.MustGetEnv("DB_NAME")
		dbPort  = helper.MustGetEnv("DB_PORT")
		TZ      = helper.MustGetEnv("TZ")
		sslMode = helper.MustGetEnv("SSL_MODE")
	)

	// dsn
	dsn := fmt.Sprintf(`
		host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s`,
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
		sslMode,
		TZ,
	)

	log.Print("dsn:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		// set without default transaction
		// will call manually per-case query
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Warn("Connected to database Failed:", err)
	}
	log.Warn("Connected to database")

	return db
}
