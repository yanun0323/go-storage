package internal

import (
	"fmt"
	"mime"
	"net/http"
	"sort"
)

func Response(format string, a ...any) any {
	return struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf(format, a...),
	}
}

func GetContentTypeAndExtension(data []byte) (ct, ext string) {
	// Detect content type
	contentType := http.DetectContentType(data)

	// Get extension from content type
	extensions, _ := mime.ExtensionsByType(contentType)

	sort.Slice(extensions, func(i, j int) bool {
		if extensions[i] == ".jpg" {
			return true
		}

		return len(extensions[i]) < len(extensions[j])
	})

	if len(extensions) != 0 {
		return contentType, extensions[0]
	}

	return contentType, ""
}
