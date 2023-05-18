package main

import (
	"fmt"
	"bufio"
	"os"
)

type sastavdala struct {
	veids, modelis string
	cena float64
	id int
}
type list struct {
	all []sastavdala
}

func Add(Scanner *bufio.Scanner, l list) list {
	var new sastavdala
	for new.veids == "" {
	fmt.Println("Veids: ")
	Scanner.Scan()
	new.veids = Scanner.Text()		
	}
	for new.modelis == "" {
	fmt.Println("Modelis: ")
	Scanner.Scan()
	new.modelis = Scanner.Text()
	}
	for new.cena == 0 {
	fmt.Println("Cena: ")
	fmt.Scan(&new.cena)		
	}
	new.id = GiveID(l)
	l.all = append(l.all, new)
	return l
}
func GiveID(l list) int {
	return len(l.all)+1
}
func List(l list){
		fmt.Println("-Personālā datora sastāvdaļa-")	
	for i := range l.all {
		fmt.Printf("%v)\n", i+1)
		fmt.Printf("ID: %v\n",l.all[i].id)
		fmt.Printf("Veids: %s\n",l.all[i].veids)
		fmt.Printf("Modelis: %s\n", l.all[i].modelis)
		fmt.Printf("Cena: %v\n", l.all[i].cena)
	}
}
func Change(l list) list {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("starting...")
	var action int
	var l list
	for {
		fmt.Println(`Type:
		1 - to add
		2 - to read
		3 - to change
		4 - to exit`)
		fmt.Scan(&action)
		if action == 1 {
			l = Add(scanner, l)
		}
		if action == 2 {
			List(l)
		}
		if action == 3 {
			fmt.Println()
		}
		if action == 4 {
			os.Exit(1)
		}
	}
}
