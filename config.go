package firewood

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var Conf *ServerConf

type HttpConf struct {
	Port int `yaml:"port"`
}

type DbConf struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	PoolSize int    `yaml:"pool_size"`
}

type CacheConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	PoolSize int    `yaml:"pool_size"`
	Password string `yaml:"password"`
}

type ServerConf struct {
	AccessLog string `yaml:"access_log"`
	ErrorLog  string `yaml:"error_log"`
	LogLevel  string `yaml:"log_level"`

	Http  *HttpConf  `yaml:"http"`
	Db    *DbConf    `yaml:"db"`
	Cache *CacheConf `yaml:"cache"`
}

func InitServerConf(file string) (*ServerConf, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	c, err := parse(buf)
	if err != nil {
		return nil, err
	}

	Conf = c

	InitLog(c)

	return c, nil
}

func parse(d []byte) (*ServerConf, error) {
	c := &ServerConf{}

	if err := yaml.Unmarshal(d, c); err != nil {
		return nil, err
	}

	return c, nil
}
