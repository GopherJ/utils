package util

/**
 * exec.go
 * 2017-12-28 20:59:10
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "os/exec"
    "runtime"
)

func Exec(s string)(string, error) {
    cmd := new(exec.Cmd)
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/C", s)
    } else if runtime.GOOS == "linux" {
        cmd = exec.Command("bash", "-c", s)
    }
    output, err := cmd.Output()
    if err != nil {
        return "", err
    }
    return string(output), nil
}

func MustExec(s string)(string) {
    output, err := Exec(s)
    if err != nil {
        panic(err)
    }
    return output
}
