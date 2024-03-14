package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1:]
	for _, ch := range s {
		if ch == "galaxy" || ch == "01" || ch == "galaxy 01" {
			fmt.Println("Alert !!!")
			break
		}
	}
}
