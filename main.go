package main

import (
	"bufio"
	"fmt"
	"os"
)

func openFile(filepath string) {
	marks := []string{}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening file", err.Error())
	}

}

func main() {
	fmt.Println("this is working")
}
