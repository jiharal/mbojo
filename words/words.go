package words

import (
	"strings"
)

// tamat (true)
// eko (false)
func Func1(val string) bool {
	arr := strings.Split(val, "")
	arr1 := []string{}
	for i, _ := range arr {
		arr1 = append(arr1, arr[(len(arr)-i)-1])
	}

	val1 := strings.Join(arr, "")
	val2 := strings.Join(arr1, "")

	if val1 == val2 {
		return true
	}
	return false
}
