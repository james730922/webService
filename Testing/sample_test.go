package Testing

import (
	"WebServer/Core/Logger"
	"testing"
)

func TestSample(t *testing.T) {
	data := map[string]interface{}{}
	for i:=0; i <= 60; i++ {
		response, err := SendHttpCommand("sample", "sampleFunc", data)
		if err != nil {
			Logger.SysLog.Errorf("[Test]sampleFunc cmd error :", err)
		} else {
			Logger.SysLog.Info("[Test]sampleFunc resp :", response)
		}
	}

}
