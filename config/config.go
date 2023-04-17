package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Dbname   string `yaml:"dbname"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"SSLMode"`
	} `yaml:"database"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logrus.Print("gather config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			text := "Error reading configuration"
			help, _ := cleanenv.GetDescription(instance, &text)
			logrus.Println(help)
			logrus.Fatal(err)
		}
	})
	return instance
}
