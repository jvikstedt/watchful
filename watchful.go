package watchful

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Executable interface {
	Identifier() string
	Instruction() Instruction
	Execute(map[string]interface{}) (map[string]InputValue, error)
}

type ParamType int

const (
	ParamInt ParamType = iota
	ParamString
	ParamFloat
	ParamBool
	ParamArray
	ParamDynamic = 900
)

type Param struct {
	Type     ParamType `json:"type"`
	Name     string    `json:"name"`
	Required bool      `json:"required"`
}

type Instruction struct {
	Dynamic bool    `json:"dynamic"`
	Input   []Param `json:"input"`
	Output  []Param `json:"output"`
}

type DynamicSource struct {
	TaskID     int    `json:"taskID"`
	OutputName string `json:"outputName"`
}

type InputValue struct {
	Type ParamType   `json:"type"`
	Val  interface{} `json:"val"`
}

type tempValue struct {
	Type      ParamType       `json:"type"`
	Val       json.RawMessage `json:"val"`
	ActualVal interface{}     `json:"-"`
}

func h(tv *tempValue) error {
	var err error
	d := json.NewDecoder(bytes.NewReader(tv.Val))
	d.UseNumber()

	switch tv.Type {
	case ParamInt:
		var val json.Number
		if err := d.Decode(&val); err != nil {
			return err
		}
		if tv.ActualVal, err = val.Int64(); err != nil {
			return err
		}
	case ParamFloat:
		var val json.Number
		if err := d.Decode(&val); err != nil {
			return err
		}
		if tv.ActualVal, err = val.Float64(); err != nil {
			return err
		}
	case ParamString:
		var val string
		if err := d.Decode(&val); err != nil {
			return err
		}
		tv.ActualVal = val
	case ParamBool:
		var val bool
		if err := d.Decode(&val); err != nil {
			return err
		}
		tv.ActualVal = val
	case ParamArray:
		var val []tempValue
		if err := d.Decode(&val); err != nil {
			return err
		}

		var values []InputValue
		for _, v := range val {
			h(&v)
			values = append(values, InputValue{Type: v.Type, Val: v.ActualVal})
		}
		tv.ActualVal = values
	case ParamDynamic:
		var val DynamicSource
		if err := d.Decode(&val); err != nil {
			return err
		}
		tv.ActualVal = val
	}

	return nil
}

func (iv *InputValue) UnmarshalJSON(b []byte) error {
	temp := tempValue{}
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}
	if err := h(&temp); err != nil {
		return err
	}

	iv.Type = temp.Type
	iv.Val = temp.ActualVal

	return nil
}

func (iv InputValue) Value() (driver.Value, error) {
	return json.Marshal(&iv)
}

func (iv *InputValue) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if bv, err := driver.String.ConvertValue(value); err == nil {
		if b, ok := bv.([]byte); ok {
			return json.Unmarshal(b, &iv)
		}
	}
	return fmt.Errorf("Failed to scan InputValue")
}
