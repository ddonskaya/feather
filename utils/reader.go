package utils

import "io"

func ReadData(conn io.Reader, buf []byte, size int) (int, error) {
	var (
		totalBytesRead = 0
		readErr        error
		bytesRead      = 0
	)

	for totalBytesRead < size && readErr == nil {
		bytesRead, readErr = conn.Read(buf[totalBytesRead:])
		totalBytesRead += bytesRead
	}

	return totalBytesRead, readErr
}