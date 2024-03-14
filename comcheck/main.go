package main

import "os"

func main() {
	s := os.Args[1:]
	for _, ch := range s {
		if ch == "galaxy" || ch == "01" || ch == "galaxy 01" {
			println("Alert !!!")
			break
		}
	}
}
