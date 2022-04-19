package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gorpc_study/nacos_test/main/config"
)

func main() {
	sc := []constant.ServerConfig{
		{
			IpAddr: "localhost",
			Port:   8848,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         "92eb12b0-1a27-41ec-a6f8-4a8bb1611e56", //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev",
	})
	fmt.Println("GetConfig,config :" + content)

	c := config.ServerConfig{}
	// json
	err = json.Unmarshal([]byte(content), &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
