# viper 使用

> 简单使用


```go
package main

import "github.com/spf13/viper"

type UserConfig struct {
	Name string `mapstructure:"name"`
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
	println(u.Name)
}

```

> 读取env获取对应环境数据