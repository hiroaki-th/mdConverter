package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gomarkdown/markdown"
)

func main() {
	start := time.Now()
	if len(os.Args) < 3 {
		fmt.Println("please input mdFileName and new fileName that is exported in HTML")
		fmt.Println("you must run this command like 'go run main.go [.md fileName] [output fileName]'")
		return
	}

	fileName := os.Args[1]

	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("cannot read file: %s", err)
		return
	}

	html := markdown.ToHTML(content, nil, nil)

	newFile, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Printf("cannot create new file: %s", err)
		return
	}

	size, err := newFile.WriteString(string(html))
	if err != nil {
		fmt.Printf("cannot write binary to new file: %s", err)
		return
	}

	fmt.Printf("total binary size: %d", size)
	fmt.Println("successfully export new file. please check html file")
	fmt.Println("this process ended in ", time.Since(start))
}
