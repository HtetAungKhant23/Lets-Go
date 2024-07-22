package fileutils

import "os"

func WriteFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644) // Owner: Read (4) + Write (2) = 6, Group: Read(4) = 4, Others: Read(4) = 4
}
