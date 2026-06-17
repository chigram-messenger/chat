package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env"`
	GRPCServer `yaml:"grpc-server"`
	Postgres   `yaml:"postgres"`
}
type GRPCServer struct {
	Host string `yaml:"host"`
	Port int64  `yaml:"port"`
}
type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"db_name"`
	Password string `yaml:"password" env:"DB_POSTGRES_PASSWORD"`
	Sslmode  string `yaml:"sslmode"`
}

func MustLoad() *Config {
	url := fetchPathURL()
	if url == "" {
		panic("empty path config")
	}
	cfg, err := LoadWithUrl(url)
	if err != nil {

		panic("dont reading config, err: " + err.Error())
	}
	return cfg

}
func LoadWithUrl(url string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(url, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
func fetchPathURL() string {
	var url string
	flag.StringVar(&url, "config", "", "path is config file")
	flag.Parse()
	if url == "" {
		url = os.Getenv("CONFIG_PATH")
	}
	fmt.Println(url)
	return url
}
