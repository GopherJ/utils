package util

/**
 * columns.go
 * 2017-12-27 19:05:24
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
	"os/exec"
    "bytes"
    "regexp"
    "io"
    "os"
    "runtime"
)

func Columns()(int) {
    output := new(bytes.Buffer)

    if runtime.GOOS == "windows" {
        cmd1 := exec.Command("mode", "con")
        cmd2 := exec.Command("findStr", "Columns")

        r, w := io.Pipe()
        cmd1.Stdout = w
        cmd2.Stdin  = r

        cmd2.Stdout = output

        cmd1.Start()
        cmd2.Start()

        go func(){
            defer w.Close()
            cmd1.Wait()
        }()
        cmd2.Wait()
    } else {
        cmd := exec.Command("echo", os.Getenv("COLUMNS"))
        cmd.Stdout = output
        cmd.Run()
    }

    re := regexp.MustCompile(`\d+`)
    width := re.FindString(output.String())

    return MustAtoi(width)
}
