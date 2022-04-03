package model

type App struct {
	Database        Database `yaml:"database"`
	Redis           Redis    `yaml:"redis"`
	FlushAllForTest bool     `yaml:"flushAllForTest"`
}

type Database struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
	Address  string `yaml:"address"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

type Redis struct {
	Address     string `yaml:"address"`
	Network     string `yaml:"network"`
	Password    string `yaml:"password"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxActive   int    `yaml:"maxActive"`
	IdleTimeout int    `yaml:"idleTimeout"`
}

type AppConfig struct {
	App App `yaml:"app"`
}
