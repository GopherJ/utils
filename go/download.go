package util

/**
 * download.go
 * 2017-12-29 22:42:23
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "io/ioutil"
    "net/http"
    "net/url"
    "path"
    "io"
    "os"
)

func Download(URL string)(error) {
    resp, err := http.Get(URL)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    u, err := url.Parse(URL)
    if err != nil {
        return err
    }
    fileName := path.Base(u.Path)

    err = ioutil.WriteFile(fileName, data, 0755)
    if err != nil {
        return err
    }

    return nil
}

func ContentLength(url string)(int, error) {
    resp, err := http.Head(url)
    if err != nil {
        return 0, err
    }

    return MustAtoi(resp.Header.Get("Content-Length")), nil
}

func DownloadWithBar(URL string)(error) {
    Size, err := ContentLength(URL)
    if err != nil {
        return err
    }

    u, err := url.Parse(URL)
    if err != nil {
        return err
    }
    fileName := path.Base(u.Path)
    fd, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    bar, err := NewBarInt(Size)
    if err != nil {
        return err
    }
    bar.SetText(fileName)

    resp, err := http.Get(URL)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    b := make([]byte, 1024)
    for {
        n, err := resp.Body.Read(b)
        if err == io.EOF {
            break
        }

        if err != nil {
            return err
        }

        fd.Write(b[0:n])
        bar.Tickn(n)
        bar.WriteTo(os.Stdout)
    }

    return nil
}
