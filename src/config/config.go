package config

func Init() {
	env()
	connectDatabase()
	autoMigrate()
}
