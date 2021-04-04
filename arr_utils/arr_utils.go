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
