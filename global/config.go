package global

// Config 配置结构体
type Config struct {
	App   App   `mapstructure:"app" yaml:"app"`
	Log   Log   `mapstructure:"log" yaml:"log"`
	Redis Redis `mapstructure:"redis" yaml:"redis"`
	Jwt   Jwt   `mapstructure:"jwt" yaml:"jwt"`
	Mongo Mongo `mapstructure:"mongo" yaml:"mongo"`
	Mysql Mysql `mapstructure:"mysql" yaml:"mysql"`
}

// App 配置
type App struct {
	Port int    `mapstructure:"port" yaml:"port"`
	ENV  string `mapstructure:"port" yaml:"port"`
	Name string `mapstructure:"name" yaml:"name"`
}

// Social Server 配置
type Social struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port int    `mapstructure:"port" yaml:"port"`
}

type Log struct {
	Path  string `mapstructure:"path" yaml:"path"`
	Level string `mapstructure:"level" yaml:"level"`
}

// Redis 配置
type Redis struct {
	Addr        string `mapstructure:"addr" yaml:"addr"`
	Password    string `mapstructure:"password" yaml:"password"`
	DB          int    `mapstructure:"db" yaml:"db"`
	MaxRetries  int    `mapstructure:"maxRetries" yaml:"maxRetries"`
	PoolSize    int    `mapstructure:"poolSize" yaml:"poolSize"`
	MinIdleConn int    `mapstructure:"minIdleConn" yaml:"minIdleConn"`
}

type Jwt struct {
	SigningKey string `mapstructure:"signingKey" yaml:"signingKey"`
	WhiteList  string `mapstructure:"whiteList" yaml:"whiteList"`
}

// MongoDB 配置
type Mongo struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     string `mapstructure:"port" yaml:"port"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Db       string `mapstructure:"db" yaml:"db"`
}

// MongoDB 配置
type Mysql struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     string `mapstructure:"port" yaml:"port"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Db       string `mapstructure:"db" yaml:"db"`
}
