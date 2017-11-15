package model

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/jvikstedt/watchful"
)

type InputValue struct {
	Type watchful.ParamType `json:"type"`
	Val  interface{}        `json:"val"`
}

type tempValue struct {
	Type      watchful.ParamType `json:"type"`
	Val       json.RawMessage    `json:"val"`
	ActualVal interface{}        `json:"-"`
}

func h(tv *tempValue) error {
	var err error
	d := json.NewDecoder(bytes.NewReader(tv.Val))
	d.UseNumber()

	switch tv.Type {
	case watchful.ParamInt:
		var val json.Number
		if err := d.Decode(&val); err != nil {
			return err
		}
		if tv.ActualVal, err = val.Int64(); err != nil {
			return err
		}
	case watchful.ParamFloat:
		var val json.Number
		if err := d.Decode(&val); err != nil {
			return err
		}
		if tv.ActualVal, err = val.Float64(); err != nil {
			return err
		}
	case watchful.ParamString:
		var val string
		if err := d.Decode(&val); err != nil {
			return err
		}
		tv.ActualVal = val
	case watchful.ParamBool:
		var val bool
		if err := d.Decode(&val); err != nil {
			return err
		}
		tv.ActualVal = val
	case watchful.ParamArray:
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
