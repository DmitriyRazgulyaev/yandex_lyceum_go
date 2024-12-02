package workingWithFiles

import (
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFilename string, startpos int) error {
	inputFile, err := os.OpenFile(inputFilename, os.O_RDONLY, 0600)

	if err != nil {
		return err
	}

	_, err2 := os.Create(outFilename)

	if err2 != nil {
		return err2
	}

	buffer := make([]byte, 100)

	_, err1 := inputFile.Seek(int64(startpos), 0)

	if err1 != nil {
		return err1
	}

	outputFile, err := os.OpenFile(outFilename, os.O_APPEND, 0600)

	if err != nil {
		return err
	}

	n, err3 := inputFile.Read(buffer)
	for err3 != io.EOF {
		n, err3 = inputFile.Read(buffer)
		if err3 != nil && err3 != io.EOF {
			return err3
		}

		err1 := os.WriteFile(outFilename, buffer[:n], 0666)

		if err1 != nil {
			return err1
		}
	}
	inputFile.Close()
	outputFile.Close()

	return nil
}
