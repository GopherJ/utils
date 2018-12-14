package util

/**
 * byteOrder.go
 * 2017-12-26 23:02:27
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "encoding/binary"
    "bytes"
    "unsafe"
)

func EndianReadFloat64(buf []byte, endian string, r *float64)(err error) {
    b := bytes.NewReader(buf)
    if endian == "LE" {
        err = binary.Read(b, binary.LittleEndian, r)
    }
    if endian == "BE" {
        err = binary.Read(b, binary.BigEndian, r)
    }
    return err
}

func EndianReadInt(buf []byte, endian string, r *int)(err error) {
    b := bytes.NewReader(buf)
    if endian == "LE" {
        err = binary.Read(b, binary.LittleEndian, r)
    }
    if endian == "BE" {
        err = binary.Read(b, binary.BigEndian, r)
    }
    return err
}

func EndianReadUint(buf []byte, endian string, r *uint)(err error) {
    b := bytes.NewReader(buf)
    if endian == "LE" {
        err = binary.Read(b, binary.LittleEndian, r)
    }
    if endian == "BE" {
        err = binary.Read(b, binary.BigEndian, r)
    }
    return err
}


func CpuEndian()(string) {
    t := 0x1
    const INT_SIZE = int(unsafe.Sizeof(0))

    b := (*[INT_SIZE]byte)(unsafe.Pointer(&t))
    if b[0] == 1 {
        return "BE"
    }
    return "LE"
}

