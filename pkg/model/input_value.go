package model

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type InputValue struct {
	Val interface{}
}

func (iv *InputValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(&iv.Val)
}
func (iv *InputValue) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	return d.Decode(&iv.Val)
}

func (iv *InputValue) Value() (driver.Value, error) {
	if iv == nil {
		return nil, nil
	}
	return json.Marshal(&iv.Val)
}
func (iv *InputValue) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if bv, err := driver.String.ConvertValue(value); err == nil {
		if b, ok := bv.([]byte); ok {
			d := json.NewDecoder(bytes.NewReader(b))
			d.UseNumber()
			return d.Decode(&iv.Val)
		}
	}
	return fmt.Errorf("Failed to scan InputValue")
}
