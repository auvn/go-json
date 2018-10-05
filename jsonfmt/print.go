package jsonfmt

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Println(args ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, args...)
}

func Fprintln(w io.Writer, args ...interface{}) (n int, err error) {
	bb, err := json.MarshalIndent(args, "", " ")
	if err != nil {
		return 0, err
	}

	return fmt.Fprintf(w, string(bb))
}
