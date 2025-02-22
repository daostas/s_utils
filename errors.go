package s_utils

import (
	"errors"
	"fmt"
)

func Error(httpStatus int, args ...interface{}) error {
	msg := fmt.Sprintf("%d", httpStatus)
	for _, arg := range args {
		msg += fmt.Sprintf(" %v", arg)
	}
	return errors.New(msg)
}
