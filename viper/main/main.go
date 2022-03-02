package main

import "github.com/spf13/viper"

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
	config.SetConfigFile("viper/main/config-production.yaml")
	u := UserConfig{}
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	println(config.GetString("name"))
	if err := config.Unmarshal(&u); err != nil {
		panic(err)
	}
	println(u.Name, u.Mysql.Port, u.Mysql.Name, u.Mysql.Password)
}
