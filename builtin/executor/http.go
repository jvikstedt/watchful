package executor

import (
	"fmt"
	"net/http"

	"github.com/jvikstedt/watchful/manager"
)

type HTTP struct{}

func (h HTTP) Name() string {
	return "http"
}

func (h HTTP) Instructions() manager.Instruction {
	return manager.Instruction{
		Takes: []manager.Param{
			manager.Param{Type: manager.ParamString, Name: "url", Required: true},
		},
		Returns: []manager.Param{
			manager.Param{Type: manager.ParamInt, Name: "statusCode"},
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

	return map[string]interface{}{
		"statusCode": response.StatusCode,
	}, nil
}
