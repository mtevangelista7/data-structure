package sequentialsearch

func Search(list []int, key int) int {
	for index, item := range list {
		if item == key {
			return index
		}
	}

	return -1
}
