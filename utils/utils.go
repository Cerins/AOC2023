package utils

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Map[T, U any](ts []T, f func(T, int) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i], i)
	}
	return us
}

func FirstOcc[T comparable](arr []T, item T) int {
	for i, vitem := range arr {
		if item == vitem {
			return i
		}
	}
	return -1

}

func IndexOf[T comparable](element T, data []T) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func IsDigit(c byte) bool {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func ToDigit(c byte) int {
	return int(c - '0')
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
