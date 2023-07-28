package model

import (
	"fmt"
	"github.com/hashicorp/vault/sdk/framework"
)

type FieldDataWrapper struct {
	*framework.FieldData
}

// NewFieldDataWrapper returns FieldDataWrapper
func NewFieldDataWrapper(data *framework.FieldData) *FieldDataWrapper {
	return &FieldDataWrapper{
		FieldData: data,
	}
}

// GetString get string value from framework.FieldData with specific key,
// returns defaultValue when the key doesn't exist or the value of the key is not type of string.
func (f *FieldDataWrapper) GetString(key string, defaultValue string) string {
	valueInterface := f.Get(key)
	if valueInterface == nil {
		return defaultValue
	}
	value, ok := valueInterface.(string)
	if !ok {
		return defaultValue
	}
	return value
}

// MustGetString get string value from framework.FieldData with specific key,
// returns error when the key doesn't exist or the value of the key is not type of string.
func (f *FieldDataWrapper) MustGetString(key string) (string, error) {
	valueInterface, ok := f.GetOk(key)
	if !ok {
		return "", f.errorResolve(key)
	}
	value, ok := valueInterface.(string)
	if !ok {
		return "", f.errorTypeMismatch(key)
	}
	return value, nil
}

// GetStringSlice get string slice value from framework.FieldData with specific key,
// returns defaultValue when the key doesn't exist or the value of the key is not type of string slice.
func (f *FieldDataWrapper) GetStringSlice(key string, defaultValue []string) []string {
	valueInterface, ok := f.GetOk(key)
	if !ok {
		return defaultValue
	}
	interfaceArray, ok := valueInterface.([]interface{})
	if !ok {
		return defaultValue
	}
	var value []string
	for _, strInterface := range interfaceArray {
		str, ok := strInterface.(string)
		if !ok {
			return defaultValue
		}
		value = append(value, str)
	}
	return value
}

// MustGetStringSlice get string slice value from framework.FieldData with specific key,
// returns error when the key doesn't exist or the value of the key is not type of string slice.
func (f *FieldDataWrapper) MustGetStringSlice(key string) ([]string, error) {
	valueInterface, ok := f.GetOk(key)
	if !ok {
		return nil, f.errorResolve(key)
	}
	interfaceArray, ok := valueInterface.([]interface{})
	if !ok {
		return nil, f.errorTypeMismatch(key)
	}
	var value []string
	for _, strInterface := range interfaceArray {
		str, ok := strInterface.(string)
		if !ok {
			return nil, f.errorTypeMismatch(key)
		}
		value = append(value, str)
	}
	return value, nil
}

// GetBool get bool value from framework.FieldData with specific key,
// returns defaultValue when the key doesn't exist or the value of the key is not type of bool.
func (f *FieldDataWrapper) GetBool(key string, defaultValue bool) bool {
	valueInterface, ok := f.GetOk(key)
	if !ok {
		return defaultValue
	}
	value, ok := valueInterface.(bool)
	if !ok {
		return defaultValue
	}
	return value
}

// MustGetBool get bool value from framework.FieldData with specific key,
// returns error when the key doesn't exist or the value of the key is not type of bool.
func (f *FieldDataWrapper) MustGetBool(key string) (bool, error) {
	valueInterface := f.Get(key)
	if valueInterface == nil {
		return false, f.errorResolve(key)
	}
	value, ok := valueInterface.(bool)
	if !ok {
		return false, f.errorTypeMismatch(key)
	}
	return value, nil
}

func (f *FieldDataWrapper) errorResolve(key string) error {
	return fmt.Errorf("failed to resolve value with key %v", key)
}

func (f *FieldDataWrapper) errorTypeMismatch(key string) error {
	return fmt.Errorf("failed to resolve value with key %v, type mismatch", key)
}
