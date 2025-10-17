package config

type API struct {
	Server APIServerConfig `yaml:"server"`
	JWT    JWTConfig       `yaml:"jwt"`
}

type APIServerConfig struct {
	ListenAddress string `yaml:"listen_address"`
	ListenPort    string `yaml:"listen_port"`
}

func (a *APIServerConfig) ListenString() string {
	return a.ListenAddress + ":" + a.ListenPort
}

type JWTConfig struct {
	SecretKey     string `yaml:"secret_key"`
	TokenDuration int    `yaml:"expiration_minutes"` // in minutes
}
