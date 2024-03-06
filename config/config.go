package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sd-client/httpclient"
	"sd-client/logger"
)

// Conf 配置文件
var Conf *Config

// DefaultPath 默认配置路径
var DefaultPath = "../../config.yml"

var allConfig *Config

type Job struct {
	IpMonitor int64 `yaml:"ipMonitor"`
	IpPing    int64 `yaml:"ipPing"`
	SyncCloud int64 `yaml:"syncCloud"`
}

type Bind struct {
	Enable          bool   `yaml:"enable"`
	ZonesDir        string `yaml:"zonesDir"`
	ZoneFilePrefix  string `yaml:"zoneFilePrefix"`
	ZoneFileSuffix  string `yaml:"zoneFileSuffix"`
	ReloadConfigCmd string `yaml:"reloadConfigCmd"`
}

type Sd struct {
	CloudHost string   `yaml:"cloudHost"`
	Job       Job      `yaml:"job"`
	Bind      Bind     `yaml:"bind"`
	CallHosts []string `yaml:"callHosts"`
}

// Config 配置文件
type Config struct {
	Port string                `yaml:"port"`
	Log  logger.LogConfig      `yaml:"log"`
	Sd   Sd                    `yaml:"sd"`
	Http httpclient.HttpConfig `yaml:"http"`
}

// GetAllConfig Get all configurations
func GetAllConfig() (*Config, error) {
	if allConfig == nil {
		return nil, errors.New("config is nil")
	}
	return allConfig, nil
}

// NewConfig
func NewConfig(path string) (*Config, error) {
	if path == "" {
		path = DefaultPath
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &Conf)
	if err != nil {
		return nil, err
	}
	allConfig = Conf
	return Conf, nil
}
