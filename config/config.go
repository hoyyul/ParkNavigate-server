package config

type Config struct {
	Server *Server `yaml:"server"`
	Mysql  *Mysql  `yaml:"mysql"`
	Redis  *Redis  `yaml:"redis"`
}

type Server struct {
	Address string `yaml:"address"`
	Env     string `yaml:"env"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	PoolSize int    `yaml:"pool_size"`
}
