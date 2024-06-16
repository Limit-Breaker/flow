package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type Etcd struct {
}

type Redis struct {
}

func MustLoad(path string) {
	viper.SetConfigFile(path)
	cfgb, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(cfgb))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	// 启动参数
	appCfg := viper.Sub("app")
	if appCfg == nil {
		panic("config not found settings.application")
	}
	ApplicationConfig = InitApplication(appCfg)
	dbCfg := viper.Sub("database")
	if dbCfg == nil {
		panic("config not found settings.database")
	}
	DB = InitDB(dbCfg)
}
