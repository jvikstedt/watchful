package builtin

import (
	"fmt"

	"github.com/jvikstedt/watchful"
)

type MoreThan struct{}

func (e MoreThan) Identifier() string {
	return "more_than"
}

func (e MoreThan) Instruction() watchful.Instruction {
	return watchful.Instruction{
		Input: []watchful.Param{
			watchful.Param{Type: watchful.ParamInt, Name: "value", Required: true},
			watchful.Param{Type: watchful.ParamInt, Name: "actual", Required: true},
		},
		Output: []watchful.Param{},
	}
}

func (e MoreThan) Execute(params map[string]interface{}) (map[string]watchful.InputValue, error) {
	result := map[string]watchful.InputValue{}

	actual, ok := params["actual"].(float64)
	if !ok {
		return result, fmt.Errorf("Checker %s received unknown datatype %T", e.Identifier(), params["actual"])
	}

	value, ok := params["value"].(float64)
	if !ok {
		return result, fmt.Errorf("Checker %s received unknown datatype %T", e.Identifier(), params["value"])
	}

	if actual < value {
		return result, fmt.Errorf("Checker %s expected: %f to be more than %f", e.Identifier(), actual, value)
	}

	return result, nil
}
