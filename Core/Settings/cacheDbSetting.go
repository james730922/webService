package Settings

type CacheDbConf struct {
	Host        string `default:"127.0.0.1"`
	Port        int    `default:"6379"`
	Password    string `default:""`
	MaxIdle     int    `default:"30"`
	MaxActive   int    `default:"20"`
	IdleTimeout int    `default:"200"`
	Wait        bool   `default:"true"`
}
