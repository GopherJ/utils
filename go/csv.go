package util

/**
 * csv.go
 * 2017-12-27 22:39:46
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "encoding/csv"
    "os"
)

func CsvWriteFile(f string, data [][]string)(error) {
    fd, err := os.Create(f)
    if err != nil {
        return err
    }
    defer fd.Close()
    w := csv.NewWriter(fd)
    return w.WriteAll(data)
}

func CsvReadFile(f string)([][]string, error) {
    fd, err := os.Open(f)
    if err != nil {
        return nil, err
    }
    defer fd.Close()
    r := csv.NewReader(fd)
    return r.ReadAll()
}

func CsvHeading(f string)([]string) {
    data, err := CsvReadFile(f)
    if err != nil {
        panic(err)
    }
    return data[0]
}
