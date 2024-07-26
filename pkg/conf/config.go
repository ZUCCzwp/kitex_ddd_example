package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// Load 从文件加载配置到 v，支持 .json, .yaml 和 .yml 文件。
func Load(file string, v any) error {
	viper.SetConfigFile(file)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file, %s", err)
	}

	if err := viper.Unmarshal(v); err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}

	return nil
}

// MustLoad 从文件加载配置到 v，支持 .json, .yaml 和 .yml 文件，如果加载失败则退出。
func MustLoad(path string, v any) {
	if err := Load(path, v); err != nil {
		log.Fatalf("error: config file %s, %s", path, err.Error())
	}
}
