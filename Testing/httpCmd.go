package Testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	TestServer     = "http://:8080/@/api"
)

func SendHttpCommand(id, name string, command interface{}) (interface{}, error) {
	data := map[string]interface{}{
		"cmd_id":       id,
		"cmd_name":     name,
		"cmd_data":     command,
	}
	path := fmt.Sprintf("%s/%s", id, name)
	runCmd := postRequest(path, data)
	cmdResult, _ := json.Marshal(runCmd)
	return string(cmdResult), nil
}

func postRequest(path string, data interface{}) string {
	marshalData, _ := json.Marshal(data)
	postRequest, _ := http.Post(TestServer+"/"+path, "application/json", bytes.NewBuffer(marshalData))
	postRespBody, _ := ioutil.ReadAll(postRequest.Body)
	return string(postRespBody)
}
