package config

type Environment struct {
	MySQL
	MigrateDatabase
}

type MySQL struct {
	User            string
	Pass            string
	Host            string
	Port            int
	DBName          string
	MigrateDatabase string
}

type MigrateDatabase struct {
	MigrateStatus bool
}
