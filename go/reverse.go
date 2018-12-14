package util

/**
 * reverse.go
 * 2017-12-26 23:45:55
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

func RStringSlice(in []string)(out []string) {
    if len(in) == 0 {
        return []string{}
    }

    for i:=len(in)-1; i>=0; i-- {
       out = append(out, in[i])
    }
    return out
}

func RIntSlice(in []int)(out []int) {
    if len(in) == 0 {
        return []int{}
    }

    for i:=len(in)-1; i>=0; i-- {
       out = append(out, in[i])
    }
    return out
}

func RString(in string)(string) {
    L := len(in)
    b := []byte{}

    for i:=L-1; i>=0; i-- {
        b = append(b, in[i])
    }

    return string(b)
}
