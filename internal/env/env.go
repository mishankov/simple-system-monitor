package env

import (
	"os"
	"strconv"
	"strings"
)

type Env struct{}

func New() *Env {
	return &Env{}
}

func (e *Env) GetStringOrDefault(name, def string) string {
	value := os.Getenv(name)

	if len(value) == 0 {
		return def
	}

	return strings.TrimSpace(value)
}

func (e *Env) GetIntOrDefault(name string, def int) (int, error) {
	strValue := e.GetStringOrDefault(name, "")

	if len(strValue) == 0 {
		return def, nil
	}

	return strconv.Atoi(strValue)
}
