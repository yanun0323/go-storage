package internal

import "fmt"

func Response(format string, a ...any) any {
	return struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf(format, a...),
	}
}
