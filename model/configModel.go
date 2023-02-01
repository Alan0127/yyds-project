package model

type App struct {
	Database        Database      `yaml:"database"`
	Redis           Redis         `yaml:"redis"`
	FlushAllForTest bool          `yaml:"flushAllForTest"`
	Kafka           Kafka         `yaml:"kafka"`
	Oss             Oss           `yaml:"oss"`
	ElasticSearch   ElasticSearch `yaml:"elasticSearch"`
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

type Kafka struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
	TimeOut int    `yaml:"timeOut"`
}

type Oss struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
	Path            string `yaml:"path"`
}

type ElasticSearch struct {
	Url      string `yaml:"url"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type AppConfig struct {
	App App `yaml:"app"`
}
