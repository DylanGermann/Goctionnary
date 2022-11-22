package main

import (
	"flag"
	"fmt"
	"os"

	"projets.perso/Goctionnary/dictionary"
)

func main() {
	action := flag.String("action", "list", "Action to perform on the CLI")

	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()

	flag.Parse()
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Arg(0))
	case "remove":
		actionDelete(d, flag.Arg(0))
	default:
		fmt.Printf("Unknown action: %v\n", *action)
	}
}

func actionDefine(d *dictionary.Dictionary, word string) {
	entry, err := d.Get(word)
	handleErr(err)
	fmt.Println(entry)
}

func actionDelete(d *dictionary.Dictionary, word string) {
	err := d.Remove(word)
	handleErr(err)
	fmt.Printf("'%v' was remove from the dictionnary\n", word)
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleErr(err)
	fmt.Printf("'%v' added to the dictionay\n", word)
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleErr(err)
	fmt.Println("Dictionnary content")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionnary error : %v", err)
		os.Exit(1)
	}
}
