package functions

import (
	"strconv"
	"strings"
)

func GenerateCode(oldCode, charAt string) string {
	splitKey := strings.SplitAfter(oldCode, charAt)
	convertToInt, _ := strconv.Atoi(splitKey[1])
	count := convertToInt + 1
	resultReplace := strings.Replace(splitKey[1], strconv.Itoa(convertToInt), strconv.Itoa(count), 8)
	newCode := splitKey[0] + resultReplace
	return newCode
}
