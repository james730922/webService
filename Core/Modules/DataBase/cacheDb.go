package DataBase

import (
	"WebServer/Core/Logger"
	"WebServer/Core/Settings"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

/*var (
	CacheDbClient *CacheDB
)*/

type CacheDB struct {
	IDB
	clientPool *redis.Pool
}

func ConnectWithCacheDB(config *Settings.CacheDbConf) (*CacheDB, error) {
	baseConnectString := fmt.Sprintf("%s:%d", config.Host, config.Port)
	Logger.SysLog.Infof("[CacheDb] Connecting to Cache Service -> %s", baseConnectString)

	client := &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: time.Duration(config.IdleTimeout) * time.Millisecond,
		Wait:        config.Wait,
		Dial: func() (redis.Conn, error) {
			Logger.SysLog.Debug("[CacheDb] Dial Connects To The CacheUtils Server")
			c, err := redis.Dial("tcp", baseConnectString)
			if err != nil {
				return nil, err
			}
			if config.Password != "" {
				if _, err := c.Do("AUTH", config.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	CacheDbClient := &CacheDB{
		clientPool: client,
	}
	// Checking cache can be connected
	err := CacheDbClient.pingHealth()
	if err != nil {
		Logger.SysLog.Debug("[CacheDB] ping Health Error")
		return nil, err
	}
	return CacheDbClient, nil
}

func (r *CacheDB) GetClient() redis.Conn {
	return r.clientPool.Get()
}

func (r *CacheDB) pingHealth() error {
	conn := r.GetClient()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		return err
	}
	return nil
}

func (r *CacheDB) GetInt(key string) (int, error) {
	conn := r.clientPool.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("GET", key))
	if err != nil {
		return 0, err
	}

	return reply, nil
}

func (r *CacheDB) SetIncr(key string, time int64) error {
	conn := r.clientPool.Get()
	defer conn.Close()
	_, err := conn.Do("INCR", key)
	if err != nil {
		return err
	}
	if time > 0 {
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}
	}
	return nil
}


