package helper

import "strconv"

func StrToUint(source string) uint {
	tmp, _ := strconv.ParseUint(source, 10, 32)

	result := uint(tmp)

	return result
}
