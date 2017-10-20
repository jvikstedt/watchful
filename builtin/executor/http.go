package executor

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jvikstedt/watchful/manager"
)

type HTTP struct{}

func (h HTTP) Identifier() string {
	return "http"
}

func (h HTTP) Instruction() manager.Instruction {
	return manager.Instruction{
		Takes: []manager.Param{
			manager.Param{Type: manager.ParamString, Name: "url", Required: true},
		},
		Returns: []manager.Param{
			manager.Param{Type: manager.ParamInt, Name: "statusCode"},
			manager.Param{Type: manager.ParamBytes, Name: "body"},
		},
	}
}

func (h HTTP) Execute(params map[string]interface{}) (map[string]interface{}, error) {
	url, ok := params["url"].(string)
	if !ok {
		return nil, fmt.Errorf("Expected url to be a string but was %T", params["url"])
	}

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"statusCode": response.StatusCode,
		"body":       bodyBytes,
	}, nil
}
