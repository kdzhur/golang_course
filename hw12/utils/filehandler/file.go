package filehandler

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadAll(file *os.File) []byte {
	var res []byte

	reader := bufio.NewReader(file)
	buf := make([]byte, 16)

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatalln(err)
			}
			break
		}
		res = append(res, buf[0:n]...)
	}
	return res
}

func CreateAndWriteToFile(text []byte, filePath string) {
	fo, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	if _, err := fo.Write(text); err != nil {
		panic(err)
	}
}
