## Goutil
```
Some outils for golang
```

## Install
```go
go get github.com/gopherj/goutil
```

## Example
```go
package main

import (
    "github.com/gopherj/goutil"
    "fmt"
    "os"
)
func main(){
    fd, err := os.Open("file.txt")
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error while opening file")
    }
    lines, err := util.ReadLine(fd)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error while reading lines")
    }

    fmt.Println(lines)
}
```

## List
```golang
func NewBuffer()(*util.Buffer)
func (b *util.Buffer) Append(interface{})(*util.Buffer)
func (b *util.Buffer) Cyan(s string)(*util.Buffer) 
func (b *util.Buffer) Blue(s string)(*util.Buffer) 
func (b *util.Buffer) Black(s string)(*util.Buffer) 
func (b *util.Buffer) Red(s string)(*util.Buffer) 
func (b *util.Buffer) Yellow(s string)(*util.Buffer) 
func (b *util.Buffer) Green(s string)(*util.Buffer) 
func (b *util.Buffer) Magenta(s string)(*util.Buffer) 
func (b *util.Buffer) White(s string)(*util.Buffer) 
func ColorWriteTo(w *os.File, s string) 

func ReadLine(r io.Reader)([]string, error)
func WordCount(r io.Reader)(map[string]int, error)

func RStringSlice(s []string)([]string) 
func RIntSlice(s []int)([]int) 

func MustAtoi(s string)(int) 
func Columns()(int)
func TimeNow()(string) 


func Utf8ToGbk(in []byte) ([]byte, error)
func GbkToUtf8(in []byte) ([]byte, error)
func Utf8ToGb2312(in []byte) ([]byte, error)
func Gb2312ToUtf8(in []byte) ([]byte, error)

func PL(s, r string, n int) string 
func PR(s, r string, n int) string 
func PC(s, r string, n int) string 

func GenPrime(bits int)(*big.Int) 
func NewBigInt(s string, b int)(*big.Int) 
func IsProbablePrime(n *big.Int, p int)(bool) 


func Index(s []string, t string)(int) 
func Include(s []string, t string)(bool) 
func Map(s []string, f func(string)string)([]string) 
func Filter(s []string, f func(string)bool)([]string) 
func All(s []string, f func(string)bool)(bool) 
func Any(s []string, f func(string)bool)(bool) 

func Zip(dst string, files []string)(error) 
func Unzip(src string, dst string)(error)

func CsvWriteFile(f string, data [][]string)(error)
func CsvReadFile(f string)([][]string, error)
func CsvHeading(f string)([]string)


func IsNumber(s string)(bool) 
func IsChinese(s string)(bool) 
func IsEnglish(s string)(bool) 
func IsEmail(s string)(bool) 
func IsPhoneOfChina(s string)(bool) 
func IsIdentityCard(s string)(bool) 

func Exec(s string)(string, error) 
func MustExec(s string)(string) 


func RandToken(n int)(string) 

func Find(root string, pattern string)([]string, error) 

func Download(url string)(error) 
func ContentLength(url string)(int, error)
func DownloadWithBar(url string)(error) 

func LoadMarkdown(f string)([]byte, error) 
```
