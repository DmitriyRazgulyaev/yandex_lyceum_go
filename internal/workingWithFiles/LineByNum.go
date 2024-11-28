package workingWithFiles

import (
	"bufio"
	"os"
)

func LineByNum(inputFileName string, lineNum int) string {
	if lineNum < 0 {
		return ""
	}
	file, err := os.Open(inputFileName)
	if err != nil {
		return ""
	}

	fileScanner := bufio.NewScanner(file)

	defer file.Close()
	for fileScanner.Scan() {
		if lineNum == 0 {
			return fileScanner.Text()
		}
		lineNum--
	}
	return ""
}
