package util

/**
 * ReadLine.go
 * 2017-12-24 11:48:09
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "bufio"
    "io"
)

func ReadLine(r io.Reader) (s []string, err error) {
    R := bufio.NewReader(r)

    for {
        line, err := R.ReadString('\n')

        if err == io.EOF {
            break
        }

        if err != nil {
            return nil, err
        }

        s = append(s, line)
    }

    return s, nil
}
