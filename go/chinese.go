package util

/**
 * chinese.go
 * 2017-12-27 14:10:23
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "bytes"
    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
    "io/ioutil"
)

func Utf8ToGbk(in []byte) ([]byte, error){
    reader := transform.NewReader(bytes.NewReader(in), simplifiedchinese.GBK.NewEncoder())
    out, err := ioutil.ReadAll(reader)
    if err != nil {
        return nil, err
    }
    return out, nil
}

func GbkToUtf8(in []byte) ([]byte, error){
    reader := transform.NewReader(bytes.NewReader(in), simplifiedchinese.GBK.NewDecoder())
        out, err := ioutil.ReadAll(reader)
        if err != nil {
            return nil, err
        }
    return out, nil
}

func Utf8ToGb2312(in []byte) ([]byte, error){
    reader := transform.NewReader(bytes.NewReader(in), simplifiedchinese.HZGB2312.NewEncoder())
    out, err := ioutil.ReadAll(reader)
    if err != nil {
        return nil, err
    }
    return out, nil
}

func Gb2312ToUtf8(in []byte) ([]byte, error){
    reader := transform.NewReader(bytes.NewReader(in), simplifiedchinese.HZGB2312.NewDecoder())
    out, err := ioutil.ReadAll(reader)
    if err != nil {
        return nil, err
    }
    return out, nil
}
