package main

import (
	"fmt"

	"github.com/jvikstedt/watchful"
)

var Extension = HelloWorld{}

type HelloWorld struct{}

func (h HelloWorld) Identifier() string {
	return "hello_world"
}

func (h HelloWorld) Instruction() watchful.Instruction {
	return watchful.Instruction{
		Input: []watchful.Param{
			watchful.Param{Type: watchful.ParamString, Name: "name", Required: true},
		},
		Output: []watchful.Param{
			watchful.Param{Type: watchful.ParamString, Name: "response"},
		},
	}
}

func (h HelloWorld) Execute(params map[string]interface{}) (map[string]watchful.InputValue, error) {
	result := map[string]watchful.InputValue{}

	name, ok := params["name"]
	if !ok {
		return result, fmt.Errorf("Can't greet you if I don't know your name :/")
	}

	asStr, ok := name.(string)
	if !ok {
		return result, fmt.Errorf("Are you crazy? Your name can't be type %T", name)
	}

	result["response"] = watchful.InputValue{Type: watchful.ParamString, Val: fmt.Sprintf("Hello %s\n", asStr)}
	return result, nil
}
