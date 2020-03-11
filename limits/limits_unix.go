// +build !windows,!plan9

package limits

import (
	"fmt"
	"syscall"

	log "github.com/p9c/logi"
)

const (
	fileLimitWant = 32768
	fileLimitMin  = 1024
)

// SetLimits raises some process limits to values which allow pod and associated utilities to run.
func SetLimits() error {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.L.Error(err)
		return err
	}
	if rLimit.Cur > fileLimitWant {
		return nil
	}
	if rLimit.Max < fileLimitMin {
		err = fmt.Errorf("need at least %v file descriptors",
			fileLimitMin)
		return err
	}
	if rLimit.Max < fileLimitWant {
		rLimit.Cur = rLimit.Max
	} else {
		rLimit.Cur = fileLimitWant
	}
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.L.Error(err)
		// try min value
		rLimit.Cur = fileLimitMin
		err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
		if err != nil {
			log.L.Error(err)
			return err
		}
	}
	return nil
}
