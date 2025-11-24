package config

import (
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	cfg  *TomlConfig
	once sync.Once
)

// Load env: dev/prod/test or ""
func Load(env string) (*TomlConfig, error) {
	var err error

	once.Do(func() {
		cfg, err = loadInternal(env)
	})

	return cfg, err
}

func loadInternal(env string) (*TomlConfig, error) {

	v := viper.New()
	v.SetConfigType("toml")
	v.AddConfigPath("./configs")

	// choose file
	if env != "" {
		v.SetConfigName(fmt.Sprintf("config.%s", env))
	} else {
		v.SetConfigName("config")
	}

	// environment variables
	v.SetEnvPrefix("APP") // APP_DB_HOST
	v.AutomaticEnv()

	// default config
	c := DefaultConfig()

	// try load file
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("[Config] missing config file, using defaults.")
	} else {
		if err := v.Unmarshal(&c); err != nil {
			return nil, fmt.Errorf("unmarshal error: %w", err)
		}
		fmt.Println("[Config] loaded:", v.ConfigFileUsed())
	}

	// optional: watch file changes
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("[Config] file changed:", e.Name)
	})

	return c, nil
}

func GetConfig() *TomlConfig {
	if cfg == nil {
		panic("config not loaded: call Load() first")
	}
	return cfg
}
