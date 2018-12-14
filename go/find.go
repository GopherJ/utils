package util

/**
 * find.go
 * 2017-12-29 15:30:19
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "path/filepath"
    "regexp"
    "os"
)

func Find(root string, pattern string)([]string, error) {
    root = filepath.Clean(root)

    root, err := filepath.Abs(root)
    if err != nil {
        return nil, err
    }

    re := regexp.MustCompile(pattern)
    files := []string{}

    err = filepath.Walk(root, func(path string, doc os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if doc.IsDir() {
            return nil
        }

        if re.MatchString(doc.Name()) {
            files = append(files, path)
        }

        return nil
    })

    if err != nil {
        return nil, err
    }

    return files, nil
}
