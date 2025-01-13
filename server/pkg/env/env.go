package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("env variable %s is not found", key)
	}
	// remove quotes that were read from Makefile
	value = strings.Trim(value, "\"")
	return value, nil
}

func GetEnvAsInt(key string) (int, error) {
	valueStr, err := GetEnv(key)
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func GetEnvAsBool(key string) (bool, error) {
	valueStr, err := GetEnv(key)
	if err != nil {
		return false, err
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return false, err
	}
	return value, nil
}

func MustGetEnv(key string) string {
	v, err := GetEnv(key)
	if err != nil {
		panic(err)
	}
	return v
}

func MustGetEnvAsInt(key string) int {
	v, err := GetEnvAsInt(key)
	if err != nil {
		panic(err)
	}
	return v
}

func MustGetEnvAsBool(key string) bool {
	v, err := GetEnvAsBool(key)
	if err != nil {
		panic(err)
	}
	return v
}
