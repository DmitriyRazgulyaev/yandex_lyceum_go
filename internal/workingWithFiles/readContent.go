package workingWithFiles

import "os"

func ReadContent(filename string) string {
	file, err := os.ReadFile(filename)

	if err != nil {
		return ""
	}
	return string(file)
}
