package helpers

func Map[T, U any](data []T, f func(T) U) []U {

	res := make([]U, 0, len(data))

	for _, e := range data {
		res = append(res, f(e))
	}

	return res
}

func Contains[T comparable](arr []T, item T) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}

	return false
}

func RemoveFromSliceByIndex[T any](slice []T, s int) []T {
	newArr := make([]T, (len(slice) - 1))
	newArr = append(slice[:s], slice[s+1:]...)
	return newArr
}
