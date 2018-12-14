package util

/**
 * slice.go
 * 2017-12-27 13:20:19
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

func Index(s []string, t string)(int) {
    for i, v := range s {
        if v == t {
            return i
        }
    }
    return -1
}

func Include(s []string, t string)(bool) {
    return Index(s, t) != -1
}

func Map(s []string, f func(string)string)([]string) {
    o := make([]string, len(s))
    for i, v := range s {
        o[i] = f(v)
    }
    return o
}

func Filter(s []string, f func(string)bool)(o []string) {
   for _, v := range s {
        if f(v) {
            o = append(o, v)
        }
   }
   return o
}

func All(s []string, f func(string)bool)(bool) {
    for _, v := range s {
        if !f(v) {
            return false
        }
    }
    return true
}

func Any(s []string, f func(string)bool)(bool) {
    for _, v := range s {
        if f(v) {
            return true
        }
    }
    return false
}
