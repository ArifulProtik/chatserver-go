package config

import "github.com/spf13/viper"

type Config struct {
	AppInfo app
}

type app struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Status  string `yaml:"status"`
}

// New maps all enviornment variable to Config

func New(path string, filename string) (*Config, error) {
	var c Config
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(filename)
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil

}
