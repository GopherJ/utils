package util

/**
 * validate.go
 * 2017-12-28 14:05:53
 *
 * GopherJ<cocathecafe@gmail.com>
 * https://github.com/GopherJ
 */

import (
	"regexp"
)

func IsNumber(s string)(bool) {
	if m, _ := regexp.MatchString(`^[0-9]+$`, s); !m {
		return false
	}
	return true
}

func IsChinese(s string)(bool) {
	if m, _ := regexp.MatchString(`^[\x{4e00}-\x{9fa5}]+$`, s); !m {
		return false
	}
	return true
}

func IsEnglish(s string)(bool) {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", s); !m {
		return false
	}
	return true
}

func IsEmail(s string)(bool) {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, s); !m {
		return false
	} else {
		return true
	}
}

func IsPhoneOfChina(s string)(bool) {
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, s); !m {
		return false
	}
	return true
}

func IsIdentityCard(s string)(bool) {
	if m, _ := regexp.MatchString(`^(\d{15})$`, s); m {
		return true
	}
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, s); m {
		return true
	}
	return false
}
