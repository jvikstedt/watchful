package builtin

import (
	"fmt"

	"github.com/jvikstedt/watchful"
)

type Equal struct{}

func (e Equal) Identifier() string {
	return "equal"
}

func (e Equal) Instruction() watchful.Instruction {
	return watchful.Instruction{
		Input: []watchful.Param{
			watchful.Param{Type: watchful.ParamAny, Name: "value", Required: true},
			watchful.Param{Type: watchful.ParamAny, Name: "actual", Required: true},
		},
		Output: []watchful.Param{},
	}
}

func (e Equal) Execute(params map[string]interface{}) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	switch v := params["actual"].(type) {
	case float64:
		t, ok := params["value"].(float64)
		if !ok {
			return result, fmt.Errorf("Checker %s expected type float but got: %T", e.Identifier(), params["value"])
		}
		if t != v {
			return result, fmt.Errorf("Checker %s expected: %f got: %f", e.Identifier(), v, t)
		}
	case string:
		t, ok := params["value"].(string)
		if !ok {
			return result, fmt.Errorf("Checker %s was expecting type string but got %T", e.Identifier(), params["value"])
		}
		if t != v {
			return result, fmt.Errorf("Checker %s expected: %s got: %s", e.Identifier(), v, t)
		}
	default:
		return result, fmt.Errorf("Checker %s received unknown datatype %T", e.Identifier(), v)
	}
	return result, nil
}
