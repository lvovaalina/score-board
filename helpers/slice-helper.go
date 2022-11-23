package helpers

type mapf func(interface{}) interface{}

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
