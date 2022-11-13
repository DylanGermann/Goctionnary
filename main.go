package main

import (
	"fmt"
	"os"

	"projets.perso/Goctionnary/dictionary"
)

func main() {
	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionnary error : %v", err)
		os.Exit(1)
	}
}
