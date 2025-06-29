package utils

import (
	"github.com/rotisserie/eris"
	"os"
)

type EnvVarChecker struct {
	envVars []string
}

func NewEnvVarChecker(vars ...string) *EnvVarChecker {
	return &EnvVarChecker{
		envVars: vars,
	}
}

func (evc *EnvVarChecker) Check() error {
	for _, key := range evc.envVars {
		_, exists := os.LookupEnv(key)

		if !exists {
			return eris.Errorf("Env var %s not set", key)
		}
	}

	return nil
}
