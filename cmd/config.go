package main

import (
	"github.com/spf13/viper"
)

func setupConfig() {
	viper.SetEnvPrefix("kviewer")

	viper.AutomaticEnv()

	viper.SetDefault("bootstrap.servers", "localhost:29092")
}
