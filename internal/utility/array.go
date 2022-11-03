package utility

// Pop the last element from a slice
func Pop(arr *[]uint) uint {
	len := len(*arr)
	rv := (*arr)[len-1]
	*arr = (*arr)[:len-1]
	return rv
}
