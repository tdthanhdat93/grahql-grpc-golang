package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	APIConf  Server     `json:"api"`
	GRPCConf GRPCServer `json:"grpc"`
}

type GRPCServer struct {
	FlightGRPCConf Server `json:"flight"`
	PlaneGRPCConf  Server `json:"plane"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfig(filepath string) *Config {
	config := &Config{}
	configFile, err := os.ReadFile(filepath)
	if err != nil {
		SetDefaultConfig(config)
	}

	err = json.Unmarshal(configFile, config)
	if err != nil {
		SetDefaultConfig(config)
	}

	return config
}

func SetDefaultConfig(c *Config) {
	c.APIConf.Host = "localhost"
	c.APIConf.Port = "4000"

	c.GRPCConf.FlightGRPCConf.Host = "localhost"
	c.GRPCConf.FlightGRPCConf.Port = "3000"

	c.GRPCConf.PlaneGRPCConf.Host = "localhost"
	c.GRPCConf.PlaneGRPCConf.Port = "3001"
}
