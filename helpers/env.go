package helpers

import (
	"os"
)

// EnvOr takes an envionment variable string and if that env key is not set, returns the default which is passed as a
// seperate argument.
func EnvOr(e, d string) string {
	env := os.Getenv(e)
	if env == "" {
		env = d
	}
	return env
}
