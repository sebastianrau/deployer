package verbose

import (
	"fmt"
	"io"
)

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	if w != nil {
		return fmt.Fprintf(w, format, a...)
	}
	return 0, nil
}
