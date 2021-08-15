package utils

import (
	"fmt"
)

func Handle(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg)
	}
}


