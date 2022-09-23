package config

import "github.com/spf13/viper"

func setDefaultConfig() {
	viper.SetDefault("Data.LogConfig.EnableConsole", true)
	viper.SetDefault("Data.LogConfig.ConsoleJSONFormat", false)
	viper.SetDefault("Data.LogConfig.ConsoleLevel", "debug")
	viper.SetDefault("Data.LogConfig.EnableFile", true)
	viper.SetDefault("Data.LogConfig.FileJSONFormat", true)
	viper.SetDefault("Data.LogConfig.FileLevel", "debug")
	viper.SetDefault("Data.LogConfig.FileLocation", "./photon.log")
	viper.SetDefault("Data.Env", "dev")
	viper.SetDefault("Data.Port", "9876")
	viper.SetDefault("Data.Verbose", true)
}
