package env

import (
	"github.com/criyle/go-judge/pkg/macsandbox"
	"github.com/criyle/go-judge/pkg/pool"
)

var defaultRead = []string{
	"/",
}

var defaultWrite = []string{
	"/tmp",
	"/dev/null",
	"/var/tmp",
}

// NewBuilder build a environment builder
func NewBuilder(c Config) (pool.EnvBuilder, error) {
	b := macsandbox.NewBuilder("", defaultRead, defaultWrite, c.NetShare)
	c.Info("created mac sandbox at", "")
	return b, nil
}
