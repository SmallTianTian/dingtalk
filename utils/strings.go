package utils

import (
	"strconv"
)

func Int64S(i int64) string {
	return strconv.FormatInt(i, 10)
}

func IsEmpty(s string) bool {
	return s == ""
}
