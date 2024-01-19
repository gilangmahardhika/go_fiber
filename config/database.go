package config

type DatabaseConfig struct {
	User         string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

func ConnectDB(database *DatabaseConfig) string {
	return database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DatabaseName + "?parseTime=true"
}
