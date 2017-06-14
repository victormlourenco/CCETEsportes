package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
	configName string = "log"
)

func init() {
	setDefaultValues()
	readFile()
}

func setDefaultValues() {
	viper.SetDefault(configName+".FilePath", "./")
	viper.SetDefault(configName+".FileName", "main.log")
}

func readFile() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		fmt.Printf("Error in config file: %s (using default values)\n", err)
	}
}

// Get retorna o valor do item especificado na configuração como interface{}
func Get(name string) interface{} {
	return viper.Get(configName + "." + name)
}

// GetString retorna o valor do item especificado na configuração como string
func GetString(name string) string {
	return viper.GetString(configName + "." + name)
}

// GetBool retorna o valor do item especificado na configuração como bool
func GetBool(name string) bool {
	return viper.GetBool(configName + "." + name)
}

// GetFloat64 retorna o valor do item especificado na configuração como float64
func GetFloat64(name string) float64 {
	return viper.GetFloat64(configName + "." + name)
}

// GetInt retorna o valor do item especificado na configuração como int
func GetInt(name string) int {
	return viper.GetInt(configName + "." + name)
}

// GetStringMap retorna o valor do item especificado na configuração como map[string]interface{}
func GetStringMap(name string) map[string]interface{} {
	return viper.GetStringMap(configName + "." + name)
}

// GetStringMapString retorna o valor do item especificado na configuração como map[string]string
func GetStringMapString(name string) map[string]string {
	return viper.GetStringMapString(configName + "." + name)
}

// GetStringSlice retorna o valor do item especificado na configuração como []string
func GetStringSlice(name string) []string {
	return viper.GetStringSlice(configName + "." + name)
}

// GetTime retorna o valor do item especificado na configuração como time.Time
func GetTime(name string) time.Time {
	return viper.GetTime(configName + "." + name)
}

// GetDuration retorna o valor do item especificado na configuração como time.Duration
func GetDuration(name string) time.Duration {
	return viper.GetDuration(configName + "." + name)
}
