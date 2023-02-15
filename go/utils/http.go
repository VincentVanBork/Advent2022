package utils

import "io"

func ReadResponseBody(body io.ReadCloser) []byte {
	data, err := io.ReadAll(body)
	Check(err)
	defer body.Close()
	return data
}
