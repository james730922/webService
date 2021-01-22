package Foundation

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"time"
)

type Cmd struct {
	CmdId     string      `json:"cmd_id"`
	CmdName   string      `json:"cmd_name"`
	CmdData   interface{} `json:"cmd_data"`
}

type SimpleResp struct {
	Count     int         `json:"count"`          //request count
	Data      interface{} `json:"data,omitempty"`
}

var limitRate = 60

func RouteApis(c iris.Context) {
	//Requests times checking
	cmdTimeKey := fmt.Sprintf("cmdtimes_%s_%d", c.RemoteAddr(), time.Now().Minute())
	times, err := GetServer().GetCacheDb().GetInt(cmdTimeKey)
	if err == nil && times >= limitRate{
		_, _ = c.Text("Error")
		return
	}
	var timeEXPIRE int64 = 60
	setErr := GetServer().GetCacheDb().SetIncr(cmdTimeKey, timeEXPIRE)
	if setErr == nil {
		times++
	}
	// Command Api
	cmdId := c.Params().Get("cmdId")
	cmdName := c.Params().Get("cmdName")
	cmdData, _ := c.GetBody()
	command := &Cmd{
		CmdId:       cmdId,
		CmdName:     cmdName,
		CmdData:     cmdData,
	}

	system := GetServer().GetSystem(command.CmdId)
	if system == nil {
		_, _ = c.Text("System not found")
		return
	}
	_ = system.RunHttpCommand(command)
	_, _ = c.Text("%d", times)
}



