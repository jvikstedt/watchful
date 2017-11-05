package builtin

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jvikstedt/watchful"
)

type HTTP struct{}

func (h HTTP) Identifier() string {
	return "http"
}

func (h HTTP) Instruction() watchful.Instruction {
	return watchful.Instruction{
		Input: []watchful.Param{
			watchful.Param{Type: watchful.ParamString, Name: "url", Required: true},
		},
		Output: []watchful.Param{
			watchful.Param{Type: watchful.ParamInt, Name: "statusCode"},
			watchful.Param{Type: watchful.ParamString, Name: "body"},
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
		"body":       string(bodyBytes),
	}, nil
}
