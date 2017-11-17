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
			watchful.Param{Type: watchful.ParamInt, Name: "value", Required: true},
			watchful.Param{Type: watchful.ParamInt, Name: "actual", Required: true},
		},
		Output: []watchful.Param{},
	}
}

func (e Equal) Execute(params map[string]interface{}) (map[string]watchful.InputValue, error) {
	result := map[string]watchful.InputValue{}

	actual, ok := params["actual"].(int64)
	if !ok {
		return result, fmt.Errorf("Checker %s received unknown datatype %T", e.Identifier(), params["actual"])
	}

	value, ok := params["value"].(int64)
	if !ok {
		return result, fmt.Errorf("Checker %s received unknown datatype %T", e.Identifier(), params["value"])
	}

	if actual != value {
		return result, fmt.Errorf("Checker %s expected: %d got: %d", e.Identifier(), value, actual)
	}

	return result, nil
}
