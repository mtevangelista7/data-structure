package binarysearch

// lista ordenada
func Search(list []int, key int) int {
	for index, item := range list {
		if item == key {
			return index
		}

		mid := len(list) / 2

		if list[mid] < key {
			list = list[mid:]
		} else {
			list = list[:mid]
		}
	}

	return -1
}
