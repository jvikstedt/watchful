package executor

import (
	"fmt"
	"strconv"

	"github.com/jvikstedt/watchful/manager"
)

type Equal struct{}

func (e Equal) Identifier() string {
	return "equal"
}

func (e Equal) Instruction() manager.Instruction {
	return manager.Instruction{
		Takes: []manager.Param{
			manager.Param{Type: manager.ParamString, Name: "value", Required: true},
			manager.Param{Type: manager.ParamAny, Name: "actual", Required: true},
		},
		Returns: []manager.Param{},
	}
}

func (e Equal) Execute(params map[string]interface{}) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	switch v := params["actual"].(type) {
	case int:
		t, err := strconv.Atoi(params["value"].(string))
		if err != nil {
			return result, err
		}
		if t != v {
			return result, fmt.Errorf("Checker %s expected: %d got: %d", e.Identifier(), v, t)
		}
	default:
		return result, fmt.Errorf("Checker %s received unknown datatype %T", e.Identifier(), v)
	}
	return result, nil
}
