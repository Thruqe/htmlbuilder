package htmlbuilder

import "strconv"

// px converts a numeric value to a px-suffixed CSS string.
// px(16) == "16px"
func px[T int | float64](v T) string {
	return formatNum(v) + "px"
}

// pct converts a numeric value to a percentage CSS string.
// pct(50) == "50%"
func pct[T int | float64](v T) string {
	return formatNum(v) + "%"
}

// rem converts a numeric value to a rem-suffixed CSS string.
// rem(1.5) == "1.5rem"
func rem[T int | float64](v T) string {
	return formatNum(v) + "rem"
}

func Px[T int | float64](v T) string  { return formatNum(v) + "px" }
func Pct[T int | float64](v T) string { return formatNum(v) + "%" }
func Rem[T int | float64](v T) string { return formatNum(v) + "rem" }

func formatNum[T int | float64](v T) string {
	switch n := any(v).(type) {
	case int:
		return strconv.Itoa(n)
	case float64:
		return strconv.FormatFloat(n, 'f', -1, 64)
	default:
		return ""
	}
}
