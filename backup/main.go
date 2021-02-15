package main 

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"path/filepath"
  "os/exec"
)

func processText(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fmt.Println(fileName +  "is opened")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		ruby := "ruby"
		cmd := "/tmp/bitcoin/address-gen.rb"
		res := exec.Command(ruby, cmd, scanner.Text())
		_, err := res.Output()
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(fileName + "is close")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getAllFile () []string{
	var files []string
	root := "datafile"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func main() {
	allFile := getAllFile()
	fmt.Println(allFile)
	for _, text := range allFile {
		if text != "datafile" {
			fmt.Println(text + "is runing")
			go processText(text)
		}
	}
	select{}
}

