package helpers

import (
	"fmt"
	"strconv"
)

func ParseToInt(source string) (int, error) {
	result, errConv := strconv.Atoi(source)
	if errConv != nil {
		fmt.Println("Error convering to int:", errConv)
		return -1, errConv
	}

	return result, nil
}
