package errors

import (
	"fmt"
	"os"
)

func Fatal(code int, msg string, err error) {
	fmt.Println(msg, err)
	os.Exit(code)
}
