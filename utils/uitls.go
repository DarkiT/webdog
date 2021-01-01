package utils

import (
	"fmt"
	"os"
	"strings"
)

// return domain, port
func SegmentHost(host string) (string, string) {
	domain := ""
	port := "80"
	str := strings.Split(host, ":")
	// domain/ip
	domain = str[0]
	if len(str) > 1 {
		// port
		port = str[1]
	}
	return domain, port
}

// remove last rune
func RemoveLastRune(s string) string {
	r := []rune(s)
	return string(r[:len(r)-1])
}

// check file exists
func IsFileExisted(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

// mr (method##router)
func GetMr(method, router string) string {
	mr := fmt.Sprintf("%s##%s", method, router)
	return mr
}

// against injection attacks
func ContainsDotDot(v string) bool {
	if !strings.Contains(v, "..") {
		return false
	}
	for _, ent := range strings.FieldsFunc(v, isSlashRune) {
		if ent == ".." {
			return true
		}
	}
	return false
}

func isSlashRune(r rune) bool { return r == '/' || r == '\\' }
