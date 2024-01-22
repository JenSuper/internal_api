package global

import "strings"

var (
	Token = "TOKEN"
	Data  = "DATA"
	Pre   = "Internal"
)

func BuildKeys(arg ...string) string {
	if len(arg) == 0 {
		return ""
	}
	var key strings.Builder
	key.WriteString(Pre)
	key.WriteString(":")
	for _, v := range arg {
		key.WriteString(strings.ReplaceAll(v, ":", "&"))
		key.WriteString("-")
	}
	return strings.TrimSuffix(key.String(), "-")
}

func BuildMultiKeys(arg ...string) string {
	if len(arg) == 0 {
		return ""
	}
	var key strings.Builder
	key.WriteString(Pre)
	key.WriteString(":")
	for _, v := range arg {
		key.WriteString(v)
		key.WriteString(":")
	}
	return strings.TrimSuffix(key.String(), ":")
}
