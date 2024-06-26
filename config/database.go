package config

import "fmt"

// DatabaseConfig stores the generic configuration for database.
type DatabaseConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	Name        string
	MaxPoolSize int
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:        viperString("db_host", "127.0.0.1"),
		Port:        viperInt("db_port", 5432),
		Name:        viperString("database", "postgres"),
		Username:    viperString("db_user", "postgres"),
		Password:    viperString("db_password", "postgres"),
		MaxPoolSize: viperInt("db_maxpoolsize", 50),
	}
}

func (dc DatabaseConfig) String() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable", dc.Name, dc.Username, dc.Password, dc.Host, dc.Port)
}
