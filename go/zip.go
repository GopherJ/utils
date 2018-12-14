package util

/**
 * zip.go
 * 2017-12-27 22:08:42
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "archive/zip"
    "io"
    "os"
    "path/filepath"
    "strings"
)

func Zip(dst string, files []string)(error) {

    fd, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer fd.Close()

    w := zip.NewWriter(fd)
    defer w.Close()

    for _, file := range files {

        f, err := os.Open(file)
        if err != nil {
            return err
        }
        defer f.Close()

        s, err := f.Stat()
        if err != nil {
            return err
        }

        h, err := zip.FileInfoHeader(s)
        if err != nil {
            return err
        }

        h.Method = zip.Deflate

        wr, err := w.CreateHeader(h)
        if err != nil {
            return err
        }
        _, err = io.Copy(wr, f)
        if err != nil {
            return err
        }
    }
    return nil
}


func Unzip(src string, dst string)(error) {

    files := []string{}

    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()

    for _, f := range r.File {

        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        fpath := filepath.Join(dst, f.Name)
        files = append(files, fpath)

        if f.FileInfo().IsDir() {
            os.MkdirAll(fpath, os.ModePerm)
        } else {
            var fdir string
            if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
                fdir = fpath[:lastIndex]
            }

            err = os.MkdirAll(fdir, os.ModePerm)
            if err != nil {
                return err
            }
            f, err := os.OpenFile(
                fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer f.Close()

            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
    }
    return nil
}
