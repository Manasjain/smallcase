package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"smallcase/config/database"
)

const (
	envVar = "AppConf"
)

var (
	// errNilConfig is returned when a nil reference is passed in as Un/Marshaler reference
	errNilConfig = errors.New("Config object is empty.")
)
var Config = &struct {
	HostAndPort string       `json:"hostAndPort"`
	Mysql       *MysqlConfig `json:"mysql"`
}{}

type MysqlConfig struct {
	Dsn      string `json:"dsn"`
	HostName string `json:"hostname"`
	Password string `json:"password"`
	Username string `json:"username"`
	DBName   string `json:"dbName"`
}

// LoadJSONFile gets your config from the json file,
// and fills your struct with the option
func LoadJSONFile(filename string, config interface{}) error {
	if config == nil {
		return errNilConfig
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, config)
}

//LoadJSONEnvPath Gets config file path from os.Getenv & calls LoadJSONFile
// but returns error if no path is set
func LoadJSONEnvPath(envVar string, config interface{}) error {
	if config == nil {
		return errNilConfig
	}

	filename := os.Getenv(envVar)
	if filename == "" {
		return fmt.Errorf("Env var is empty: %s", envVar)
	}
	log.Printf(" : loading config from envVar %s, file = %s", envVar, filename)
	return LoadJSONFile(filename, config)
}

// LoadJSONEnvPathOrPanic calls LoadJSONEnvPath but panics on error
func LoadJSONEnvPathOrPanic(envVar string, config interface{}) {
	if err := LoadJSONEnvPath(envVar, config); err != nil {
		panic(fmt.Errorf("failed to load config file with error %s", err))
	}
}

func init() {
	LoadJSONEnvPathOrPanic(envVar, Config)
	Config.Mysql.Dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.Mysql.Username, Config.Mysql.Password, Config.Mysql.HostName, Config.Mysql.DBName)
	database.Initialize(Config.Mysql.Dsn)
}
