package mergesort

// Mergesort is a parallel merge sort
func Mergesort(arr []int, workers int) []int {
	if len(arr) <= 1 {
		return arr
	}

	var arr1, arr2 []int
	mid := len(arr) / 2

	if workers > 1 {
		chanel := make(chan int)
		numWorkers1 := workers / 2
		go func() {
			arr1 = Mergesort(arr[:mid], numWorkers1)
			chanel <- 1
		}()
		arr2 = Mergesort(arr[mid:], workers-numWorkers1)
		<-chanel // Wait untill the left part is sorted
	} else {
		arr1 = Mergesort(arr[:mid], 1)
		arr2 = Mergesort(arr[mid:], 1)
	}
	return merge(arr1, arr2)
}

func merge(arr1, arr2 []int) []int {
	arr := make([]int, len(arr1)+len(arr2))

	i := 0
	for ; len(arr1)*len(arr2) > 0; i++ {
		if arr1[0] < arr2[0] {
			arr[i] = arr1[0]
			arr1 = arr1[1:]
		} else {
			arr[i] = arr2[0]
			arr2 = arr2[1:]
		}
	}
	for j := 0; j < len(arr1); j, i = j+1, i+1 {
		arr[i] = arr1[j]
	}
	for j := 0; j < len(arr2); j, i = j+1, i+1 {
		arr[i] = arr2[j]
	}
	return arr
}
