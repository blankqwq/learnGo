package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

type MysqlConfig struct {
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
}

type UserConfig struct {
	Name  string      `mapstructure:"name"`
	Mysql MysqlConfig `mapstructure:"mysql"`
}

func main() {
	config := viper.New()
	config.AutomaticEnv()
	debug := config.GetString("DEBUG")
	configPrefix := "config"
	if debug == "true" {
		config.SetConfigFile(fmt.Sprintf("viper/ch2/%s-debug.yaml", configPrefix))
	} else {
		config.SetConfigFile(fmt.Sprintf("viper/ch2/%s-production.yaml", configPrefix))
	}

	u := UserConfig{}
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	println(config.GetString("name"))
	if err := config.Unmarshal(&u); err != nil {
		panic(err)
	}
	println(u.Name, u.Mysql.Port, u.Mysql.Name, u.Mysql.Password)

	//watch
	config.WatchConfig()
	config.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("config change : %s", in.Name)
		_ = config.ReadInConfig()
		_ = config.Unmarshal(&u)
		fmt.Printf("%v\n", u)
	})
	time.Sleep(500 * time.Second)
}
