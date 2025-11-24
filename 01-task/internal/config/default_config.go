package config

func DefaultConfig() *TomlConfig {
	return &TomlConfig{
		AppName: "task-app",

		MySQL: MySQLConfig{
			Host:        "127.0.0.1",
			Port:        3306,
			User:        "root",
			Password:    "",
			Name:        "app",
			TablePrefix: "",
		},

		Log: LogConfig{
			Path:  "./logs",
			Level: "info",
		},

		StaticPath: PathConfig{
			FilePath: "./static",
		},

		Blockchain: BlockchainConfig{
			RPC_URL:        "",
			ChainID:        1,
			GasLimit:       2000000,
			GasTipCapGwei:  2,
			GasFeeCapGwei:  30,
			CounterAddress: "",
		},
	}
}
