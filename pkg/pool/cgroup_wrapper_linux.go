package pool

import (
	"time"

	"github.com/criyle/go-judge/pkg/envexec"
	"github.com/criyle/go-sandbox/pkg/cgroup"
)

var (
	_ Cgroup = &wCgroup{}
)

type wCgroup cgroup.Cgroup

func (c *wCgroup) SetMemoryLimit(s envexec.Size) error {
	return (*cgroup.Cgroup)(c).SetMemoryLimitInBytes(uint64(s))
}

func (c *wCgroup) SetProcLimit(l uint64) error {
	return (*cgroup.Cgroup)(c).SetPidsMax(l)
}

func (c *wCgroup) CPUUsage() (time.Duration, error) {
	t, err := (*cgroup.Cgroup)(c).CpuacctUsage()
	return time.Duration(t), err
}

func (c *wCgroup) MemoryUsage() (envexec.Size, error) {
	s, err := (*cgroup.Cgroup)(c).MemoryMaxUsageInBytes()
	if err != nil {
		return 0, err
	}
	return envexec.Size(s), nil
	// not really useful if creates new
	// cache, err := (*cgroup.CGroup)(c).FindMemoryStatProperty("cache")
	// if err != nil {
	// 	return 0, err
	// }
	// return envexec.Size(s - cache), err
}

func (c *wCgroup) AddProc(pid int) error {
	return (*cgroup.Cgroup)(c).AddProc(pid)
}

func (c *wCgroup) Reset() error {
	if err := (*cgroup.Cgroup)(c).SetCpuacctUsage(0); err != nil {
		return err
	}
	if err := (*cgroup.Cgroup)(c).SetMemoryMaxUsageInBytes(0); err != nil {
		return err
	}
	return nil
}

func (c *wCgroup) Destroy() error {
	return (*cgroup.Cgroup)(c).Destroy()
}
