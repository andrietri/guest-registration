package migration

import (
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/gorm"
)

// func MigrationRubenv
func MigrationRubenv(db *gorm.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "database/migrationSchema",
	}
	migrate.SetTable("migrations")

	sql, _ := db.DB()
	n, err := migrate.Exec(sql, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println("Error occcured:", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
