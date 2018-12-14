package util

/**
 * makdown.go
 * 2017-12-29 23:50:27
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"regexp"
)

func LoadMarkdown(f string) ([]byte, error) {
	input, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	output := blackfriday.Run(input,
		blackfriday.WithExtensions(blackfriday.CommonExtensions))

	return output, nil
}

func MarkdownHeading(f string) ([]string, error) {
	md, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	re, err := regexp.Compile(`(?m:^#{1,6} .+$)`)
	if err != nil {
		return nil, err
	}

	return re.FindAllString(string(md), -1), nil
}
