package util

/**
 * color.go
 * 2017-12-27 10:31:47
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "github.com/mattn/go-colorable"
    "io"
    "os"
    "runtime"
)

func ColorWriteTo(w *os.File, s string) {
    if runtime.GOOS == "windows" {
        io.WriteString(colorable.NewColorable(w), s)
    } else {
        io.WriteString(w, s)
    }
}

func (b *Buffer) Cyan(s string)(*Buffer) {
    b.Append("\x1b[36m").Append(s).Append("\x1b[0m")
    return b
}

func (b *Buffer) Blue(s string)(*Buffer) {
    b.Append("\x1b[34m").Append(s).Append("\x1b[0m")
    return b
}

func (b *Buffer) Black(s string)(*Buffer) {
    b.Append("\x1b[30m").Append(s).Append("\x1b[0m")
    return b
}

func (b *Buffer) Red(s string)(*Buffer) {
    b.Append("\x1b[31m").Append(s).Append("\x1b[0m")
    return b
}

func (b *Buffer) Yellow(s string)(*Buffer) {
    b.Append("\x1b[33m").Append(s).Append("\x1b[0m")
    return b
}

func (b *Buffer) Green(s string)(*Buffer) {
    b.Append("\x1b[32m").Append(s).Append("\x1b[0m")
    return b
}

func (b *Buffer) Magenta(s string)(*Buffer) {
    b.Append("\x1b[35m").Append(s).Append("\x1b[0m")
    return b
}

func (b *Buffer) White(s string)(*Buffer) {
    b.Append("\x1b[37m").Append(s).Append("\x1b[0m")
    return b
}
