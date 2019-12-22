package mergesort

// Mergesort is a parallel merge sort
func Mergesort(arr []int, workers int) []int {

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
	for j := 0; len(arr1) > 0; j, i = j+1, i+1 {
		arr[i] = arr1[j]
	}
	for j := 0; len(arr2) > 0; j, i = j+1, i+1 {
		arr[i] = arr2[j]
	}
	return arr
}
