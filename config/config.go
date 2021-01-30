package config

// Configurations exported
type Configurations struct {
	Database DatabaseConfiguration
}

// DatabaseConfigurations exported
type DatabaseConfiguration struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBURL      string
	DBPort     string
}
