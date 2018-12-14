package util

/**
 * Buffer.go
 * 2017-12-24 15:23:20
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "bytes"
    "strconv"
)


type Buffer struct {
    *bytes.Buffer
}

func NewBuffer() *Buffer {
    return &Buffer{
        bytes.NewBuffer([]byte{}),
    }
}

func (b *Buffer) Append(s interface{}) *Buffer {
    switch val := s.(type) {
    case int:
        b.WriteString(strconv.Itoa(val))
    case uint:
        b.WriteString(strconv.FormatUint(uint64(val), 10))
    case rune:
        b.WriteRune(val)
    case byte:
        b.WriteByte(val)
    case string:
        b.WriteString(val)
    case []byte:
        b.Write(val)
    case []rune:
        for _, r := range val {
            b.WriteRune(r)
        }
    }

    return b
}
