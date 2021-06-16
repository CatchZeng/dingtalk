package configs

import (
	"log"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// InitConfig reads in configs file and ENV variables if set.
func InitConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Panic(err)
	}

	// Search configs in home directory with name ".dingtalk" (without extension).
	configPath := path.Join(home, ".dingtalk")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		return
	}
	log.Println("using configs file:", viper.ConfigFileUsed())
}

// GetConfig get configs with key
func GetConfig(key string) (string, error) {
	// If a configs file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		return viper.GetString(key), nil
	}
	return "", err
}
