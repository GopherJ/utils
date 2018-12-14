package util

/**
 * WordCount.go
 * 2017-12-24 12:01:21
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "bufio"
    "io"
    "strings"
)

func WordCount(r io.Reader)(map[string]int, error) {
    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords)

    words := map[string]int{}
    for scanner.Scan() {
        word := strings.ToLower(scanner.Text())
        words[word]++
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return words, nil
}
