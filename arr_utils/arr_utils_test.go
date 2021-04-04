package arr_utils

import "testing"

func TestFindPosByContainsValue(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	posValThree := 2 // pos 3 in arr

	res, err := FindPosByValueInt(arr, arr[posValThree])

	if res != posValThree || err != nil {
		t.Fatalf("FindPosByValue(%v, %v) = (%v, nil), but now res: %v, err: `%v`",
			arr, arr[posValThree], posValThree, res, err)
	}
}

func TestFindPosByNotContainsValue(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	res, err := FindPosByValueInt(arr, -1)

	if res != -1 || err == nil {
		t.Fatalf("FindPosByValue(%v, %v) = (-1, `arr is not contains this val`), but now res: %v, err: nil",
			arr, 0, res)
	}
}

func TestFindPosByValueWithVoidArr(t *testing.T) {
	var arr []int
	res, err := FindPosByValueInt(arr, 0)

	if res != -1 || err == nil {
		t.Fatalf("FindPosByValue(%v, %v) = (-1, `arr is void`), but now res: %v, err: nil",
			arr, 0, res)
	}
}
