package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gtfo_search/binaries"
)

func main() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorYellow := "\033[33m"
	colorCyan := "\033[36m"
	bold := "\033[1m"

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		path_slice := strings.Split(scanner.Text(), "/")
		file_name := path_slice[len(path_slice)-1]

		_, ok := binaries.Data[file_name]
		if ok {
			fmt.Println(string(colorCyan), "File:", file_name, string(colorReset))

			for name, list_of_exploit_methods := range binaries.Data[file_name] {
				fmt.Println("[+]", string(bold), name, string(colorReset))
				for _, exploit := range list_of_exploit_methods {

					if exploit.Type == "description" {
						fmt.Println(string(colorYellow))
					} else {
						fmt.Print(string(colorRed))
					}
					fmt.Println(exploit.Value, string(colorReset))
				}
				fmt.Println()
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
