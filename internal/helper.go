package internal

import (
	"fmt"
	"mime"
	"net/http"
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

	if len(extensions) != 0 {
		return contentType, extensions[0]
	}

	return contentType, ""
}
