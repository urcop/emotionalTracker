package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"os"
)

type HttpServer struct {
	Host string `yaml:"host" env-default:"0.0.0.0" env-required:"true"`
	Port string `yaml:"port" env-default:"8000" env-required:"true"`
}

type Db struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	Name     string `yaml:"db_name" env-default:"postgres"`
	SslMode  string `yaml:"sslmode"`
}

type GrpcServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Env        string     `yaml:"env" env-required:"true"`
	HttpServer HttpServer `yaml:"http_server"`
	Db         Db         `yaml:"db"`
	GrpcServer GrpcServer `yaml:"grpc_server" env-required:"true"`
}

func Make() *Config {

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist " + path)
	}

	var config Config

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		slog.Error("cannot read config file", err)
	}

	return &config
}

func (s *Config) HttpHost() string {
	return s.HttpServer.Host
}

func (s *Config) HttpPort() string {
	return s.HttpServer.Port
}

func (s *Config) PostgresHost() string {
	return s.Db.Host
}

func (s *Config) PostgresPort() string {
	return s.Db.Port
}

func (s *Config) PostgresUser() string {
	return s.Db.User
}

func (s *Config) PostgresPassword() string {
	return s.Db.Password
}

func (s *Config) PostgresName() string {
	return s.Db.Name
}

func (s *Config) EnvLevel() string {
	return s.Env
}

func (s *Config) GrpcHost() string {
	return s.GrpcServer.Host
}

func (s *Config) GrpcPort() string {
	return s.GrpcServer.Port
}

func (s *Config) SslMode() string {
	return s.Db.SslMode
}
