package Settings

type WebConf struct {
	RunMode      string `default:"debug"`
	HttpPort     int    `default:"8084"`
	ReadTimeout  int    `default:"1800"`
	WriteTimeout int    `default:"1800"`
}