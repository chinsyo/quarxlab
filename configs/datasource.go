package configs

import (
	"github.com/BurntSushi/toml"
	"log"
	"fmt"
)

type DataSource interface {
	func Init(path string)
	func GetDSN() string
}

type MySQLConfig struct {
	Name string
	Host string
	Port uint16
	Username string
	Password string
}

func (conf *MySQLConfig) Init(path string)  {
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		log.Fatal(err)
		panic(err)
	}
} 

func (conf *MySQLConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Name)
} 

type SQLiteConfig struct {
	Path string
}

func (conf *SQLiteConfig) Init(path string)  {
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func (conf *SQLiteConfig) GetDSN() string {
	return fmt.Sprintf("%s", conf.Path)
}
