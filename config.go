package main

import (
	"github.com/spf13/viper"
)

var (
	Config *viper.Viper
)

func LoadConf() {
	Config = viper.New()
	Config.SetConfigName("configuration")
	Config.AddConfigPath(".")
	err := Config.ReadInConfig()
	if err != nil {
		Logger.Fatalf("[CONFIG] Error reading configuration : %s", err)
	}
	Config.SetDefault("metaFile", "metadata.yaml")
	Config.SetDefault("buildFolder", "build")
	Config.SetDefault("textFolder", "content")
	Config.SetDefault("pictureFolder", "img")
	Config.SetDefault("templateFolder", "templates")
	Config.SetDefault("translateFolder", "translations")
	Config.SetDefault("latexEngine", "lualatex")
}
