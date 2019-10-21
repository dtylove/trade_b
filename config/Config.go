package config

import (
	"encoding/json"
	"io/ioutil"
)

var Conf Config

type Config struct {
	DataSource DataSourceConfig `json:"dataSource"`
	Redis      RedisConfig      `json:"redis"`
	Kafka      KafkaConfig      `json:"kafka"`
	PushServer PushServerConfig `json:"pushServer"`
	RestServer RestServerConfig `json:"restServer"`
	JwtSecret  string           `json:"jwtSecret"`
}

type DataSourceConfig struct {
	DriverName string `json:"driverName"`
	Addr       string `json:"addr"`
	Database   string `json:"database"`
	User       string `json:"user"`
	Password   string `json:"password"`
	Migrate    bool   `json:"migrate"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

type KafkaConfig struct {
	Brokers []string `json:"brokers"`
}

type PushServerConfig struct {
	Addr string `json:"addr"`
	Path string `json:"path"`
}

type RestServerConfig struct {
	Addr string `json:"addr"`
}

func init() {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &Conf)
	if err != nil {
		panic(err)
	}
}
