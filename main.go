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
	d.Add("python", "A interpreted langage")
	words, entries, _ := d.List()
	for _, words := range words {
		fmt.Println(entries[words])
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionnary error : %v", err)
		os.Exit(1)
	}
}
