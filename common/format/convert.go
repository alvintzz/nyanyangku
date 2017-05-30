package format

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//ToString convert type (int, int64, float32, float64, byte, and []bytes) to string
//Parameter p is optional and only used in converting float
func ToString(n interface{}, p ...int) string {
	var t string

	switch n.(type) {
	case int:
		t = strconv.Itoa(n.(int))
	case int64:
		t = strconv.FormatInt(n.(int64), 10)
	case float32:
		if len(p) > 0 {
			t = strconv.FormatFloat(float64(n.(float32)), 'f', p[0], 64)
		} else {
			t = strconv.FormatFloat(float64(n.(float32)), 'f', -1, 64)
		}
	case float64:
		if len(p) > 0 {
			t = strconv.FormatFloat(n.(float64), 'f', p[0], 64)
		} else {
			t = strconv.FormatFloat(n.(float64), 'f', -1, 64)
		}
	case byte:
		t = string(n.(byte))
	case []byte:
		t = string(n.([]byte))
	case string:
		t = n.(string)
	}

	return t
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func ToInt(str string) int {
	value, _ := strconv.Atoi(str)
	return value
}

func ToInt64(str string) int64 {
	value, _ := strconv.ParseInt(str, 10, 64)
	return value
}

func ToFloat64(str string) float64 {
	value, _ := strconv.ParseFloat(str, 64)
	return value
}

func ArrayToInt(strs []string) []int {
	arr := []int{}
	for _, val := range strs {
		arr = append(arr, ToInt(val))
	}
	return arr
}

func ArrayToInt64(strs []string) []int64 {
	arr := []int64{}
	for _, val := range strs {
		arr = append(arr, ToInt64(val))
	}
	return arr
}

func ArrayToFloat64(strs []string) []float64 {
	arr := []float64{}
	for _, val := range strs {
		arr = append(arr, ToFloat64(val))
	}
	return arr
}

func getIndonesianMonth(month time.Month) string {
	switch month {
	case 1:
		return "Januari"
	case 2:
		return "Februari"
	case 3:
		return "Maret"
	case 4:
		return "April"
	case 5:
		return "Mei"
	case 6:
		return "Juni"
	case 7:
		return "Juli"
	case 8:
		return "Agustus"
	case 9:
		return "September"
	case 10:
		return "Oktober"
	case 11:
		return "November"
	default:
		return "Desember"
	}
}

func DateFormat(t time.Time) string {
	monthString := getIndonesianMonth(t.Month())
	layout := fmt.Sprintf("2 %s 2006", monthString)
	return t.Format(layout)
}

func DateTimeFormat(t time.Time) string {
	monthString := getIndonesianMonth(t.Month())
	tzString := "WIB"
	layout := fmt.Sprintf("2 %s 2006, 15:04 %s", monthString, tzString)
	return t.Format(layout)
}

func LPad(s string, l int, c string) string {
	lr := len([]rune(s))
	if lr > l {
		return s
	}
	r := []rune(strings.Repeat(c, l-lr) + s)
	return string(r[len(r)-l:])
}

func RPad(s string, l int, c string) string {
	lr := len([]rune(s))
	if lr > l {
		return s
	}
	r := []rune(s + strings.Repeat(c, l))
	return string(r[:l])
}
