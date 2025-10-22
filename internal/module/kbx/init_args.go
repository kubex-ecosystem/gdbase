// Package kbx provides utilities for working with initialization arguments.
package kbx

import (
	"os"
	"reflect"
)

type InitArgs struct {
	ConfigFile     string
	ConfigType     string
	EnvFile        string
	LogFile        string
	Name           string
	Debug          bool
	ReleaseMode    bool
	IsConfidential bool
	Port           string
	Bind           string
	Address        string
	PubCertKeyPath string
	PubKeyPath     string
	Pwd            string
}

func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetValueOrDefault[T any](value T, defaultValue T) (T, reflect.Type) {
	v := any(value)
	if v == nil {
		return defaultValue, reflect.TypeOf(v)
	}
	return value, reflect.TypeFor[T]()
}
