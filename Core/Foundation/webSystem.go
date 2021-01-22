package Foundation

import (
	"WebServer/Core/Logger"
	"strings"
)

///* Http System*///

type IWebSystem interface {
	RunHttpCommand(*Cmd) interface{}
}

type WebSystem struct {
	IWebSystem,
	httpCmdMap map[string]func(IWebRequest) interface{}
}

func (sys *WebSystem) RegisterHttp(operator string, f func(IWebRequest) interface{}) {
	if sys.httpCmdMap == nil {
		sys.httpCmdMap = make(map[string]func(IWebRequest) interface{})
	}
	cmdName := strings.ToLower(operator)
	sys.httpCmdMap[cmdName] = f
	Logger.SysLog.Debugf("[Http] `%s` Registered", cmdName)
}

func (sys *WebSystem) RunHttpCommand(data *Cmd) interface{} {
	RequestData := &WebRequest{CmdData: data.CmdData}
	cmdId := strings.ToLower(data.CmdId)
	cmdName := strings.ToLower(data.CmdName)
	if httpFunc, httpFuncExist := sys.httpCmdMap[cmdName]; httpFuncExist {
		Logger.SysLog.Infof("[HttpCommand] cmdId : %s, cmdName : %s", cmdId, cmdName)
		return httpFunc(RequestData)
	}
	return nil
}

///* Request*///

type IWebRequest interface {
	Raw() interface{}
}

type WebRequest struct {
	CmdData  interface{}
}

func (req *WebRequest) Raw() interface{} {
	return req.CmdData
}



