package config

type TomlConfig struct {
	AppName    string           `mapstructure:"app_name"`
	MySQL      MySQLConfig      `mapstructure:"mysql"`
	Log        LogConfig        `mapstructure:"log"`
	StaticPath PathConfig       `mapstructure:"static_path"`
	Blockchain BlockchainConfig `mapstructure:"blockchain"`
}

// MySQL相关配置
type MySQLConfig struct {
	Host        string `mapstructure:"host"`
	Name        string `mapstructure:"name"`
	Password    string `mapstructure:"password"`
	Port        int    `mapstructure:"port"`
	TablePrefix string `mapstructure:"table_prefix"`
	User        string `mapstructure:"user"`
}

// 日志保存地址
type LogConfig struct {
	Path  string `mapstructure:"path"`
	Level string `mapstructure:"level"`
}

// 相关地址信息，例如静态文件地址
type PathConfig struct {
	FilePath string `mapstructure:"file_path"`
}

// 区块链相关配置
type BlockchainConfig struct {
	RPC_URL        string `mapstructure:"rpc_url"`
	ChainID        int64  `mapstructure:"chain_id"`
	PrivateKey     string `mapstructure:"private_key"`
	GasLimit       uint64 `mapstructure:"gas_limit"`
	GasTipCapGwei  int64  `mapstructure:"gas_tip_cap_gwei"`
	GasFeeCapGwei  int64  `mapstructure:"gas_fee_cap_gwei"`
	CounterAddress string `mapstructure:"counter_address"`
}
