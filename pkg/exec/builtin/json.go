package builtin

import (
	"encoding/json"
	"fmt"

	"github.com/jvikstedt/watchful"
)

type JSON struct{}

func (j JSON) Identifier() string {
	return "json"
}

func (j JSON) Instruction() watchful.Instruction {
	return watchful.Instruction{
		Dynamic: true,
		Input: []watchful.Param{
			watchful.Param{Type: watchful.ParamBytes, Name: "rawjson", Required: true},
		},
		Output: []watchful.Param{},
	}
}

func (j JSON) Execute(params map[string]interface{}) (map[string]interface{}, error) {
	rawjson, ok := params["rawjson"].([]byte)
	if !ok {
		return nil, fmt.Errorf("Expected rawjson to be a []byte but was %T", params["rawjson"])
	}

	original := make(map[string]interface{})

	err := json.Unmarshal(rawjson, &original)
	if err != nil {
		return nil, err
	}

	return j.flatten(original), nil
}

func (j JSON) flatten(m map[string]interface{}) map[string]interface{} {
	o := make(map[string]interface{})
	for k, v := range m {
		switch child := v.(type) {
		case map[string]interface{}:
			nm := j.flatten(child)
			for nk, nv := range nm {
				o[k+"."+nk] = nv
			}
		default:
			o[k] = v
		}
	}
	return o
}
