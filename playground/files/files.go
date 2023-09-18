package files

import (
	"bufio"
	"fmt"
	"os"
)

func WriteToFile(fileName string, content string) error {
	path := fmt.Sprintf("./playground/examples/%s", fileName)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprint(f, content)
	return nil
}

func ReadFromFile(fileName string) (string, error) {
	path := fmt.Sprintf("./playground/examples/%s", fileName)
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if scanner.Scan() {
		return scanner.Text(), nil
	} else {
		return "", nil
	}
}
