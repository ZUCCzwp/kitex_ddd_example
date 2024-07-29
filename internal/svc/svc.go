package svc

import (
	"database/sql"
	"fmt"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/config"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/domain/repository"
	"github.com/ZUCCzwp/kitex_ddd_example/internal/infrastructure/persistence"
	"os"
)

type ServiceContext struct {
	config   config.Config
	userRepo repository.UserRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn, err := NewDBConnection()
	if err != nil {
		panic(err.Error())
	}
	defer mysqlConn.Close()
	return &ServiceContext{
		config:   c,
		userRepo: persistence.NewUserRepository(mysqlConn),
	}
}

// NewDBConnection returns initialized sql.DB
func NewDBConnection() (*sql.DB, error) {
	user := getEnvWithDefault("DB_USER", "root")
	password := getEnvWithDefault("DB_PASSWORD", "")
	host := getEnvWithDefault("DB_HOST", "localhost")
	port := getEnvWithDefault("DB_PORT", "3306")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/ddd_sample?parseTime=true", user, password, host, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getEnvWithDefault(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}
