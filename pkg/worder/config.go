package worder

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type RemoteDictionary struct {
	Name string
	Url  string
	Path string
	Dest string
}

type Config struct {
	AppName            string             `mapstructure:"app_name"`
	AppVersion         string             `mapstructure:"version"`
	DebugMode          bool               `mapstructure:"debug_mode"`
	RunDir             string             `mapstructure:"run_dir"`
	RemoteDictionaries []RemoteDictionary `mapstructure:"dictionaries"`
}

func Initalize() Config {

	// initial application config
	var config Config
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/worder/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.worder") // call multiple times to add many search paths
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file:", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("failed to unmarshal configuration: %s", err)
	}
	log.Println("configuration successfully loaded -> ", config)

	// Setup local working directory
	if _, err := os.Stat(config.RunDir); os.IsNotExist(err) {
		err := os.Mkdir(config.RunDir, os.ModePerm)
		if err != nil {
			log.Fatal("Error creating: ", config.RunDir, err)
		}
	}

	return config
}
