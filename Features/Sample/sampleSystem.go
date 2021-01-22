package Sample

import (
	WebCore "WebServer/Core/Foundation"
	"WebServer/Core/Logger"
)

type System struct {
	WebCore.WebSystem
}

func New() WebCore.IWebSystem {
	sys := new(System)
	sys.RegisterHttp("sampleFunc", sys.SampleFunc)
	return sys
}

func (echo *System) SampleFunc(CmdData WebCore.IWebRequest) interface{} {
	Logger.SysLog.Info("[SampleSystem] This is sample func")
	return nil
}




