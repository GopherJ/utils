package util

/**
 * time.go
 * 2017-12-27 14:12:45
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
	"time"
)

const (
	DateTime = "2006-01-02 15:04:05"
)

func Time() string {
	return time.Now().Local().Format(DateTime)
}
