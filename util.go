package main
import (
	"errors"
	"fmt"
	"runtime"
)


func Err(err error, args ...interface{}) error {
	_, file, line, _ := runtime.Caller(1)
	loc := fmt.Sprintf(" @ %s : %d\n", file, line)

	msg := loc + fmt.Sprint(args)
	if err == nil {
		err = errors.New(msg)		
	}
	return errors.New(err.Error() + "\n!!-- " + msg)
}
