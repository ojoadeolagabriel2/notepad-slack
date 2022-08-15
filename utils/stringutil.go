package utils

import "strconv"

func ToInt64(data string) int64 {
	if n, err := strconv.ParseInt(data, 10, 64); err == nil {
		return n
	}
	panic("could not parse " + data)
}
