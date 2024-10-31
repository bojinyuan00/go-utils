package comfunc

// StringInSlice string in slices(判断字符串是否在切片内)
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
