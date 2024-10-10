package main

import (
	"main/internal"

	"github.com/spf13/viper"
	"github.com/yanun0323/pkg/config"
	"github.com/yanun0323/pkg/logs"
)

func main() {
	if err := config.Init("config", true, "./"); err != nil {
		logs.Fatalf("init config, err: %v", err)
	}

	conf := internal.Config{}
	if err := viper.Unmarshal(&conf); err != nil {
		logs.Fatalf("unmarshal config, err: %v", err)
	}

	if err := internal.Start(conf); err != nil {
		logs.Fatalf("start app, err: %v", err)
	}

	logs.Infof("app shutdown")
}
