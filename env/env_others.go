// +build !windows,!linux,!darwin

package env

import (
	"errors"
	"runtime"

	"github.com/criyle/go-judge/pkg/pool"
)

func NewBuilder(c Config) (pool.EnvBuilder, error) {
	return nil, errors.New("environment is not support on this platform" + runtime.GOOS)
}
