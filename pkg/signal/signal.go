// Package signal -----------------------------
// @file      : signal.go
// @author    : Jimmy
// @contact   : 2114272829@qq.com
// @time      : 2024/7/29 下午5:13
// -------------------------------------------
package signal

import (
	"github.com/ZUCCzwp/kitex_ddd_example/pkg/onelog"
	"os"
	"os/signal"
	"syscall"
)

var done = make(chan struct{})

func init() {
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

		for {
			v := <-signals
			switch v {
			case syscall.SIGTERM:
				stopOnSignal()
				gracefulStop(signals, syscall.SIGTERM)
			case syscall.SIGINT:
				stopOnSignal()
				gracefulStop(signals, syscall.SIGINT)
			default:
				onelog.Errorf("got unregistered signal:%v", v)
			}
		}
	}()
}

// Done returns the channel that notifies the process quitting.
func Done() <-chan struct{} {
	return done
}

func stopOnSignal() {
	select {
	case <-done:
		// already closed
	default:
		close(done)
	}
}
