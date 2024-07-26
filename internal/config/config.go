package config

type Config struct {
	Name     string `yaml:"Name"`
	Port     int    `yaml:"Port"`
	ListenOn string `yaml:"ListenOn"`
	Mode     string `yaml:"Mode"`
}
