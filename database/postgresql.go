package database

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

// Config defines the environment variables necessary for the app to run
type Config struct {
	PostgresHost       string `mapstructure:"POSTGRES_HOST"`
	GolangPostgresHost string `mapstructure:"GOLANG_POSTGRES_HOST"`
	PostgresPort       string `mapstructure:"POSTGRES_PORT"`
	PostgresUser       string `mapstructure:"POSTGRES_USER"`
	PostgresDBName     string `mapstructure:"POSTGRES_DB"`
	PostgresPassword   string `mapstructure:"POSTGRES_PASSWORD"`
}

// LoadConfig read the .env file
func LoadConfig(path string) (config Config, err error) {
	// Read file path
	viper.AddConfigPath(path)
	// set config file and path
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	// watching changes in app.env
	viper.AutomaticEnv()
	// reading the config file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// SetUpPostgres set up Postgresql with environment variables
func SetUpPostgres() *Postgres {

	// load app.env file data to struct
	config, err := LoadConfig(".")

	// handle errors
	if err != nil {
		log.Fatalf("can't load environment app.env: %v", err)
	}

	// Host: using the host for docker-compose
	postgres, err := NewPostgres(config.GolangPostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresDBName, config.PostgresPassword)

	// Host: for local connection
	// postgres, err := NewPostgres(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresDBName, config.PostgresPassword)

	if err != nil {
		log.Fatal(err.Error())
	}
	return postgres
}

// NewPostgres constructor for postgres
func NewPostgres(host, port, user, dbname, password string) (*Postgres, error) {
	connStr := "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable"
	db, err := sql.Open("postgres", fmt.Sprintf(connStr, host, port, user, dbname, password))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}
