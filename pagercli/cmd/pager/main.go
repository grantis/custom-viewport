package main

import (
	"fmt"
	"io"
	"os"

	"github.com/yourusername/viewportpager"
)

func main() {
	var content, title string

	// Check if a file argument is provided
	if len(os.Args) > 1 {
		filename := os.Args[1]
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Could not read file:", err)
			os.Exit(1)
		}
		content = string(data)
		title = filename
	} else {
		// Check for piped input
		fi, err := os.Stdin.Stat()
		if err != nil {
			fmt.Println("Could not stat STDIN:", err)
			os.Exit(1)
		}

		if fi.Mode()&os.ModeCharDevice == 0 {
			bytes, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println("Could not read from STDIN:", err)
				os.Exit(1)
			}
			content = string(bytes)
			title = "Piped Input"
		} else {
			content = "Welcome to the Terminal Pager!"
			title = "Go View Pager"
		}
	}

	// Start the pager
	if err := viewportpager.StartPager(content, title); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
