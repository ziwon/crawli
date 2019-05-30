package config

import (
	"fmt"
	"os"
	"path"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
	Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error
}

var defaultConfig *viper.Viper

func Config() Provider {
	return defaultConfig
}

func LoadConfigProvider(appName string) Provider {
	return readViperConfig(appName)
}

func init() {
	defaultConfig = readViperConfig("CRAWLI")
}

func readViperConfig(appName string) *viper.Viper {
	v := viper.New()
	home, err := homedir.Dir()
	if err != nil {
		os.Exit(1)
	}

	cfgFile := path.Join(home, ".crawli", "config", "config.toml")
	if _, err := os.Stat(cfgFile); err == nil {
		v.SetConfigFile(cfgFile)
	}

	v.SetEnvPrefix(appName)
	v.AutomaticEnv()

	// global defaults
	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	return v
}
