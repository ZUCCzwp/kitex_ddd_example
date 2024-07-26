package conf

import (
	"fmt"
	"testing"
)

type Config struct {
	Port  string
	Env   string
	Mysql struct {
		Host     string
		Port     int
		Database string
		Username string
		Password string
	}
	Redis struct {
		Host      string
		Port      int
		Password  string
		Databases int
	}
}

var config Config

func TestMustLoadYaml(t *testing.T) {
	MustLoad("config.yaml", &config)
	fmt.Println(config)
	fmt.Println(config.Mysql)
	fmt.Println(config.Redis)
}

func TestMustLoadToml(t *testing.T) {
	MustLoad("config.toml", &config)
	fmt.Println(config)
	fmt.Println(config.Mysql)
	fmt.Println(config.Redis)
}

func TestMustLoadJson(t *testing.T) {
	MustLoad("config.json", &config)
	fmt.Println(config)
	fmt.Println(config.Mysql)
	fmt.Println(config.Redis)
}
