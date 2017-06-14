package config

import (
	"fmt"

	"git.sirius.tivit.com.br/tivit-go-to/worker/backend/src/worker/lib/log"

	"github.com/spf13/viper"
)

const (
	configName string = "main"
)

func init() {
	setDefaultValues()
	readFile()
}

func setDefaultValues() {
	viper.SetDefault(configName+".ServerAddress", "localhost:8080")
}

func readFile() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		log.Error(log.Content{
			Message: fmt.Sprintf("Error in config file: %s (using defaults)", err),
		})
	}
}

// GetString retorna o valor do item especificado na configuração como string
func GetString(name string) string {
	return viper.GetString(configName + "." + name)
}
