package server

import (
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	cfg := NewConfig()

	var err error
	// run the server
	go func() {
		err = Run(cfg)
	}()

	// provide some time to run the server
	time.Sleep(2 * time.Second)
	cfg.StopCh <- syscall.SIGTERM
	// give some time for server shutdown
	time.Sleep(1 * time.Second)
	assert.Nil(t, err, "expected error nil, server should shutdown gracefully")
}
