package sharedutils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// Int32Pointer ...
func Int32Pointer(i int32) *int32 {
	return &i
}

// Ptr2Str is Pointer to String if nil return ""
func Ptr2Str(value *string) string {
	if value != nil {
		return *value
	}
	return ""
}

// Ptr2Int is Pointer to Int if nil return -1
func Ptr2Int(value *int) int {
	if value != nil {
		return *value
	}
	return -1
}

// Ptr2Float is Pointer to Float if nil return -1
func Ptr2Float(value *float64) float64 {
	if value != nil {
		return *value
	}
	return -1
}

// Str2Int64 is Pointer to Float if nil return -1
func Str2Int64(value string) int64 {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return -1
	}
	return i
}

// Str2Unix is String to Unix
func Str2Unix(value string) time.Time {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Unix(i, 0)
}

// Str2Uint is String to Uint
func Str2Uint(value string) uint {
	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return uint(i)
}

// Str2Duration is String to Duration
func Str2Duration(value string) time.Duration {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Duration(i) * time.Second
}

// UInt2Str is UInt to String if nil return ""
func UInt2Str(value uint) string {
	return strconv.FormatUint(uint64(value), 10)
}

// Str2Ptr is Pointer to String if nil return ""
func Str2Ptr(value string) *string {
	return &value
}

// StrToKg is string to float64 with type
func StrToKg(input string) (float64, string) {
	rawData := strings.Split(input, " ")
	if len(rawData) != 2 {
		return 0, ""
	}
	value, err := strconv.ParseFloat(rawData[0], 64)
	if err != nil {
		return 0, ""
	}
	return value, rawData[2]
}

func PriceCalc(totalPrice float64) (float64, float64) {
	vatInt := int64(math.Round(totalPrice / float64(11) * 100))
	unitPrice := float64(int64(totalPrice*100)-vatInt) / 100
	vat := float64(vatInt) / 100

	strUnitPrice := fmt.Sprintf("%.2f", unitPrice)
	strVAT := fmt.Sprintf("%.2f", vat)

	uprice, _ := strconv.ParseFloat(strUnitPrice, 64)
	vprice, _ := strconv.ParseFloat(strVAT, 64)

	return uprice, vprice
}
