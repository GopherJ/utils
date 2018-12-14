package util

/**
 * bar.go
 * 2017-12-31 21:30:01
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
    "os"
    "bytes"
    "math"
    "text/template"
    "errors"
    "strings"
)

type Bar struct {
    Filled  string
    Empty   string

    Total   float64
    Value   float64
    Percent float64

    Width   int
    Text    string

    Tpl     *template.Template
}

var (
    ValueOverflow = errors.New("Value Overflow")
)

func NewBar(total float64)(*Bar, error) {
    b :=  &Bar{
             Filled :  "█",
             Empty  :  "░",

             Total  :  total,
    }

    b.SetWidth(60)

    err := b.SetTemplate(`{{.Percent | printf "%3.0f"}}% {{.Text}}`)
    if err != nil {
        return nil, err
    }

    return b, nil
}

func NewBarInt(total int)(*Bar, error) {
    b, err := NewBar(math.Abs(float64(total)))
    if err != nil {
        return nil, err
    }
    return b, nil
}

func (b *Bar) SetTemplate(tpl string)(error) {
    t, err := template.New("").Parse(tpl)
    if err != nil {
        return err
    }

    b.Tpl = t
    return nil
}

func (b *Bar) SetText(text string)(*Bar) {
    b.Text = text
    return b
}

func (b *Bar) SetWidth(width int)(*Bar) {
    cols := Columns()
    b.Width = width * cols / 100
    return b
}

func (b *Bar) Tick(val float64)(error) {
    if b.Value + val > b.Total {
        return ValueOverflow
    }

    b.Value  += val
    b.Percent = b.Value / b.Total * 100

    return nil
}

func (b *Bar) Tickn(val int)(error) {
   err := b.Tick(math.Abs(float64(val)))
   if err != nil {
        return err
   }
   return nil
}

func (b *Bar) WriteTo(w *os.File) {
    P := b.Value / b.Total

    buf := new(bytes.Buffer)
    buf.WriteRune('\r')

    buf.WriteString(strings.Repeat(b.Filled, int(math.Floor(P * float64(b.Width)))))
    buf.WriteString(strings.Repeat(b.Empty, int(math.Ceil((1 - P) * float64(b.Width)))))

    b.Tpl.Execute(buf, b)
    buf.WriteTo(w)
}


