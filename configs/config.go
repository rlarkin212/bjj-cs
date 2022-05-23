package configs

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Version string `mapstructure:"version"`
	Rest    Rest   `mapstructure:"rest"`
	GQL     GQL    `mapstructure:"gql"`
	Mongo   Mongo  `mapstructure:"mongo"`
}

type Rest struct {
	Port string `mapstructure:"port"`
}

type GQL struct {
	Port string `mapstructure:"port"`
}

type Mongo struct {
	URL        string `mapstructure:"url"`
	DB         string `mapstructure:"db"`
	Collection string `mapstructure:"collection"`
	Timeout    int64  `mapstructure:"timeout"`
}

func LoadConfig(path, name, extension string) (*Config, error) {
	c := &Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(extension)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	replaceEnvVars("${", "}")

	err = viper.Unmarshal(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func replaceEnvVars(prefix, suffix string) {
	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)
		if strings.HasPrefix(v, prefix) && strings.HasSuffix(v, suffix) {
			viper.Set(k, os.ExpandEnv(v))
		}
	}
}
