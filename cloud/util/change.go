package util

import (
	"strconv"
)

func StringToInt(str string) (num int) {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return
}
