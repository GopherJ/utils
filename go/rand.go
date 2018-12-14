package util

/**
 * rand.go
 * 2017-12-28 20:48:48
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "crypto/rand"
    "fmt"
)

func RandToken(n int)(string) {
    b := make([]byte, n)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}
