package config

import "github.com/spf13/viper"

type Application struct {
	Name    string
	Host    string
	Port    string
	IsHttps bool
	Mode    string
}

var ApplicationConfig = new(Application)

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		Name:    cfg.GetString("name"),
		Host:    cfg.GetString("host"),
		Port:    cfg.GetString("port"),
		IsHttps: cfg.GetBool("isHttps"),
		Mode:    cfg.GetString("mode"),
	}
}
