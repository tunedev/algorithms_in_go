package main

func main() {

}

func binarySearch(haystack []int, needle int) bool {
	lo := 0
	hi := len(haystack)
	for lo < hi {
		mid := lo + (hi-lo)/2
		midVal := haystack[mid]
		if midVal == needle {
			return true
		}
		if midVal > needle {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return false
}
