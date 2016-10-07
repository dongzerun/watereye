package conf

import (
	"io/ioutil"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	GlobalConfig Config
)

type Config struct {
	LogConfig  LogConfig  `toml:"golog"` // 日志配置
	HttpConfig HttpConfig `toml:"http"`
	// AdDBConfig       DBconfig               `toml:"addb"`         // 广告数据库配置
	// CacheConfig      map[string]Redisconfig `toml:"cache"`     // redis cache配置
}

type HttpConfig struct {
	Listen string `toml:"listen"`
}

type LogConfig struct {
	Level      string `toml:"level"`
	Console    int    `toml:"console"`
	Dir        string `toml:"dir"`
	Filename   string `toml:"filename"`
	ReserveNum int    `toml:"reserve_num"`
	Suffix     string `toml:"suffix"`
	Colorfull  int    `toml:"colorfull"`
}

type DBconfig struct {
	Usr      string `toml:"user"`
	Pwd      string `toml:"pwd"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DBname   string `toml:"db_name"`
	MaxIdle  int    `toml:"max_idle"`
	MaxOpen  int    `toml:"max_open"`
	PoolSize int    `toml:"pool_size"`
}

// idleTimeout, connectTimeout, readTimeout, writeTimeout time.Duration, poolSize
type Redisconfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	PoolSize int    `toml:"pool_size"`
	MaxIdle  int    `toml:"max_idle"`

	IdleTimeout    time.Duration `toml:"idle_timeout"`
	ConnectTimeout time.Duration `toml:"connect_timeout"`
	ReadTimeout    time.Duration `toml:"read_timeout"`
	WriteTimeout   time.Duration `toml:"write_timeout"`
}

func InitConfig(filename string) {
	GlobalConfig = NewTomlConfig(filename)
}

// 获取config
func NewTomlConfig(filename string) (conf Config) {

	var (
		data []byte
		err  error
	)

	data, err = ioutil.ReadFile(filename)

	if err != nil {
		panic("read configuration file failed " + err.Error())
	}

	if _, err = toml.Decode(string(data), &conf); err != nil {
		panic("toml decode failed " + err.Error())
	}
	return
}
