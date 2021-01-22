package Settings

type AppConf struct {
	Codename       string `default:""`
	LogStage       string `default:"dev"`
	LogLevel       string `default:"info"`
}
