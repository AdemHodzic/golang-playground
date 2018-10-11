package reader

import (
	"bufio"
	"os"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var content []string
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return content
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
