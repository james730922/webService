package Foundation

import (
	"WebServer/Core/Logger"
	"WebServer/Core/Modules/DataBase"
	"WebServer/Core/Settings"
	"reflect"
	"strings"
	"sync"
)

type Service struct {
	cacheDb         *DataBase.CacheDB
	systemMap       map[string]IWebSystem
}

var (
	service         *Service
	once            sync.Once
)

func GetServer() *Service {
	once.Do(func() {
		service = &Service{
			systemMap: make(map[string]IWebSystem),
		}
	})
	return service
}

func (s *Service) ConnectCacheDbService(config *Settings.CacheDbConf) {
	client, err := DataBase.ConnectWithCacheDB(config)
	if err != nil {
		Logger.SysLog.Errorf("[CacheDb] Try To Connect CacheUtils Database Failed -> (%s)", err)
	}
	s.cacheDb = client
}

func (s *Service) GetCacheDb() *DataBase.CacheDB {
	if s.cacheDb == nil {
		return nil
	}
	return s.cacheDb
}

func (s *Service) Register(name string, inst IWebSystem) bool {
	cmdId := strings.ToLower(name)
	if _, find := s.systemMap[cmdId]; find {
		Logger.SysLog.Warnf(
			"[Engine] Registered System: %s, %s Failed, command Duplicate!",
			cmdId,
			reflect.TypeOf(inst),
			)
		return false
	}
	Logger.SysLog.Infof("[Engine] Registered System: %s, %s", cmdId, reflect.TypeOf(inst))
	s.systemMap[name] = inst
	return true
}

func (s *Service) GetSystem(name string) IWebSystem {
	name = strings.ToLower(name)
	if system, find := s.systemMap[name]; find {
		return system
	}
	return nil
}



























