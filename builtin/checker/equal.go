package checker

import (
	"fmt"
	"strconv"
)

type Equal struct{}

func (e Equal) Identifier() string {
	return "equal"
}

func (e Equal) Check(target string, i interface{}) error {
	switch v := i.(type) {
	case int:
		t, err := strconv.Atoi(target)
		if err != nil {
			return err
		}
		if t != v {
			return fmt.Errorf("Checker %s expected: %d got: %d", e.Identifier(), v, t)
		}
	default:
		return fmt.Errorf("Checker %s received unknown datatype %T", e.Identifier(), v)
	}
	return nil
}
