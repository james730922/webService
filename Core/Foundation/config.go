package Foundation

import (
	"WebServer/Core/Logger"
	"WebServer/Core/Settings"
	"errors"
	"github.com/koding/multiconfig"
)

type IConfig interface{}

type Config struct {
	App     *Settings.AppConf
	Web     *Settings.WebConf
	CacheDB *Settings.CacheDbConf
	envMap  map[string]interface{}
	raw     map[string]string
	engine  *Engine
}

var _ IConfig = &Config{}

func (config *Config) LoadExternalEnv(envPrefix string, conf interface{}) {
	config.loadEnv(envPrefix, conf)
	config.envMap[envPrefix] = conf
}

func (config *Config) GetEnv(prefix string) (interface{}, error) {
	if val, ok := config.envMap[prefix]; ok {
		return val, nil
	}
	Logger.SysLog.Errorf("[ConfigSystem] Config Not Found in Prefix `%s`, Please Check", prefix)
	return nil, errors.New("settings not found")
}

func (config *Config) systemExternalEnv(envPrefix string, conf interface{}) {
	config.loadEnv(envPrefix, conf)
}

func (config *Config) loadEnv(envPrefix string, conf interface{}) {
	InstantiateLoader := &multiconfig.EnvironmentLoader{
		Prefix:    envPrefix,
		CamelCase: true,
	}
	err := InstantiateLoader.Load(conf)
	if err != nil {
		Logger.SysLog.Errorf("[Config] Load Env Failed, %s", err)
		return
	}
}
