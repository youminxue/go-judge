package pool

import (
	"sync"
	"time"

	"github.com/criyle/go-judge/pkg/envexec"
)

// Cgroup defines interface to limit and monitor resources consumption of a process
type Cgroup interface {
	SetMemoryLimit(envexec.Size) error
	SetProcLimit(uint64) error

	CPUUsage() (time.Duration, error)
	MemoryUsage() (envexec.Size, error)

	AddProc(int) error
	Reset() error
	Destroy() error
}

// CgroupPool implements pool of Cgroup
type CgroupPool interface {
	Get() (Cgroup, error)
	Put(Cgroup)
}

// CgroupListPool implements cgroup pool
type CgroupListPool struct {
	builder CgroupBuilder

	cgs []Cgroup
	mu  sync.Mutex
}

// NewCgroupListPool creates new cgroup pool
func NewCgroupListPool(builder CgroupBuilder) CgroupPool {
	return &CgroupListPool{builder: builder}
}

// Get gets cgroup from pool, if pool is empty, creates new one
func (w *CgroupListPool) Get() (Cgroup, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if len(w.cgs) > 0 {
		rt := w.cgs[len(w.cgs)-1]
		w.cgs = w.cgs[:len(w.cgs)-1]
		return rt, nil
	}

	cg, err := w.builder.Build()
	if err != nil {
		return nil, err
	}
	return (*wCgroup)(cg), nil
}

// Put puts cgroup into the pool
func (w *CgroupListPool) Put(c Cgroup) {
	w.mu.Lock()
	defer w.mu.Unlock()

	c.Reset()
	w.cgs = append(w.cgs, c)
}

// Shutdown destroy all cgroup
func (w *CgroupListPool) Shutdown() {
	w.mu.Lock()
	defer w.mu.Unlock()

	for _, c := range w.cgs {
		c.Destroy()
	}
}
