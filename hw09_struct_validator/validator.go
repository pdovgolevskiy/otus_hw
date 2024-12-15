package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

var (
	ErrUnsupportedOperation = errors.New("unsupported validation")
	ErrInvalidRule          = errors.New("invalid rule")
)

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	panic("implement me")
}

func findRulesFromTag(validateFullStr string) {
	rules := strings.Split(validateFullStr, "|")
	for _, rule := range rules {
		rl := strings.Split(rule, ":")

	}
}

func Validate(v interface{}) error {
	st := reflect.TypeOf(v)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if validateVal, ok := field.Tag.Lookup("validate"); ok {
			if validateVal == "" {
				continue
				//TODO
			} else {
				findRulesFromTag(validateVal)
			}
		} else {
			fmt.Println("(not specified)")
		}
	}
	// Place your code here.
	return nil
}
