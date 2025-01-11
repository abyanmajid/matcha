package reference

import (
	"fmt"
	"io"
	"net/http"

	"github.com/abyanmajid/matcha/logger"
)

func fetchContentFromURL(fileURL string, r *http.Request) (string, error) {
	if fileURL[0] == '/' {
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		fileURL = fmt.Sprintf("%s://%s%s", scheme, r.Host, fileURL)
	}

	resp, err := http.Get(fileURL)
	if err != nil {
		logger.Debug("HEYO")
		return "", fmt.Errorf("error getting file content: %w", err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Debug("222HEYO")
		return "", fmt.Errorf("error reading file content: %w", err)
	}

	return string(content), nil
}
