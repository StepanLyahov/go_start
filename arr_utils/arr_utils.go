package arr_utils

import "errors"

func FindPosByValueInt(arr []int, val int) (int, error) {
	if len(arr) == 0 {
		return -1, errors.New("arr is void")
	}

	for pos, valArr := range arr {
		if valArr == val {
			return pos, nil
		}
	}

	return -1, errors.New("arr is not contains this val")
}

func BubbleSort(arr []int) {
	len := len(arr)

	if len == 0 || len == 1 {
		return
	}

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if arr[i] <= arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}
